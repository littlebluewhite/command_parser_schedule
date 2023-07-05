package command_template

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
	List() ([]*model.CommandTemplate, error)
	Find(ids []int32) ([]*model.CommandTemplate, error)
	Create([]*model.CommandTemplate) ([]*model.CommandTemplate, error)
	//Update([]*model.CommandTemplate) error
	Delete([]*model.CommandTemplate) error
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

func (o *operate) List() ([]*model.CommandTemplate, error) {
	c := query.Use(o.db).CommandTemplate
	ctx := context.Background()
	ct, err := c.WithContext(ctx).Preload(field.Associations).Preload(c.Monitor.MConditions).Find()
	if err != nil {
		return nil, err
	}
	return ct, nil
}

func (o *operate) Find(ids []int32) ([]*model.CommandTemplate, error) {
	c := query.Use(o.db).CommandTemplate
	ctx := context.Background()
	CommandTemplates, err := c.WithContext(ctx).Preload(field.Associations).Preload(c.Monitor.MConditions).Where(c.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return CommandTemplates, nil
}

func (o *operate) Create(CommandTemplates []*model.CommandTemplate) ([]*model.CommandTemplate, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.CommandTemplate.WithContext(ctx).Create(CommandTemplates...); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return CommandTemplates, nil
}

func (o *operate) Update(CommandTemplates []*model.CommandTemplate) error {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		for _, item := range CommandTemplates {
			t := util.StructToMap(item)
			if _, err := tx.CommandTemplate.WithContext(ctx).Where(tx.CommandTemplate.ID.Eq(item.ID)).Updates(
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

func (o *operate) Delete(CommandTemplate []*model.CommandTemplate) error {
	ids := make([]int32, 0, len(CommandTemplate))
	for _, t := range CommandTemplate {
		ids = append(ids, t.ID)
	}
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if _, err := tx.CommandTemplate.WithContext(ctx).Where(
			tx.CommandTemplate.ID.In(ids...)).Delete(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
