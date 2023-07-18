package time_server

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/util/logFile"
	"context"
	"fmt"
	"time"
)

var (
	timeServerLog logFile.LogFile
)

func init() {
	timeServerLog = logFile.NewLogFile("app", "time_server")
}

type TimeServer interface {
	Start(ctx context.Context)
}

type timeServer[T any] struct {
	dbs      dbs.Dbs
	duration time.Duration
}

func NewTimeServer[T any](dbs dbs.Dbs, duration time.Duration) TimeServer {
	return &timeServer[T]{
		dbs:      dbs,
		duration: duration,
	}
}

func (ts *timeServer[T]) Start(ctx context.Context) {
	timeServerLog.Info().Println("Time server start")
	ticker := time.NewTicker(ts.duration)
	ctxChild, cancel := context.WithCancel(ctx)
	defer cancel()
	for {
		select {
		case <-ctx.Done():
			timeServerLog.Info().Println("Time server stop gracefully")
			return
		case t := <-ticker.C:
			go ts.checkSchedule(ctxChild, t)
			fmt.Println("Invoked at ", t)
		}
	}
}

func (ts *timeServer[T]) checkSchedule(ctx context.Context, t time.Time) {
	select {
	case <-ctx.Done():
		return
	default:
		var cacheMap map[int]model.Schedule
		if x, found := ts.dbs.GetCache().Get("schedule"); found {
			cacheMap = x.(map[int]model.Schedule)
		}
		for _, s := range cacheMap {
			go checkScheduleActive(s, t)
		}

	}
}
