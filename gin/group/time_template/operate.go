package time_template

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
	List() ([]*model.TimeTemplate, error)
	Find(ids []int32) ([]*model.TimeTemplate, error)
	Create([]*model.TimeTemplate) ([]*model.TimeTemplate, error)
	Update([]*model.TimeTemplate) error
	Delete([]*model.TimeTemplate) error
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

func (o *operate) List() ([]*model.TimeTemplate, error) {
	t := query.Use(o.db).TimeTemplate
	ctx := context.Background()
	timeTemplates, err := t.WithContext(ctx).Preload(field.Associations).Find()
	if err != nil {
		return nil, err
	}
	return timeTemplates, nil
}

func (o *operate) Find(ids []int32) ([]*model.TimeTemplate, error) {
	t := query.Use(o.db).TimeTemplate
	ctx := context.Background()
	timeTemplates, err := t.WithContext(ctx).Preload(field.Associations).Where(t.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return timeTemplates, nil
}

func (o *operate) Create(timeTemplates []*model.TimeTemplate) ([]*model.TimeTemplate, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.TimeTemplate.WithContext(ctx).Create(timeTemplates...); err != nil {
			return err
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
