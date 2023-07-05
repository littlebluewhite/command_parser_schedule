// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.
// Code generated by gorm.io/gen. DO NOT EDIT.

package model

const TableNameMCondition = "m_condition"

// MCondition mapped from table <m_condition>
type MCondition struct {
	ID            int32   `gorm:"column:id;primaryKey;autoIncrement:true" json:"id"`
	Order         *int32  `gorm:"column:order" json:"order"`
	CalculateType *string `gorm:"column:calculate_type" json:"calculate_type"`
	PreLogicType  *string `gorm:"column:pre_logic_type" json:"pre_logic_type"`
	Value         *string `gorm:"column:value" json:"value"`
	SearchRule    *string `gorm:"column:search_rule;comment:ex: person.item.[]array.name" json:"search_rule"`
	MonitorID     *int32  `gorm:"column:monitor_id" json:"monitor_id"`
}

// TableName MCondition's table name
func (*MCondition) TableName() string {
	return TableNameMCondition
}
