package task_server

import (
	"command_parser_schedule/entry/e_task"
	"sync"
)

type chs struct {
	rec chan e_task.Task
	mu  *sync.RWMutex
}

type executeParams struct {
	templateId     int
	triggerFrom    []string
	triggerAccount string
	token          string
}

type taskRec struct {
	Token   string `json:"token"`
	TaskId  string `json:"task_id"`
	Status  string `json:"status"`
	Message string `json:"message"`
}
