package task_template

import (
	"command_parser_schedule/dal/model"
	"command_parser_schedule/gin/group/command_template"
	"fmt"
	"math"
)

func Format(ct []*model.TaskTemplate) []*TaskTemplate {
	result := make([]*TaskTemplate, 0, len(ct))
	for _, item := range ct {
		fmt.Printf("%+v\n", item)
		sResult := make([]*TaskStage, 0, len(item.Stages))
		for _, s := range item.Stages {
			var cTemplate *command_template.CommandTemplate
			if s.CommandTemplate != nil {
				cTemplate = command_template.Format([]*model.CommandTemplate{s.CommandTemplate})[0]
			} else {
				cTemplate = nil
			}
			i := TaskStage{
				ID:                s.ID,
				Name:              s.Name,
				StageNumber:       s.StageNumber,
				Mode:              s.Mode,
				CommandTemplateID: s.CommandTemplateID,
				Tag:               s.Tag,
				CommandTemplate:   cTemplate,
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
		sResult := make([]model.TaskStage, 0, len(item.Stages))
		for _, s := range item.Stages {
			i := model.TaskStage{
				Name:              s.Name,
				StageNumber:       s.StageNumber,
				Mode:              s.Mode,
				CommandTemplateID: s.CommandTemplateID,
				Tag:               s.Tag,
			}
			sResult = append(sResult, i)
		}
		i := model.TaskTemplate{
			Name:     item.Name,
			Variable: item.Variable,
			Stages:   sResult,
		}
		result = append(result, &i)
	}
	return result
}

func UpdateConvert(tt []*model.TaskTemplate, uMap map[int32]*TaskTemplateUpdate) []*model.TaskTemplate {
	for i := 0; i < len(uMap); i++ {
		u := uMap[tt[i].ID]
		if u.Name != nil {
			tt[i].Name = *u.Name
		}
		if u.Variable != nil {
			tt[i].Variable = *u.Variable
		}
		sId := make(map[int32]struct{})
		for _, s := range tt[i].Stages {
			sId[s.ID] = struct{}{}
		}
		if u.Stages != nil {
			sResult := make([]model.TaskStage, 0, len(u.Stages))
			for _, s := range u.Stages {
				_, ok := sId[int32(math.Abs(float64(s.ID)))]
				if !ok && s.ID != 0 {
					continue
				}
				ts := model.TaskStage{
					ID:                s.ID,
					Name:              s.Name,
					StageNumber:       s.StageNumber,
					Mode:              s.Mode,
					CommandTemplateID: s.CommandTemplateID,
					Tag:               s.Tag,
				}
				sResult = append(sResult, ts)
			}
			tt[i].Stages = sResult
		}
	}
	return tt
}
