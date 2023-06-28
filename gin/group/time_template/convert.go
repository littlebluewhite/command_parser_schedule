package time_template

import "command_parser_schedule/dal/model"

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
				MConditionType:  item.TimeData.MConditionType,
				MCondition:      item.TimeData.MCondition,
			},
		}
		result = append(result, &i)
	}
	return result
}
