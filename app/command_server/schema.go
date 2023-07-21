package command_server

import (
	"encoding/json"
	"time"
)

type Protocol int

const (
	http Protocol = iota
	websocket
	mqtt
	redisTopic
)

func (p Protocol) String() string {
	return [...]string{"http", "websocket", "mqtt", "redis_topic"}[p]
}

type Column int

const (
	status Column = iota
	data
)

func (c Column) String() string {
	return [...]string{"status", "data"}[c]
}

type CalculateType int

const (
	equal CalculateType = iota
	not
	less
	greater
	lessOrEqual
	greaterOrEqual
	include
	exclude
)

func (c CalculateType) String() string {
	return [...]string{"=", "!=", "<", ">", "<=", ">=", "include", "exclude"}[c]
}

type command struct {
	CommandId   string     `json:"command_id"`
	From        time.Time  `json:"from"`
	To          *time.Time `json:"to"`
	TriggerFrom []string   `json:"trigger_from"`
	TemplateID  int        `json:"template_id"`
	Template    template   `json:"template"`
}

type PreLogicType int

const (
	and PreLogicType = iota
	or
)

func (p PreLogicType) String() string {
	return [...]string{"and", "or"}[p]
}

type template struct {
	ID          int32             `json:"id"`
	Name        string            `json:"name"`
	Protocol    string            `json:"protocol"`
	Description *string           `json:"description"`
	Host        string            `json:"host"`
	Port        string            `json:"port"`
	Http        *httpsCommand     `json:"http"`
	Mqtt        *mqttCommand      `json:"mqtt"`
	Websocket   *websocketCommand `json:"websocket"`
	Redis       *redisCommand     `json:"redis"`
	Monitor     *monitor          `json:"monitor"`
}

type httpsCommand struct {
	Method            string           `json:"method"`
	URL               string           `json:"url"`
	AuthorizationType *string          `json:"authorization_type"`
	Params            json.RawMessage  `json:"params"`
	Header            json.RawMessage  `json:"header"`
	BodyType          *string          `json:"body_type"`
	Body              *json.RawMessage `json:"body"`
}

type mqttCommand struct {
	Topic   string           `json:"topic"`
	Header  json.RawMessage  `json:"header"`
	Message *json.RawMessage `json:"message"`
	Type    string           `json:"type"`
}

type websocketCommand struct {
	URL     string          `json:"url"`
	Header  json.RawMessage `json:"header"`
	Message *string         `json:"message"`
}

type redisCommand struct {
	Password *string          `json:"password"`
	Db       *int32           `json:"db"`
	Topic    *string          `json:"topic"`
	Message  *json.RawMessage `json:"message"`
	Type     string           `json:"type"`
}

type monitor struct {
	Column            string       `json:"column"`
	Timeout           int32        `json:"timeout"`
	Interval          *int32       `json:"interval"`
	CommandTemplateID int32        `json:"command_template_id"`
	MConditions       []mCondition `json:"m_conditions"`
}

type mCondition struct {
	Order         *int32  `json:"order"`
	CalculateType *string `json:"calculate_type"`
	PreLogicType  *string `json:"pre_logic_type"`
	Value         *string `json:"value"`
	SearchRule    *string `json:"search_rule"`
	MonitorID     *int32  `json:"monitor_id"`
}
