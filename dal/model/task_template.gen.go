// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"encoding/json"
	"time"
)

const TableNameTaskTemplate = "task_template"

// TaskTemplate mapped from table <task_template>
type TaskTemplate struct {
	ID        int32           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Name      string          `gorm:"column:name;not null" json:"name"`
	Variable  json.RawMessage `gorm:"column:variable;default:json_object()" json:"variable"`
	UpdatedAt *time.Time      `gorm:"column:updated_at" json:"updated_at"`
	CreatedAt *time.Time      `gorm:"column:created_at" json:"created_at"`
	Tags      json.RawMessage `gorm:"column:tags;default:json_array()" json:"tags"`
	Stages    []TaskStage     `gorm:"many2many:task_template_stage" json:"stages"`
}

// TableName TaskTemplate's table name
func (*TaskTemplate) TableName() string {
	return TableNameTaskTemplate
}
