package command_server

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"command_parser_schedule/util/logFile"
	"errors"
	"fmt"
	"sync"
	"time"
)

type CommandServer interface {
}

type commandServer struct {
	dbs  dbs.Dbs
	l    logFile.LogFile
	c    map[string]command
	pipe chan int
}

func NewCommandServer(dbs dbs.Dbs) CommandServer {
	l := logFile.NewLogFile("app", "command_server")
	c := make(map[string]command)
	s := make(chan int)
	return &commandServer{
		dbs:  dbs,
		l:    l,
		c:    c,
		pipe: s,
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
		templateId := <-c.pipe
		go func(id int) {
			c.Execute(id)
		}(templateId)
	}
}

func (c *commandServer) Execute(templateId int, triggerFrom []string) (commandId string, err error) {
	cache := c.dbs.GetCache()
	var cacheMap map[int]model.CommandTemplate
	if x, found := cache.Get("commandTemplates"); found {
		cacheMap = x.(map[int]model.CommandTemplate)
	} else {

	}
	ct, ok := cacheMap[templateId]
	from := time.Now()
	commandId = fmt.Sprintf("%v_%v_%v_%v", templateId, ct.Name, ct.Protocol, from.UnixMicro())
	//com := command{
	//	CommandId:   commandId,
	//	From:        from,
	//	TriggerFrom: triggerFrom,
	//	TemplateID:  templateId,
	//	Template:    model2template(ct)}
	if !ok {
		err = errors.New("can not find command template")
		return
	}
	return
}
