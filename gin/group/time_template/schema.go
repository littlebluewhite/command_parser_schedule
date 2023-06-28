package time_template

import (
	"encoding/json"
	"time"
)

type TimeTemplate struct {
	ID        int32      `json:"id"`
	Name      string     `json:"name"`
	UpdatedAt *time.Time `json:"updated_at"`
	CreatedAt *time.Time `json:"created_at"`
	TimeData  TimeDatum  `json:"time_data"`
}

type TimeDatum struct {
	RepeatType      *string         `json:"repeat_type"`
	StartDate       time.Time       `json:"start_date"`
	EndDate         *time.Time      `json:"end_date"`
	StartTime       string          `json:"start_time"`
	EndTime         string          `json:"end_time"`
	IntervalSeconds *int32          `json:"interval_seconds"`
	MConditionType  *string         `json:"m_condition_type"`
	MCondition      json.RawMessage `json:"m_condition"`
}

type TimeTemplateCreate struct {
	Name     string          `json:"name" binding:"required"`
	TimeData TimeDatumCreate `json:"time_data" binding:"required"`
}

type TimeDatumCreate struct {
	RepeatType      *string         `json:"repeat_type"`
	StartDate       time.Time       `json:"start_date" binding:"required"`
	EndDate         *time.Time      `json:"end_date"`
	StartTime       string          `json:"start_time" binding:"required"`
	EndTime         string          `json:"end_time" binding:"required"`
	IntervalSeconds *int32          `json:"interval_seconds"`
	MConditionType  *string         `json:"m_condition_type"`
	MCondition      json.RawMessage `json:"m_condition" binding:"required"`
}
