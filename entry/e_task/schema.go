package e_task

import (
	"command_parser_schedule/entry/e_task_template"
	"github.com/goccy/go-json"
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
	CancelFunc     func()
}

type TaskPub struct {
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
	TStatus         TStatus `json:"task_status"`
	Stages          int     `json:"stages"`
	FailedCommandId string  `json:"failed_command_id"`
	FailedMessage   string  `json:"failed_message"`
}

type TStatus int

const (
	Prepared TStatus = iota
	Process
	Success
	Failure
	Cancel
)

func (s TStatus) String() string {
	return [...]string{"Prepared", "Process", "Success", "Failure", "Cancel"}[s]
}

func (s TStatus) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.String())
}
