// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameHTTPSCommand = "https_command"

// HTTPSCommand mapped from table <https_command>
type HTTPSCommand struct {
	ID                int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CommandID         *int32  `gorm:"column:command_id" json:"command_id"`
	Method            *string `gorm:"column:method" json:"method"`
	URL               string  `gorm:"column:url;not null" json:"url"`
	AuthorizationType *string `gorm:"column:authorization_type" json:"authorization_type"`
	Params            *string `gorm:"column:params" json:"params"`
	Header            *string `gorm:"column:header" json:"header"`
	BodyType          *string `gorm:"column:body_type" json:"body_type"`
	Body              *string `gorm:"column:body" json:"body"`
}

// TableName HTTPSCommand's table name
func (*HTTPSCommand) TableName() string {
	return TableNameHTTPSCommand
}
