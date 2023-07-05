package time_template

import (
	"command_parser_schedule/dal/model"
)

func Format(tt []*model.TimeTemplate) []*TimeTemplate {
	result := make([]*TimeTemplate, 0, len(tt))
	for _, item := range tt {
		i := TimeTemplate{
			ID:        item.ID,
			Name:      item.Name,
			UpdatedAt: item.UpdatedAt,
			CreatedAt: item.CreatedAt,
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

func CreateConvert(c []*TimeTemplateCreate) []*model.TimeTemplate {
	result := make([]*model.TimeTemplate, 0, len(c))
	for _, item := range c {
		i := model.TimeTemplate{
			Name: item.Name,
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

func UpdateConvert(tt []*model.TimeTemplate, uMap map[int32]*TimeTemplateUpdate) []*model.TimeTemplate {
	for i := 0; i < len(uMap); i++ {
		u := uMap[tt[i].ID]
		if u.Name != nil {
			tt[i].Name = *u.Name
		}
		if u.TimeData != nil {
			tt[i].TimeData.RepeatType = u.TimeData.RepeatType
			tt[i].TimeData.StartDate = u.TimeData.StartDate
			tt[i].TimeData.EndDate = u.TimeData.EndDate
			tt[i].TimeData.StartTime = []byte(u.TimeData.StartTime.String())
			tt[i].TimeData.EndTime = []byte(u.TimeData.EndTime.String())
			tt[i].TimeData.IntervalSeconds = u.TimeData.IntervalSeconds
			tt[i].TimeData.ConditionType = u.TimeData.ConditionType
			tt[i].TimeData.TCondition = u.TimeData.TCondition
		}
	}
	return tt
}
