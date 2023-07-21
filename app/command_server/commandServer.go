package command_server

import (
	"command_parser_schedule/app/dbs"
	"command_parser_schedule/dal/model"
	"errors"
)

type CommandServer interface {
}

type commandServer struct {
	dbs dbs.Dbs
	c   map[string]command
}

func NewCommandServer(dbs dbs.Dbs) CommandServer {
	c := make(map[string]command)
	return &commandServer{
		dbs: dbs,
		c:   c,
	}
}

func (c *commandServer) Execute(templateId int) (commandId string, err error) {
	cache := c.dbs.GetCache()
	var cacheMap map[int]model.CommandTemplate
	if x, found := cache.Get("commandTemplates"); found {
		cacheMap = x.(map[int]model.CommandTemplate)
	} else {

	}
	ct, ok := cacheMap[templateId]
	com := command{Template: model2template(ct)}
	if !ok {
		err = errors.New("can not find command template")
		return
	}
}
