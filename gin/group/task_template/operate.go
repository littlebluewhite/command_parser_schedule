package task_template

import (
	"command_parser_schedule/dal/model"
	"command_parser_schedule/dal/query"
	"command_parser_schedule/gin/initial"
	"command_parser_schedule/util"
	"context"
	"github.com/patrickmn/go-cache"
	"gorm.io/gen/field"
	"gorm.io/gorm"
)

type Operate interface {
	List() ([]*model.TaskTemplate, error)
	Find(ids []int32) ([]*model.TaskTemplate, error)
	Create([]*model.TaskTemplate) ([]*model.TaskTemplate, error)
	//Update([]*model.TaskTemplate) error
	Delete([]*model.TaskTemplate) error
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

func (o *operate) List() ([]*model.TaskTemplate, error) {
	t := query.Use(o.db).TaskTemplate
	ctx := context.Background()
	tt, err := t.WithContext(ctx).Preload(field.Associations).Preload(t.Stages.Command_Template).Find()
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (o *operate) Find(ids []int32) ([]*model.TaskTemplate, error) {
	t := query.Use(o.db).TaskTemplate
	ctx := context.Background()
	TaskTemplates, err := t.WithContext(ctx).Preload(field.Associations).Preload(t.Stages.Command_Template).Where(c.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return TaskTemplates, nil
}

func (o *operate) Create(TaskTemplates []*model.TaskTemplate) ([]*model.TaskTemplate, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.TaskTemplate.WithContext(ctx).Create(TaskTemplates...); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return TaskTemplates, nil
}

func (o *operate) Update(TaskTemplates []*model.TaskTemplate) error {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		for _, item := range TaskTemplates {
			t := util.StructToMap(item)
			if _, err := tx.TaskTemplate.WithContext(ctx).Where(tx.TaskTemplate.ID.Eq(item.ID)).Updates(
				t); err != nil {
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

func (o *operate) Delete(TaskTemplate []*model.TaskTemplate) error {
	ids := make([]int32, 0, len(TaskTemplate))
	for _, t := range TaskTemplate {
		ids = append(ids, t.ID)
	}
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if _, err := tx.TaskTemplate.WithContext(ctx).Where(
			tx.TaskTemplate.ID.In(ids...)).Delete(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
