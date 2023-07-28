package command_server

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/util/logFile"
	"context"
	"errors"
	"fmt"
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

func (c *commandServer) Start() {
	wg := &sync.WaitGroup{}
	go func(wg *sync.WaitGroup) {
		wg.Add(1)
		c.receive()
		wg.Done()
	}(wg)
	wg.Wait()
}

func (c *commandServer) receive() {
	for {
		com := <-c.chs.rec
		go c.doCommand(com)
	}
}

func (c *commandServer) doCommand(com command) {
	ctx, cancel := context.WithTimeout(context.Background(),
		time.Duration(com.Template.Timeout)*time.Second)
	defer cancel()
	c.chs.mu.Lock()
	com.Status = Process
	c.c[com.CommandId] = com
	c.chs.mu.Unlock()
	result := c.requestProtocol(ctx, com)
}

func (c *commandServer) Execute(templateId int, triggerFrom []string) (commandId string, err error) {
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
		CommandId:   commandId,
		From:        from,
		TriggerFrom: triggerFrom,
		TemplateID:  templateId,
		Template:    model2template(ct),
	}
	c.chs.rec <- com
	return
}
