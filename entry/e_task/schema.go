package e_task

import (
	"command_parser_schedule/entry/e_task_template"
	"time"
)

type Task struct {
	TaskId         string                       `json:"task_id"`
	Token          string                       `json:"token"`
	From           time.Time                    `json:"from"`
	To             *time.Time                   `json:"to"`
	TriggerFrom    []string                     `json:"trigger_from"`
	TriggerAccount string                       `json:"trigger_account"`
	Status         Status                       `json:"status"`
	Message        string                       `json:"message"`
	TemplateID     int                          `json:"template_id"`
	Template       e_task_template.TaskTemplate `json:"template"`
}

type Status struct {
	IsSuccess       bool   `json:"is_success"`
	Stages          int    `json:"stages"`
	FailedCommandId string `json:"failed_command_id"`
	FailedMessage   string `json:"failed_message"`
}
