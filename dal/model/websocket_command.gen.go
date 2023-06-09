// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

import (
	"encoding/json"
)

const TableNameWebsocketCommand = "websocket_command"

// WebsocketCommand mapped from table <websocket_command>
type WebsocketCommand struct {
	ID                int32           `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CommandTemplateID *int32          `gorm:"column:command_template_id" json:"command_template_id"`
	URL               string          `gorm:"column:url;not null" json:"url"`
	Header            json.RawMessage `gorm:"column:header;default:json_array()" json:"header"`
	Message           *string         `gorm:"column:message" json:"message"`
}

// TableName WebsocketCommand's table name
func (*WebsocketCommand) TableName() string {
	return TableNameWebsocketCommand
}
