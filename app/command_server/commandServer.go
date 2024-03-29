package command_server

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/entry/e_command"
	"command_parser_schedule/entry/e_command_template"
	"command_parser_schedule/util/logFile"
	"context"
	"errors"
	"fmt"
	"github.com/goccy/go-json"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"sync"
	"time"
)

type CommandServer struct {
	dbs dbs.Dbs
	l   logFile.LogFile
	c   map[string]e_command.Command
	chs chs
}

func NewCommandServer(dbs dbs.Dbs) *CommandServer {
	l := logFile.NewLogFile("app", "command_server")
	c := make(map[string]e_command.Command)
	rec := make(chan e_command.Command)
	mu := new(sync.RWMutex)
	return &CommandServer{
		dbs: dbs,
		l:   l,
		c:   c,
		chs: chs{
			rec: rec,
			mu:  mu,
		},
	}
}

func (c *CommandServer) Start(removeTime time.Duration) {
	wg := &sync.WaitGroup{}
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	wg.Add(2)
	go func(wg *sync.WaitGroup) {
		c.removeFinishedCommand(ctx, removeTime)
		wg.Done()
	}(wg)
	go func(wg *sync.WaitGroup) {
		c.rdbSub(ctx)
		wg.Done()
	}(wg)
	wg.Wait()
}

func (c *CommandServer) rdbSub(ctx context.Context) {
	pubsub := c.dbs.GetRdb().Subscribe(ctx, "sendCommand")
	for {
		msg, err := pubsub.ReceiveMessage(ctx)
		if err != nil {
			panic(err)
		}
		b := []byte(msg.Payload)
		var s SendCommand
		err = json.Unmarshal(b, &s)
		if err != nil {
			continue
		}
		ep := executeParams{s.TemplateId, s.TriggerFrom, s.TriggerAccount, s.Token}
		_, err = c.execute(ep)
		if err != nil {
			c.l.Error().Println("Error executing Command")
		}
	}
}

func (c *CommandServer) doCommand(com e_command.Command) e_command.Command {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(com.Template.Timeout)*time.Millisecond)
	defer cancel()
	c.chs.mu.Lock()
	com.Status = e_command.Process
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
	// write to history in influxdb
	c.writeToHistory(com)
	// send to redis channel
	if e := c.rdbPub(ctx, com); e != nil {
		panic(e)
	}
	return com
}

func (c *CommandServer) execute(ep executeParams) (commandId string, err error) {
	ctx := context.Background()
	com, err := c.generateCommand(ep)
	// publish to redis
	_ = c.rdbPub(ctx, com)
	if err != nil {
		c.l.Error().Println(err)
		return
	}
	go func() {
		c.doCommand(com)
	}()
	commandId = com.CommandId
	return
}

func (c *CommandServer) Execute(templateId int, triggerFrom []string,
	triggerAccount string, token string) (com e_command.Command, err error) {
	ctx := context.Background()
	ep := executeParams{
		templateId:     templateId,
		triggerFrom:    triggerFrom,
		triggerAccount: triggerAccount,
		token:          token,
	}
	com, err = c.generateCommand(ep)
	// publish to redis
	_ = c.rdbPub(ctx, com)
	if err != nil {
		c.l.Error().Println(err)
		return
	}
	ch := make(chan e_command.Command)
	go func() {
		ch <- c.doCommand(com)
	}()
	com = <-ch
	return
}

func (c *CommandServer) generateCommand(ep executeParams) (com e_command.Command, err error) {
	cache := c.dbs.GetCache()
	var cacheMap map[int]model.CommandTemplate
	if x, found := cache.Get("commandTemplates"); found {
		cacheMap = x.(map[int]model.CommandTemplate)
	}
	ct, ok := cacheMap[ep.templateId]
	if !ok {
		err = cannotFindTemplate
		com = e_command.Command{Token: ep.token, Message: "can not find Command template", Status: e_command.Failure}
		return
	}
	from := time.Now()
	commandId := fmt.Sprintf("%v_%v_%v_%v", ep.templateId, ct.Name, ct.Protocol, from.UnixMicro())
	com = e_command.Command{
		CommandId:      commandId,
		Token:          ep.token,
		From:           from,
		TriggerFrom:    ep.triggerFrom,
		TriggerAccount: ep.triggerAccount,
		TemplateID:     ep.templateId,
		Template:       e_command_template.Format([]model.CommandTemplate{ct})[0],
	}
	return
}

func (c *CommandServer) CancelCommand(commandId string) error {
	c.chs.mu.RLock()
	com, ok := c.c[commandId]
	c.chs.mu.RUnlock()
	if !ok {
		return errors.New("can not find Command")
	}
	if com.Status != e_command.Process {
		return fmt.Errorf("command id %v finished. Can not cancel", commandId)
	} else {
		com.CancelFunc()
	}
	return nil
}

func (c *CommandServer) ShowCommandList() (cs []e_command.Command) {
	c.chs.mu.RLock()
	defer c.chs.mu.RUnlock()
	for _, item := range c.c {
		cs = append(cs, item)
	}
	return
}

func (c *CommandServer) removeFinishedCommand(ctx context.Context, s time.Duration) {
Loop1:
	for {
		select {
		case <-ctx.Done():
			break Loop1
		default:
			c.chs.mu.Lock()
			now := time.Now()
			for cId, item := range c.c {
				if item.Status != e_command.Process && item.To.Add(s).After(now) {
					delete(c.c, cId)
				}
			}
			time.Sleep(s)
		}
	}
}

func (c *CommandServer) writeToHistory(com e_command.Command) {
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

func (c *CommandServer) ReadFromHistory(commandId, start, stop, status string) (hc []e_command.Command) {
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
			var com e_command.Command
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

func (c *CommandServer) rdbPub(ctx context.Context, com e_command.Command) (e error) {
	cb, _ := json.Marshal(e_command.ToPub(com))
	e = c.dbs.GetRdb().Publish(ctx, "CommandRec", cb).Err()
	if e != nil {
		c.l.Error().Println("redis publish error")
		return
	}
	return
}
