package command_server

import (
	"command_parser_schedule/entry/e_command"
	"errors"
	"sync"
)

type Protocol int

const (
	https Protocol = iota
	websocket
	mqtt
	redisTopic
)

func (p Protocol) String() string {
	return [...]string{"http", "websocket", "mqtt", "redis_topic"}[p]
}

type chs struct {
	rec chan e_command.Command
	mu  *sync.RWMutex
}

type httpHeader struct {
	Key      string `json:"key"`
	Value    string `json:"value"`
	IsActive bool   `json:"is_active"`
	DataType string `json:"data_type"`
}

type doResult struct {
	statusCode int
	respData   []byte
	status     e_command.Status
	message    string
}

type analyzeResult struct {
	getSuccess  bool
	valueResult any
	arrayResult []any
}

type assertResult struct {
	order         int32
	assertSuccess bool
	preLogicType  *string
}

var (
	valueCalculate = []string{"=", "<", "<=", ">", ">=", "!="}
	sliceCalculate = []string{"include", "exclude"}
)

type executeParams struct {
	templateId     int
	triggerFrom    []string
	triggerAccount string
	token          string
}

type SendCommand struct {
	TemplateId     int      `json:"template_id"`
	TriggerFrom    []string `json:"trigger_from"`
	TriggerAccount string   `json:"trigger_account"`
	Token          string   `json:"token"`
}

var cannotFindTemplate = errors.New("can not find Command template")
