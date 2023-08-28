package task_server

import (
	"command_parser_schedule/entry/e_task"
	"command_parser_schedule/entry/e_task_template"
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

type getStagesResult struct {
	sns      []int32
	stageMap map[int32]stageMapValue
}

type stageMapValue struct {
	monitor []e_task_template.TaskStage
	execute []e_task_template.TaskStage
}
