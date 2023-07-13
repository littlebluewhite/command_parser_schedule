package header_template

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/dal/query"
	"command_parser_schedule/util"
	"context"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Operate interface {
	List() ([]model.HeaderTemplate, error)
	Find(ids []int32) ([]model.HeaderTemplate, error)
	Create([]*HeaderTemplateCreate) ([]model.HeaderTemplate, error)
	Update([]*HeaderTemplateUpdate) error
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
		panic("initial header template operate error")
	}
	return o
}

func (o *operate) getCacheMap() map[int]model.HeaderTemplate {
	var cacheMap map[int]model.HeaderTemplate
	if x, found := o.cache.Get("headerTemplates"); found {
		cacheMap = x.(map[int]model.HeaderTemplate)
	} else {
		return make(map[int]model.HeaderTemplate)
	}
	return cacheMap
}

func (o *operate) setCacheMap(cacheMap map[int]model.HeaderTemplate) {
	o.cache.Set("headerTemplates", cacheMap, cache.NoExpiration)
}

func (o *operate) listDB() ([]*model.HeaderTemplate, error) {
	t := query.Use(o.db).HeaderTemplate
	ctx := context.Background()
	headerTemplates, err := t.WithContext(ctx).Preload(field.Associations).Find()
	if err != nil {
		return nil, err
	}
	return headerTemplates, nil
}

func (o *operate) List() ([]model.HeaderTemplate, error) {
	return o.listCache()
}

func (o *operate) listCache() ([]model.HeaderTemplate, error) {
	var tt []model.HeaderTemplate
	cacheMap := o.getCacheMap()
	fmt.Println(cacheMap)
	for _, value := range cacheMap {
		tt = append(tt, value)
	}
	return tt, nil
}

func (o *operate) ReloadCache() (e error) {
	tt, err := o.listDB()
	if err != nil {
		e = err
		return
	}
	cacheMap := make(map[int]model.HeaderTemplate)
	for i := 0; i < len(tt); i++ {
		entry := tt[i]
		cacheMap[int(entry.ID)] = *entry
	}
	o.setCacheMap(cacheMap)
	return
}

func (o *operate) findDB(ids []int32) ([]*model.HeaderTemplate, error) {
	t := query.Use(o.db).HeaderTemplate
	ctx := context.Background()
	headerTemplates, err := t.WithContext(ctx).Preload(field.Associations).Where(t.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return headerTemplates, nil
}

func (o *operate) Find(ids []int32) ([]model.HeaderTemplate, error) {
	return o.findCache(ids)
}

func (o *operate) findCache(ids []int32) ([]model.HeaderTemplate, error) {
	tt := make([]model.HeaderTemplate, 0, len(ids))
	var cacheMap map[int]model.HeaderTemplate
	if x, found := o.cache.Get("headerTemplates"); found {
		cacheMap = x.(map[int]model.HeaderTemplate)
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

func (o *operate) Create(c []*HeaderTemplateCreate) ([]model.HeaderTemplate, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	cacheMap := o.getCacheMap()
	headerTemplates := CreateConvert(c)
	result := make([]model.HeaderTemplate, 0, len(headerTemplates))
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.HeaderTemplate.WithContext(ctx).CreateInBatches(headerTemplates, 100); err != nil {
			return err
		}
		for _, h := range headerTemplates {
			cacheMap[int(h.ID)] = *h
			result = append(result, *h)
		}
		o.setCacheMap(cacheMap)
		return nil
	})
	if err != nil {
		return nil, err
	}
	return result, nil
}

func (o *operate) Update(u []*HeaderTemplateUpdate) error {
	cacheMap := o.getCacheMap()
	ht := UpdateConvert(cacheMap, u)
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		for _, item := range ht {
			t := util.StructToMap(item)
			if _, err := tx.HeaderTemplate.WithContext(ctx).Where(tx.HeaderTemplate.ID.Eq(item.ID)).Updates(
				t); err != nil {
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
	for _, i := range ids {
		_, ok := cacheMap[int(i)]
		if !ok {
			return errors.New(fmt.Sprintf("id: %d not found", i))
		}
	}
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if _, err := tx.HeaderTemplate.WithContext(ctx).Where(
			tx.HeaderTemplate.ID.In(ids...)).Delete(); err != nil {
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
