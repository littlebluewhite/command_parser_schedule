package task_template

import (
	"command_parser_schedule/dal/model"
	"command_parser_schedule/gin/group/command_template"
)

func Format(ct []*model.TaskTemplate) []*TaskTemplate {
	result := make([]*TaskTemplate, 0, len(ct))
	for _, item := range ct {
		sResult := make([]*TaskStage, 0, len(item.Stages))
		for _, s := range item.Stages {
			i := TaskStage{
				Name:              s.Name,
				StageNumber:       s.StageNumber,
				Mode:              s.Mode,
				CommandTemplateID: s.CommandTemplateID,
				Tag:               s.Tag,
				CommandTemplate:   *command_template.Format([]*model.CommandTemplate{&s.CommandTemplate})[0],
			}
			sResult = append(sResult, &i)
		}
		i := TaskTemplate{
			ID:        item.ID,
			Name:      item.Name,
			Variable:  item.Variable,
			UpdatedAt: item.UpdatedAt,
			CreatedAt: item.CreatedAt,
			Stages:    sResult,
		}
		result = append(result, &i)
	}
	return result
}

func CreateConvert(c []*TaskTemplateCreate) []*model.TaskTemplate {
	result := make([]*model.TaskTemplate, 0, len(c))
	for _, item := range c {
		sResult := make([]*model.TaskStage, 0, len(item.Stages))
		for _, s := range item.Stages {
			i := model.TaskStage{
				Name:              s.Name,
				StageNumber:       s.StageNumber,
				Mode:              s.Mode,
				CommandTemplateID: s.CommandTemplateID,
				Tag:               s.Tag,
			}
			sResult = append(sResult, &i)
		}
		i := model.TaskTemplate{
			ID:       item.ID,
			Name:     item.Name,
			Variable: item.Variable,
			Stages:   sResult,
		}
		result = append(result, &i)
	}
	return result
}
