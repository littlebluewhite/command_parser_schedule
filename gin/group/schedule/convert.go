package schedule

import (
	"command_parser_schedule/dal/model"
	"fmt"
)

func Format(sd []*model.Schedule) []*Schedule {
	result := make([]*Schedule, 0, len(sd))
	for _, item := range sd {
		i := Schedule{
			ID:          item.ID,
			Name:        item.Name,
			Description: item.Description,
			TaskID:      item.TaskID,
			Enabled:     item.Enabled,
			UpdatedAt:   item.UpdatedAt,
			CreatedAt:   item.CreatedAt,
			TimeData: TimeDatum{
				RepeatType:      item.TimeData.RepeatType,
				StartDate:       item.TimeData.StartDate,
				EndDate:         item.TimeData.EndDate,
				StartTime:       string(item.TimeData.StartTime),
				EndTime:         string(item.TimeData.EndTime),
				IntervalSeconds: item.TimeData.IntervalSeconds,
				ConditionType:   item.TimeData.ConditionType,
				TCondition:      item.TimeData.TCondition,
			},
		}
		result = append(result, &i)
	}
	return result
}

func CreateConvert(c []*ScheduleCreate) []*model.Schedule {
	result := make([]*model.Schedule, 0, len(c))
	for _, item := range c {
		fmt.Printf("%+v\n", item)
		i := model.Schedule{
			Name:        item.Name,
			Description: item.Description,
			TaskID:      item.TaskID,
			Enabled:     item.Enabled,
			TimeData: model.TimeDatum{
				RepeatType:      item.TimeData.RepeatType,
				StartDate:       item.TimeData.StartDate,
				EndDate:         item.TimeData.EndDate,
				StartTime:       []byte(item.TimeData.StartTime.String()),
				EndTime:         []byte(item.TimeData.EndTime.String()),
				IntervalSeconds: item.TimeData.IntervalSeconds,
				ConditionType:   item.TimeData.ConditionType,
				TCondition:      item.TimeData.TCondition,
			},
		}
		result = append(result, &i)
	}
	return result
}

func UpdateConvert(sd []*model.Schedule, uMap map[int32]*ScheduleUpdate) []*model.Schedule {
	for i := 0; i < len(uMap); i++ {
		u := uMap[sd[i].ID]
		sd[i].Name = *u.Name
		sd[i].Description = u.Description
		sd[i].TaskID = u.TaskID
		sd[i].Enabled = *u.Enabled
		if u.TimeData != nil {
			sd[i].TimeData.RepeatType = u.TimeData.RepeatType
			sd[i].TimeData.StartDate = u.TimeData.StartDate
			sd[i].TimeData.EndDate = u.TimeData.EndDate
			sd[i].TimeData.StartTime = []byte(u.TimeData.StartTime.String())
			sd[i].TimeData.EndTime = []byte(u.TimeData.EndTime.String())
			sd[i].TimeData.IntervalSeconds = u.TimeData.IntervalSeconds
			sd[i].TimeData.ConditionType = u.TimeData.ConditionType
			sd[i].TimeData.TCondition = u.TimeData.TCondition
		}
	}
	return sd
}
