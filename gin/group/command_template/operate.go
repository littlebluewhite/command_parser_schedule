package command_template

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/dal/query"
	"context"
	"errors"
	"fmt"
	"github.com/patrickmn/go-cache"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Operate interface {
	List() ([]model.CommandTemplate, error)
	Find(ids []int32) ([]model.CommandTemplate, error)
	Create([]*CommandTemplateCreate) ([]model.CommandTemplate, error)
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
		panic("initial command template operate error")
	}
	return o
}

func (o *operate) getCacheMap() map[int]model.CommandTemplate {
	var cacheMap map[int]model.CommandTemplate
	if x, found := o.cache.Get("commandTemplates"); found {
		cacheMap = x.(map[int]model.CommandTemplate)
	} else {
		return make(map[int]model.CommandTemplate)
	}
	return cacheMap
}

func (o *operate) setCacheMap(cacheMap map[int]model.CommandTemplate) {
	o.cache.Set("commandTemplates", cacheMap, cache.NoExpiration)
}

func (o *operate) listDB() ([]*model.CommandTemplate, error) {
	c := query.Use(o.db).CommandTemplate
	ctx := context.Background()
	ct, err := c.WithContext(ctx).Preload(field.Associations).Preload(c.Monitor.MConditions).Find()
	if err != nil {
		return nil, err
	}
	return ct, nil
}

func (o *operate) listCache() ([]model.CommandTemplate, error) {
	var tt []model.CommandTemplate
	cacheMap := o.getCacheMap()
	fmt.Println(cacheMap)
	for _, value := range cacheMap {
		tt = append(tt, value)
	}
	return tt, nil
}

func (o *operate) List() ([]model.CommandTemplate, error) {
	return o.listCache()
}

func (o *operate) ReloadCache() (e error) {
	tt, err := o.listDB()
	if err != nil {
		e = err
		return
	}
	cacheMap := make(map[int]model.CommandTemplate)
	for i := 0; i < len(tt); i++ {
		entry := tt[i]
		cacheMap[int(entry.ID)] = *entry
	}
	o.setCacheMap(cacheMap)
	return
}
func (o *operate) findDB(ids []int32) ([]*model.CommandTemplate, error) {
	c := query.Use(o.db).CommandTemplate
	ctx := context.Background()
	CommandTemplates, err := c.WithContext(ctx).Preload(field.Associations).Preload(c.Monitor.MConditions).Where(c.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return CommandTemplates, nil
}

func (o *operate) findCache(ids []int32) ([]model.CommandTemplate, error) {
	tt := make([]model.CommandTemplate, 0, len(ids))
	var cacheMap map[int]model.CommandTemplate
	if x, found := o.cache.Get("commandTemplates"); found {
		cacheMap = x.(map[int]model.CommandTemplate)
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

func (o *operate) Find(ids []int32) ([]model.CommandTemplate, error) {
	return o.findCache(ids)
}

func (o *operate) Create(c []*CommandTemplateCreate) ([]model.CommandTemplate, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	cacheMap := o.getCacheMap()
	commandTemplates := CreateConvert(c)
	result := make([]model.CommandTemplate, 0, len(commandTemplates))
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.CommandTemplate.WithContext(ctx).CreateInBatches(commandTemplates, 100); err != nil {
			return err
		}
		for _, t := range commandTemplates {
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

func (o *operate) Delete(ids []int32) error {
	cacheMap := o.getCacheMap()
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if _, err := tx.CommandTemplate.WithContext(ctx).Where(
			tx.CommandTemplate.ID.In(ids...)).Delete(); err != nil {
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
