// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"encoding/json"
	"time"
)

const TableNameTimeDatum = "time_data"

// TimeDatum mapped from table <time_data>
type TimeDatum struct {
	ID              int32           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	RepeatType      *string         `gorm:"column:repeat_type" json:"repeat_type"`
	StartDate       time.Time       `gorm:"column:start_date;not null" json:"start_date"`
	EndDate         *time.Time      `gorm:"column:end_date" json:"end_date"`
	StartTime       []byte          `gorm:"column:start_time;not null" json:"start_time"`
	EndTime         []byte          `gorm:"column:end_time;not null" json:"end_time"`
	IntervalSeconds *int32          `gorm:"column:interval_seconds" json:"interval_seconds"`
	MConditionType  *string         `gorm:"column:m_condition_type" json:"m_condition_type"`
	MCondition      json.RawMessage `gorm:"column:m_condition" json:"m_condition"`
}

// TableName TimeDatum's table name
func (*TimeDatum) TableName() string {
	return TableNameTimeDatum
}
