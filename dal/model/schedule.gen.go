// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"time"
)

const TableNameSchedule = "schedule"

// Schedule mapped from table <schedule>
type Schedule struct {
	ID          int32      `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name        string     `gorm:"column:name;not null" json:"name"`
	Description *string    `gorm:"column:description" json:"description"`
	TimeDataID  int32      `gorm:"column:time_data_id;not null" json:"time_data_id"`
	TaskID      *int32     `gorm:"column:task_id" json:"task_id"`
	Enabled     bool       `gorm:"column:enabled;not null" json:"enabled"`
	UpdatedAt   *time.Time `gorm:"column:updated_at;default:CURRENT_TIMESTAMP" json:"updated_at"`
	CreatedAt   *time.Time `gorm:"column:created_at;default:now()" json:"created_at"`
	TimeData    TimeDatum  `gorm:"foreignKey:time_data_id" json:"time_data"`
}

// TableName Schedule's table name
func (*Schedule) TableName() string {
	return TableNameSchedule
}
