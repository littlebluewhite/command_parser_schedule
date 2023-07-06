package task_template

import (
	"command_parser_schedule/gin/group/command_template"
	"encoding/json"
	"time"
)

type TaskTemplate struct {
	ID        int32           `json:"id"`
	Name      string          `json:"name"`
	Variable  json.RawMessage `json:"variable"`
	UpdatedAt *time.Time      `json:"updated_at"`
	CreatedAt *time.Time      `json:"created_at"`
	Stages    []*TaskStage    `json:"stages"`
}

type TaskStage struct {
	Name              string                           `json:"name"`
	StageNumber       int32                            `json:"stage_number"`
	Mode              string                           `json:"mode"`
	CommandTemplateID *int32                           `json:"command_template_id"`
	Tag               json.RawMessage                  `json:"tag"`
	CommandTemplate   command_template.CommandTemplate `json:"command_template"`
}

type TaskTemplateCreate struct {
	ID       int32           `json:"id"`
	Name     string          `json:"name"`
	Variable json.RawMessage `json:"variable"`
	Stages   []*TaskStage    `json:"stages"`
}

type TaskStageCreate struct {
	Name              string          `json:"name"`
	StageNumber       int32           `json:"stage_number"`
	Mode              string          `json:"mode"`
	CommandTemplateID *int32          `json:"command_template_id"`
	Tag               json.RawMessage `json:"tag"`
}
