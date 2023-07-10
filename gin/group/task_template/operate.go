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
	Update([]*model.TaskTemplate) error
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
	tt, err := t.WithContext(ctx).Preload(field.Associations).Preload(t.Stages.CommandTemplate).Find()
	if err != nil {
		return nil, err
	}
	return tt, nil
}

func (o *operate) Find(ids []int32) ([]*model.TaskTemplate, error) {
	t := query.Use(o.db).TaskTemplate
	ctx := context.Background()
	TaskTemplates, err := t.WithContext(ctx).Preload(field.Associations).Preload(t.Stages.CommandTemplate).Where(t.ID.In(ids...)).Find()
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

func (o *operate) Delete(TaskTemplate []*model.TaskTemplate) error {
	ids := make([]int32, 0, len(TaskTemplate))
	sIds := make([]int32, 0, 20)
	for _, t := range TaskTemplate {
		ids = append(ids, t.ID)
		for _, s := range t.Stages {
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
