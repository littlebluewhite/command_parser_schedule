package command_server

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/util/logFile"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"sync"
	"time"
)

type CommandServer interface {
}

type commandServer struct {
	dbs dbs.Dbs
	l   logFile.LogFile
	c   map[string]command
	chs chs
}

func NewCommandServer(dbs dbs.Dbs) CommandServer {
	l := logFile.NewLogFile("app", "command_server")
	c := make(map[string]command)
	r := make(chan command)
	send := make(chan command)
	mu := new(sync.RWMutex)
	return &commandServer{
		dbs: dbs,
		l:   l,
		c:   c,
		chs: chs{
			rec:  r,
			send: send,
			mu:   mu,
		},
	}
}

func (c *commandServer) Start(removeTime time.Duration) {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		c.receive(ctx)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		c.removeFinishedCommand(ctx, removeTime)
		wg.Done()
	}(wg)
	wg.Wait()
}

func (c *commandServer) receive(ctx context.Context) {
Loop1:
	for {
		select {
		case <-ctx.Done():
			break Loop1
		default:
			com := <-c.chs.rec
			go c.doCommand(com)
		}
	}
}

func (c *commandServer) doCommand(com command) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(com.Template.Timeout)*time.Second)
	defer cancel()
	c.chs.mu.Lock()
	com.Status = Process
	com.CancelFunc = cancel
	c.c[com.CommandId] = com
	c.chs.mu.Unlock()
	result := c.requestProtocol(ctx, com)
	now := time.Now()
	com.To = &now
	com.RespData = result.respData
	com.Status = result.status
	com.Message = result.message
	com.CancelFunc = nil
	c.chs.mu.Lock()
	c.c[com.CommandId] = com
	c.chs.mu.Unlock()
	c.writeToHistory(com)
}

func (c *commandServer) Execute(templateId int, triggerFrom []string, triggerAccount string) (commandId string, err error) {
	cache := c.dbs.GetCache()
	var cacheMap map[int]model.CommandTemplate
	if x, found := cache.Get("commandTemplates"); found {
		cacheMap = x.(map[int]model.CommandTemplate)
	}
	ct, ok := cacheMap[templateId]
	if !ok {
		err = errors.New("can not find command template")
		return
	}
	from := time.Now()
	commandId = fmt.Sprintf("%v_%v_%v_%v", templateId, ct.Name, ct.Protocol, from.UnixMicro())
	com := command{
		CommandId:      commandId,
		From:           from,
		TriggerFrom:    triggerFrom,
		TriggerAccount: triggerAccount,
		TemplateID:     templateId,
		Template:       model2template(ct),
	}
	c.chs.rec <- com
	return
}

func (c *commandServer) CancelCommand(commandId string) error {
	c.chs.mu.RLock()
	com, ok := c.c[commandId]
	c.chs.mu.RUnlock()
	if !ok {
		return errors.New("can not find command")
	}
	if com.Status != Process {
		return fmt.Errorf("command id %v finished. Can not cancel", commandId)
	} else {
		com.CancelFunc()
	}
	return nil
}

func (c *commandServer) ShowCommandList() (cs []command) {
	c.chs.mu.RLock()
	defer c.chs.mu.RUnlock()
	for _, item := range c.c {
		cs = append(cs, item)
	}
	return
}

func (c *commandServer) removeFinishedCommand(ctx context.Context, s time.Duration) {
Loop1:
	for {
		select {
		case <-ctx.Done():
			break Loop1
		default:
			c.chs.mu.Lock()
			now := time.Now()
			for cId, item := range c.c {
				if item.Status != Process && item.To.Add(s).After(now) {
					delete(c.c, cId)
				}
			}
			time.Sleep(s)
		}
	}
}

func (c *commandServer) writeToHistory(com command) {
	ctx := context.Background()
	p := influxdb2.NewPoint("command_history",
		map[string]string{"command_id": com.CommandId, "status": com.Status.String()},
		map[string]interface{}{"data": com},
		com.From,
	)
	if err := c.dbs.GetIdb().Writer().WritePoint(ctx, p); err != nil {
		panic(err)
	}
}

func (c *commandServer) ReadFromHistory(commandId, start, stop, status string) (hc []command) {
	ctx := context.Background()
	stopValue := ""
	if stop != "" {
		stopValue = fmt.Sprintf(", stop: %s", stop)
	}
	statusValue := ""
	if status != "" {
		statusValue = fmt.Sprintf(`|> filter(fn: (r) => r.status == "%s"`, status)
	}
	stmt := fmt.Sprintf(`from(bucket:"schedule"
|> range(start: %s%s)
|> filter(fn: (r) => r._measurement == "command_history"
|> filter(fn: (r) => r.command_id == "%s")
|> filter(fn: (r) => r."_field" == "data")
%s
`, start, stopValue, commandId, statusValue)
	result, err := c.dbs.GetIdb().Querier().Query(ctx, stmt)
	if err == nil {
		for result.Next() {
			var com command
			v := result.Record().Value()
			if e := json.Unmarshal([]byte(v.(string)), &com); e != nil {
				panic(e)
			}
			hc = append(hc, com)
		}
	} else {
		panic(err)
	}
	return
}
