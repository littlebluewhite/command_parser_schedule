package time_template

import (
	"command_parser_schedule/dal/model"
	"command_parser_schedule/dal/query"
	"command_parser_schedule/gin/initial"
	"command_parser_schedule/util"
	"context"
	"errors"
	"github.com/patrickmn/go-cache"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Operate interface {
	List() ([]*model.TimeTemplate, error)
	Find(ids []int32) ([]*model.TimeTemplate, error)
	Create([]*model.TimeTemplate) ([]*model.TimeTemplate, error)
	Update([]*model.TimeTemplate) error
	Delete([]*model.TimeTemplate) error
	ReloadCache() error
}

type operate struct {
	db    *gorm.DB
	cache *cache.Cache
}

func NewOperate(dbs initial.Dbs) Operate {
	return &operate{
		db:    dbs.GetSql(),
		cache: dbs.GetCache(),
	}
}

func (o *operate) getCacheMap() (map[int]model.TimeTemplate, error) {
	var cacheMap map[int]model.TimeTemplate
	if x, found := o.cache.Get("timeTemplates"); found {
		cacheMap = x.(map[int]model.TimeTemplate)
	} else {
		return nil, errors.New("cache error")
	}
	return cacheMap, nil
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

func (o *operate) List() ([]*model.TimeTemplate, error) {
	return o.listCache()
}

func (o *operate) listCache() ([]*model.TimeTemplate, error) {
	var tt []*model.TimeTemplate
	cacheMap, err := o.getCacheMap()
	if err != nil {
		return nil, err
	}
	for _, value := range cacheMap {
		tt = append(tt, &value)
	}
	return tt, nil
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
	o.cache.Set("timeTemplates", cacheMap, cache.NoExpiration)
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

func (o *operate) Find(ids []int32) ([]*model.TimeTemplate, error) {
	return o.findCache(ids)
}

func (o *operate) findCache(ids []int32) ([]*model.TimeTemplate, error) {
	tt := make([]*model.TimeTemplate, 0, len(ids))
	var cacheMap map[int]model.TimeTemplate
	if x, found := o.cache.Get("timeTemplates"); found {
		cacheMap = x.(map[int]model.TimeTemplate)
	} else {
		return nil, errors.New("cache error")
	}
	for _, id := range ids {
		t := cacheMap[int(id)]
		tt = append(tt, &t)
	}
	return tt, nil
}

func (o *operate) Create(timeTemplates []*model.TimeTemplate) ([]*model.TimeTemplate, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.TimeTemplate.WithContext(ctx).Create(timeTemplates...); err != nil {
			return err
		}
		for _, t := range timeTemplates {

		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return timeTemplates, nil
}

func (o *operate) Update(timeTemplates []*model.TimeTemplate) error {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		for _, item := range timeTemplates {
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
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (o *operate) Delete(timeTemplate []*model.TimeTemplate) error {
	tId := make([]int32, 0, len(timeTemplate))
	tdId := make([]int32, 0, len(timeTemplate))
	for _, t := range timeTemplate {
		tId = append(tId, t.ID)
		tdId = append(tdId, t.TimeDataID)
	}
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if _, err := tx.TimeTemplate.WithContext(ctx).Where(
			tx.TimeTemplate.ID.In(tId...)).Delete(); err != nil {
			return err
		}
		if _, err := tx.TimeDatum.WithContext(ctx).Where(
			tx.TimeDatum.ID.In(tdId...)).Delete(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
