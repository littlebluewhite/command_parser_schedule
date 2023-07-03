package command_template

import (
	"encoding/json"
)

type CommandTemplate struct {
	ID   int32           `json:"id"`
	Name string          `json:"name"`
	Data json.RawMessage `json:"data"`
}

type CommandTemplateCreate struct {
	Name string          `json:"name" binding:"required"`
	Data json.RawMessage `json:"data" binding:"required"`
}

type CommandTemplateUpdate struct {
	ID   int32            `json:"id" binding:"required"`
	Name *string          `json:"name"`
	Data *json.RawMessage `json:"data"`
}
