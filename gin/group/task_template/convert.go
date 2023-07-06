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
		mResult := make([]*model.MCondition, 0, len(item.Monitor.MConditions))
		for _, m := range item.Monitor.MConditions {
			i := model.MCondition{
				Order:         m.Order,
				CalculateType: m.CalculateType,
				PreLogicType:  m.PreLogicType,
				Value:         m.Value,
				SearchRule:    m.SearchRule,
			}
			mResult = append(mResult, &i)
		}
		i := model.TaskTemplate{
			Name:        item.Name,
			Protocol:    item.Protocol,
			Description: item.Description,
			Host:        item.Host,
			Port:        item.Port,
		}
		if item.Http != nil {
			i.Http = &model.HTTPSTask{
				Method:            item.Http.Method,
				URL:               item.Http.URL,
				AuthorizationType: item.Http.AuthorizationType,
				Params:            item.Http.Params,
				Header:            item.Http.Header,
				BodyType:          item.Http.BodyType,
				Body:              item.Http.Body,
			}
		}
		if item.Mqtt != nil {
			i.Mqtt = &model.MqttTask{
				Topic:   item.Mqtt.Topic,
				Header:  item.Mqtt.Header,
				Message: item.Mqtt.Message,
				Type:    item.Mqtt.Type,
			}
		}
		if item.Websocket != nil {
			i.Websocket = &model.WebsocketTask{
				URL:     item.Websocket.URL,
				Header:  item.Websocket.Header,
				Message: item.Websocket.Message,
			}
		}
		if item.Redis != nil {
			i.Redis = &model.RedisTask{
				Password: item.Redis.Password,
				Db:       item.Redis.Db,
				Topic:    item.Redis.Topic,
				Message:  item.Redis.Message,
				Type:     item.Redis.Type,
			}
		}
		if item.Monitor != nil {
			i.Monitor = &model.Monitor{
				Column:      item.Monitor.Column,
				Timeout:     item.Monitor.Timeout,
				Interval:    item.Monitor.Interval,
				MConditions: mResult,
			}
		}
		result = append(result, &i)
	}
	return result
}
