// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameCondition = "condition"

// Condition mapped from table <condition>
type Condition struct {
	ID            int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	CalculateType *string `gorm:"column:calculate_type" json:"calculate_type"`
	LogicType     *string `gorm:"column:logic_type" json:"logic_type"`
	Value         *string `gorm:"column:value" json:"value"`
	SearchRule    *string `gorm:"column:search_rule;comment:ex: person.item.[]array.name" json:"search_rule"`
}

// TableName Condition's table name
func (*Condition) TableName() string {
	return TableNameCondition
}
