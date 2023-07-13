package task_template

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
	List() ([]model.TaskTemplate, error)
	Find(ids []int32) ([]model.TaskTemplate, error)
	Create([]*TaskTemplateCreate) ([]model.TaskTemplate, error)
	Update([]*TaskTemplateUpdate) error
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

func (o *operate) getCacheMap() map[int]model.TaskTemplate {
	var cacheMap map[int]model.TaskTemplate
	if x, found := o.cache.Get("taskTemplates"); found {
		cacheMap = x.(map[int]model.TaskTemplate)
	} else {
		return make(map[int]model.TaskTemplate)
	}
	return cacheMap
}

func (o *operate) setCacheMap(cacheMap map[int]model.TaskTemplate) {
	o.cache.Set("taskTemplates", cacheMap, cache.NoExpiration)
}

func (o *operate) listDB() ([]*model.TaskTemplate, error) {
	t := query.Use(o.db).TaskTemplate
	ctx := context.Background()
	tt, err := t.WithContext(ctx).Preload(field.Associations).Preload(t.Stages.CommandTemplate).Find()
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (o *operate) listCache() ([]model.TaskTemplate, error) {
	var tt []model.TaskTemplate
	cacheMap := o.getCacheMap()
	fmt.Println(cacheMap)
	for _, value := range cacheMap {
		tt = append(tt, value)
	}
	return tt, nil
}

func (o *operate) List() ([]model.TaskTemplate, error) {
	return o.listCache()
}

func (o *operate) ReloadCache() (e error) {
	tt, err := o.listDB()
	if err != nil {
		e = err
		return
	}
	cacheMap := make(map[int]model.TaskTemplate)
	for i := 0; i < len(tt); i++ {
		entry := tt[i]
		cacheMap[int(entry.ID)] = *entry
	}
	o.setCacheMap(cacheMap)
	return
}

func (o *operate) findDB(ids []int32) ([]*model.TaskTemplate, error) {
	t := query.Use(o.db).TaskTemplate
	ctx := context.Background()
	TaskTemplates, err := t.WithContext(ctx).Preload(field.Associations).Preload(t.Stages.CommandTemplate).Where(t.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return TaskTemplates, nil
}

func (o *operate) findCache(ids []int32) ([]model.TaskTemplate, error) {
	tt := make([]model.TaskTemplate, 0, len(ids))
	var cacheMap map[int]model.TaskTemplate
	if x, found := o.cache.Get("taskTemplates"); found {
		cacheMap = x.(map[int]model.TaskTemplate)
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

func (o *operate) Find(ids []int32) ([]model.TaskTemplate, error) {
	return o.findCache(ids)
}

func (o *operate) Create(c []*TaskTemplateCreate) ([]model.TaskTemplate, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	cacheMap := o.getCacheMap()
	taskTemplates := CreateConvert(c)
	result := make([]model.TaskTemplate, 0, len(taskTemplates))
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.TaskTemplate.WithContext(ctx).CreateInBatches(taskTemplates, 100); err != nil {
			return err
		}
		for _, t := range taskTemplates {
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

func (o *operate) Update(u []*TaskTemplateUpdate) error {
	cacheMap := o.getCacheMap()
	tt := UpdateConvert(cacheMap, u)
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		for _, item := range tt {
			t := util.StructToMap(item)
			sUpdate := make([]map[string]interface{}, 0, 10)
			sCreate := make([]*model.TaskStage, 0, 10)
			sDelete := make([]int32, 0, 10)
			for _, stage := range item.Stages {
				switch {
				case stage.ID < 0:
					sDelete = append(sDelete, -stage.ID)
				case stage.ID == 0:
					sCreate = append(sCreate, &stage)
				case stage.ID > 0:
					sUpdate = append(sUpdate, util.StructToMap(stage))
				}
			}
			t["stages"] = sUpdate
			delete(t, "stages")
			delete(t, "updated_at")
			delete(t, "created_at")
			if _, err := tx.TaskTemplate.WithContext(ctx).Where(tx.TaskTemplate.ID.Eq(
				item.ID)).Updates(t); err != nil {
				return err
			}
			for _, si := range sUpdate {
				delete(si, "command_template")
				if _, err := tx.TaskStage.WithContext(ctx).Where(tx.TaskStage.ID.Eq((si["id"]).(int32))).Updates(si); err != nil {
					return err
				}
			}
			if err := tx.TaskStage.WithContext(ctx).Create(sCreate...); err != nil {
				return err
			}
			tts := make([]*model.TaskTemplateStage, 0, len(sCreate))
			for _, ts := range sCreate {
				tts = append(tts, &model.TaskTemplateStage{
					TaskStageID: ts.ID, TaskTemplateID: item.ID})
			}
			if err := tx.TaskTemplateStage.WithContext(ctx).Create(tts...); err != nil {
				return err
			}
			if _, err := tx.TaskStage.WithContext(ctx).Where(tx.TaskStage.ID.In(
				sDelete...)).Delete(); err != nil {
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

func (o *operate) Delete(ids []int32) error {
	cacheMap := o.getCacheMap()
	sIds := make([]int32, 0, 20)
	for _, i := range ids {
		tt, ok := cacheMap[int(i)]
		if !ok {
			return errors.New(fmt.Sprintf("id: %d not found", i))
		}
		for _, s := range tt.Stages {
			sIds = append(sIds, s.ID)
		}
	}
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if _, err := tx.TaskTemplate.WithContext(ctx).Where(
			tx.TaskTemplate.ID.In(ids...)).Delete(); err != nil {
			return err
		}
		if _, err := tx.TaskStage.WithContext(ctx).Where(
			tx.TaskStage.ID.In(sIds...)).Delete(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
