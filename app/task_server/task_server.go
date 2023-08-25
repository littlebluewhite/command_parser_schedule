package task_server

import (
	"command_parser_schedule/app/command_server"
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/entry/e_task"
	"command_parser_schedule/util/logFile"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"sync"
	"time"
)

type TaskServer interface {
}

type taskServer struct {
	dbs dbs.Dbs
	l   logFile.LogFile
	t   map[string]e_task.Task
	cs  command_server.CommandServer
	chs chs
}

func NewTaskServer(dbs dbs.Dbs) TaskServer {
	l := logFile.NewLogFile("app", "task_server")
	t := make(map[string]e_task.Task)
	rec := make(chan e_task.Task)
	mu := new(sync.RWMutex)
	cs := command_server.NewCommandServer(dbs)
	return &taskServer{
		dbs: dbs,
		l:   l,
		t:   t,
		cs:  cs,
		chs: chs{
			rec: rec,
			mu:  mu,
		},
	}
}

func (t *taskServer) Start(removeTime time.Duration) {
	t.cs.Start(removeTime)
}

func (t *taskServer) receive(ctx context.Context) {
Loop1:
	for {
		select {
		case <-ctx.Done():
			break Loop1
		default:
			task := <-t.chs.rec
			go t.doTask(task)
		}
	}
}

func (t *taskServer) Execute(ep executeParams) (taskId string, err error) {
	ctx := context.Background()
	cache := t.dbs.GetCache()
	var cacheMap map[int]model.TaskTemplate
	if x, found := cache.Get("taskTemplates"); found {
		cacheMap = x.(map[int]model.TaskTemplate)
	}
	tt, ok := cacheMap[ep.templateId]
	if !ok {
		err = errors.New("can not find task template")
	}
	from := time.Now()
	taskId = fmt.Sprintf("%v_%v_%v", ep.templateId, tt.Name, from.UnixMicro())
	task := e_task.Task{
		TaskId:         taskId,
		Token:          ep.token,
		From:           from,
		TriggerFrom:    ep.triggerFrom,
		TriggerAccount: ep.triggerAccount,
		TemplateID:     ep.templateId,
	}
	t.chs.rec <- task
	// publish to redis
	tr := taskRec{Token: ep.token, TaskId: taskId}
	_ = t.rdbPub(ctx, tr)
	return
}

func (t *taskServer) rdbPub(ctx context.Context, tr taskRec) (e error) {
	trb, _ := json.Marshal(tr)
	e = t.dbs.GetRdb().Publish(ctx, "taskRec", trb).Err()
	if e != nil {
		t.l.Error().Println("redis publish error")
		return
	}
	return
}
