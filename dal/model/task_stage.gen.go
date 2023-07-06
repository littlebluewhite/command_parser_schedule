// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"encoding/json"
)

const TableNameTaskStage = "task_stage"

// TaskStage mapped from table <task_stage>
type TaskStage struct {
	ID                int32            `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name              *string          `gorm:"column:name" json:"name"`
	StageNumber       *int32           `gorm:"column:stage_number" json:"stage_number"`
	Mode              *string          `gorm:"column:mode" json:"mode"`
	CommandTemplateID *int32           `gorm:"column:command_template_id" json:"command_template_id"`
	Tag               *json.RawMessage `gorm:"column:tag" json:"tag"`
	CommandTemplate   CommandTemplate  `gorm:"foreignKey:command_template_id" json:"command_template"`
}

// TableName TaskStage's table name
func (*TaskStage) TableName() string {
	return TableNameTaskStage
}
