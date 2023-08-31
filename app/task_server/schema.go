package task_server

import (
	"command_parser_schedule/entry/e_task"
	"command_parser_schedule/entry/e_task_template"
	"errors"
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

type getStagesResult struct {
	sns      []int32
	stageMap map[int32]stageMapValue
}

type stageMapValue struct {
	monitor []e_task_template.TaskStage
	execute []e_task_template.TaskStage
}

var cannotFindTemplate = errors.New("can not find task template")
