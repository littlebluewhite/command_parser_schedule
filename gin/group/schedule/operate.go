package schedule

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/app/time_server"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/dal/query"
	"command_parser_schedule/entry/e_schedule"
	"command_parser_schedule/util"
	"context"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Operate interface {
	List() ([]model.Schedule, error)
	Find(ids []int32) ([]model.Schedule, error)
	Create([]*e_schedule.ScheduleCreate) ([]model.Schedule, error)
	Update([]*e_schedule.ScheduleUpdate) error
	Delete([]int32) error
	ReloadCache() error
}

type operate struct {
	db    *gorm.DB
	cache *cache.Cache
	ts    time_server.TimeServer
}

func NewOperate(dbs dbs.Dbs, ts time_server.TimeServer) Operate {
	o := &operate{
		db:    dbs.GetSql(),
		cache: dbs.GetCache(),
		ts:    ts,
	}
	err := o.ReloadCache()
	if err != nil {
		panic("initial time template operate error")
	}
	return o
}

func (o *operate) getCacheMap() map[int]model.Schedule {
	var cacheMap map[int]model.Schedule
	if x, found := o.cache.Get("Schedules"); found {
		cacheMap = x.(map[int]model.Schedule)
	} else {
		return make(map[int]model.Schedule)
	}
	return cacheMap
}

func (o *operate) setCacheMap(cacheMap map[int]model.Schedule) {
	o.cache.Set("Schedules", cacheMap, cache.NoExpiration)
}

func (o *operate) listDB() ([]*model.Schedule, error) {
	t := query.Use(o.db).Schedule
	ctx := context.Background()
	schedules, err := t.WithContext(ctx).Preload(field.Associations).Find()
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (o *operate) listCache() ([]model.Schedule, error) {
	var s []model.Schedule
	cacheMap := o.getCacheMap()
	fmt.Println(cacheMap)
	for _, value := range cacheMap {
		s = append(s, value)
	}
	return s, nil
}

func (o *operate) List() ([]model.Schedule, error) {
	return o.listCache()
}

func (o *operate) ReloadCache() (e error) {
	s, err := o.listDB()
	if err != nil {
		e = err
		return
	}
	cacheMap := make(map[int]model.Schedule)
	for i := 0; i < len(s); i++ {
		entry := s[i]
		cacheMap[int(entry.ID)] = *entry
	}
	o.setCacheMap(cacheMap)
	o.ts.ReloadSchedule(cacheMap)
	return
}

func (o *operate) findDB(ids []int32) ([]*model.Schedule, error) {
	t := query.Use(o.db).Schedule
	ctx := context.Background()
	schedules, err := t.WithContext(ctx).Preload(field.Associations).Where(t.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (o *operate) findCache(ids []int32) ([]model.Schedule, error) {
	s := make([]model.Schedule, 0, len(ids))
	var cacheMap map[int]model.Schedule
	if x, found := o.cache.Get("Schedules"); found {
		cacheMap = x.(map[int]model.Schedule)
	} else {
		return nil, errors.New("cache error")
	}
	for _, id := range ids {
		t, ok := cacheMap[int(id)]
		if !ok {
			return nil, fmt.Errorf("id: %v not found", id)
		}
		s = append(s, t)
	}
	return s, nil
}

func (o *operate) Find(ids []int32) ([]model.Schedule, error) {
	return o.findCache(ids)
}

func (o *operate) Create(c []*e_schedule.ScheduleCreate) ([]model.Schedule, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	cacheMap := o.getCacheMap()
	Schedules := e_schedule.CreateConvert(c)
	result := make([]model.Schedule, 0, len(Schedules))
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.Schedule.WithContext(ctx).CreateInBatches(Schedules, 100); err != nil {
			return err
		}
		for _, t := range Schedules {
			cacheMap[int(t.ID)] = *t
			result = append(result, *t)
		}
		o.setCacheMap(cacheMap)
		o.ts.ReloadSchedule(cacheMap)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *operate) Update(u []*e_schedule.ScheduleUpdate) error {
	cacheMap := o.getCacheMap()
	s := e_schedule.UpdateConvert(cacheMap, u)
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		for _, item := range s {
			s := util.StructToMap(item)
			td := s["time_data"].(map[string]interface{})
			util.MapDeleteNil(s)
			delete(s, "time_data")
			delete(s, "updated_at")
			delete(s, "created_at")
			delete(td, "id")
			if _, err := tx.Schedule.WithContext(ctx).Where(tx.Schedule.ID.Eq(item.ID)).Updates(
				s); err != nil {
				return err
			}
			if _, err := tx.TimeDatum.WithContext(ctx).Where(tx.TimeDatum.ID.Eq(item.TimeDataID)).Updates(
				td); err != nil {
				return err
			}
			cacheMap[int(item.ID)] = *item
		}
		o.setCacheMap(cacheMap)
		o.ts.ReloadSchedule(cacheMap)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (o *operate) Delete(ids []int32) error {
	cacheMap := o.getCacheMap()
	tdId := make([]int32, 0, len(ids))
	for _, i := range ids {
		t, ok := cacheMap[int(i)]
		if !ok {
			return errors.New(fmt.Sprintf("id: %d not found", i))
		}
		tdId = append(tdId, t.TimeDataID)
	}
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if _, err := tx.Schedule.WithContext(ctx).Where(
			tx.Schedule.ID.In(ids...)).Delete(); err != nil {
			return err
		}
		if _, err := tx.TimeDatum.WithContext(ctx).Where(
			tx.TimeDatum.ID.In(tdId...)).Delete(); err != nil {
			return err
		}
		for _, id := range ids {
			delete(cacheMap, int(id))
		}
		o.setCacheMap(cacheMap)
		o.ts.ReloadSchedule(cacheMap)
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
