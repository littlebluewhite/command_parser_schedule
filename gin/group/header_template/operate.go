package header_template

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
	List() ([]*model.HeaderTemplate, error)
	Find(ids []int32) ([]*model.HeaderTemplate, error)
	Create([]*model.HeaderTemplate) ([]*model.HeaderTemplate, error)
	Update([]*model.HeaderTemplate) error
	Delete([]*model.HeaderTemplate) error
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

func (o *operate) List() ([]*model.HeaderTemplate, error) {
	t := query.Use(o.db).HeaderTemplate
	ctx := context.Background()
	ht, err := t.WithContext(ctx).Find()
	if err != nil {
		return nil, err
	}
	return ht, nil
}

func (o *operate) Find(ids []int32) ([]*model.HeaderTemplate, error) {
	t := query.Use(o.db).HeaderTemplate
	ctx := context.Background()
	HeaderTemplates, err := t.WithContext(ctx).Preload(field.Associations).Where(t.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return HeaderTemplates, nil
}

func (o *operate) Create(HeaderTemplates []*model.HeaderTemplate) ([]*model.HeaderTemplate, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.HeaderTemplate.WithContext(ctx).Create(HeaderTemplates...); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return HeaderTemplates, nil
}

func (o *operate) Update(HeaderTemplates []*model.HeaderTemplate) error {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		for _, item := range HeaderTemplates {
			t := util.StructToMap(item)
			if _, err := tx.HeaderTemplate.WithContext(ctx).Where(tx.HeaderTemplate.ID.Eq(item.ID)).Updates(
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

func (o *operate) Delete(HeaderTemplate []*model.HeaderTemplate) error {
	ids := make([]int32, 0, len(HeaderTemplate))
	for _, t := range HeaderTemplate {
		ids = append(ids, t.ID)
	}
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if _, err := tx.HeaderTemplate.WithContext(ctx).Where(
			tx.HeaderTemplate.ID.In(ids...)).Delete(); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return err
	}
	return nil
}
