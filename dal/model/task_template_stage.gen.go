// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameTaskTemplateStage = "task_template_stage"

// TaskTemplateStage mapped from table <task_template_stage>
type TaskTemplateStage struct {
	ID             int32  `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	TaskTemplateID *int32 `gorm:"column:task_template_id" json:"task_template_id"`
	StageID        *int32 `gorm:"column:stage_id" json:"stage_id"`
}

// TableName TaskTemplateStage's table name
func (*TaskTemplateStage) TableName() string {
	return TableNameTaskTemplateStage
}