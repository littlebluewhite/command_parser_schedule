package time_template

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/dal/query"
	"command_parser_schedule/entry/e_time_template"
	"command_parser_schedule/util"
	"context"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Operate interface {
	List() ([]model.TimeTemplate, error)
	Find(ids []int32) ([]model.TimeTemplate, error)
	Create([]*e_time_template.TimeTemplateCreate) ([]model.TimeTemplate, error)
	Update([]*e_time_template.TimeTemplateUpdate) error
	Delete([]int32) error
	ReloadCache() error
}

type operate struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewOperate(dbs dbs.Dbs) Operate {
	o := &operate{
		db:    dbs.GetSql(),
		cache: dbs.GetCache(),
	}
	err := o.ReloadCache()
	if err != nil {
		panic("initial time template operate error")
	}
	return o
}

func (o *operate) getCacheMap() map[int]model.TimeTemplate {
	var cacheMap map[int]model.TimeTemplate
	if x, found := o.cache.Get("timeTemplates"); found {
		cacheMap = x.(map[int]model.TimeTemplate)
	} else {
		return make(map[int]model.TimeTemplate)
	}
	return cacheMap
}

func (o *operate) setCacheMap(cacheMap map[int]model.TimeTemplate) {
	o.cache.Set("timeTemplates", cacheMap, cache.NoExpiration)
}

func (o *operate) listDB() ([]*model.TimeTemplate, error) {
	t := query.Use(o.db).TimeTemplate
	ctx := context.Background()
	timeTemplates, err := t.WithContext(ctx).Preload(field.Associations).Find()
	if err != nil {
		return nil, err
	}
	return timeTemplates, nil
}

func (o *operate) listCache() ([]model.TimeTemplate, error) {
	var tt []model.TimeTemplate
	cacheMap := o.getCacheMap()
	fmt.Println(cacheMap)
	for _, value := range cacheMap {
		tt = append(tt, value)
	}
	return tt, nil
}

func (o *operate) List() ([]model.TimeTemplate, error) {
	return o.listCache()
}

func (o *operate) ReloadCache() (e error) {
	tt, err := o.listDB()
	if err != nil {
		e = err
		return
	}
	cacheMap := make(map[int]model.TimeTemplate)
	for i := 0; i < len(tt); i++ {
		entry := tt[i]
		cacheMap[int(entry.ID)] = *entry
	}
	o.setCacheMap(cacheMap)
	return
}

func (o *operate) findDB(ids []int32) ([]*model.TimeTemplate, error) {
	t := query.Use(o.db).TimeTemplate
	ctx := context.Background()
	timeTemplates, err := t.WithContext(ctx).Preload(field.Associations).Where(t.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return timeTemplates, nil
}

func (o *operate) findCache(ids []int32) ([]model.TimeTemplate, error) {
	tt := make([]model.TimeTemplate, 0, len(ids))
	var cacheMap map[int]model.TimeTemplate
	if x, found := o.cache.Get("timeTemplates"); found {
		cacheMap = x.(map[int]model.TimeTemplate)
	} else {
		return nil, errors.New("cache error")
	}
	for _, id := range ids {
		t, ok := cacheMap[int(id)]
		if !ok {
			return nil, fmt.Errorf("id: %v not found", id)
		}
		tt = append(tt, t)
	}
	return tt, nil
}

func (o *operate) Find(ids []int32) ([]model.TimeTemplate, error) {
	return o.findCache(ids)
}

func (o *operate) Create(c []*e_time_template.TimeTemplateCreate) ([]model.TimeTemplate, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	cacheMap := o.getCacheMap()
	timeTemplates := e_time_template.CreateConvert(c)
	result := make([]model.TimeTemplate, 0, len(timeTemplates))
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.TimeTemplate.WithContext(ctx).CreateInBatches(timeTemplates, 100); err != nil {
			return err
		}
		for _, t := range timeTemplates {
			cacheMap[int(t.ID)] = *t
			result = append(result, *t)
		}
		o.setCacheMap(cacheMap)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *operate) Update(u []*e_time_template.TimeTemplateUpdate) error {
	cacheMap := o.getCacheMap()
	tt := e_time_template.UpdateConvert(cacheMap, u)
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		for _, item := range tt {
			t := util.StructToMap(item)
			td := t["time_data"].(map[string]interface{})
			delete(t, "time_data")
			delete(t, "updated_at")
			delete(t, "created_at")
			delete(td, "id")
			if _, err := tx.TimeTemplate.WithContext(ctx).Where(tx.TimeTemplate.ID.Eq(item.ID)).Updates(
				t); err != nil {
				return err
			}
			if _, err := tx.TimeDatum.WithContext(ctx).Where(tx.TimeDatum.ID.Eq(item.TimeDataID)).Updates(
				td); err != nil {
				return err
			}
			cacheMap[int(item.ID)] = *item
		}
		o.setCacheMap(cacheMap)
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
		if _, err := tx.TimeTemplate.WithContext(ctx).Where(
			tx.TimeTemplate.ID.In(ids...)).Delete(); err != nil {
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
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
