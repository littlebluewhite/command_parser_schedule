package schedule

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
	List() ([]*model.Schedule, error)
	Find(ids []int32) ([]*model.Schedule, error)
	Create([]*model.Schedule) ([]*model.Schedule, error)
	Update([]*model.Schedule) error
	Delete([]*model.Schedule) error
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

func (o *operate) List() ([]*model.Schedule, error) {
	t := query.Use(o.db).Schedule
	ctx := context.Background()
	schedules, err := t.WithContext(ctx).Preload(field.Associations).Find()
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (o *operate) Find(ids []int32) ([]*model.Schedule, error) {
	t := query.Use(o.db).Schedule
	ctx := context.Background()
	schedules, err := t.WithContext(ctx).Preload(field.Associations).Where(t.ID.In(ids...)).Find()
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (o *operate) Create(schedules []*model.Schedule) ([]*model.Schedule, error) {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if err := tx.Schedule.WithContext(ctx).Create(schedules...); err != nil {
			return err
		}
		return nil
	})
	if err != nil {
		return nil, err
	}
	return schedules, nil
}

func (o *operate) Update(schedules []*model.Schedule) error {
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		for _, item := range schedules {
			s := util.StructToMap(item)
			td := s["time_data"].(map[string]interface{})
			util.MapDeleteNil(s)
			delete(s, "time_data")
			delete(s, "updated_at")
			delete(s, "created_at")
			delete(td, "id")
			if _, err := tx.Schedule.WithContext(ctx).Where(tx.Schedule.ID.Eq(item.ID)).Updates(
				s); err != nil {
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

func (o *operate) Delete(schedule []*model.Schedule) error {
	sId := make([]int32, 0, len(schedule))
	tdId := make([]int32, 0, len(schedule))
	for _, t := range schedule {
		sId = append(sId, t.ID)
		tdId = append(tdId, t.TimeDataID)
	}
	q := query.Use(o.db)
	ctx := context.Background()
	err := q.Transaction(func(tx *query.Query) error {
		if _, err := tx.Schedule.WithContext(ctx).Where(
			tx.Schedule.ID.In(sId...)).Delete(); err != nil {
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
