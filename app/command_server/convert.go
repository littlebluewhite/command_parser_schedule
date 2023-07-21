package command_server

import "command_parser_schedule/dal/model"

func model2template(mc model.CommandTemplate) template {
	t := template{
		ID:          mc.ID,
		Name:        mc.Name,
		Protocol:    mc.Protocol,
		Description: mc.Description,
		Host:        mc.Host,
		Port:        mc.Port,
	}
	if mc.Http != nil {
		t.Http = &httpsCommand{
			Method:            mc.Http.Method,
			URL:               mc.Http.URL,
			AuthorizationType: mc.Http.AuthorizationType,
			Params:            mc.Http.Params,
			Header:            mc.Http.Header,
			BodyType:          mc.Http.BodyType,
			Body:              mc.Http.Body,
		}
	}
	if mc.Mqtt != nil {
		t.Mqtt = &mqttCommand{
			Topic:   mc.Mqtt.Topic,
			Header:  mc.Mqtt.Header,
			Message: mc.Mqtt.Message,
			Type:    mc.Mqtt.Type,
		}
	}
	if mc.Websocket != nil {
		t.Websocket = &websocketCommand{
			URL:     mc.Websocket.URL,
			Header:  mc.Websocket.Header,
			Message: mc.Websocket.Message,
		}
	}
	if mc.Redis != nil {
		t.Redis = &redisCommand{
			Password: mc.Redis.Password,
			Db:       mc.Redis.Db,
			Topic:    mc.Redis.Topic,
			Message:  mc.Redis.Message,
			Type:     mc.Redis.Type,
		}
	}
	if mc.Monitor != nil {
		mResult := make([]mCondition, 0, len(mc.Monitor.MConditions))
		for _, m := range mc.Monitor.MConditions {
			i := mCondition{
				Order:         m.Order,
				CalculateType: m.CalculateType,
				PreLogicType:  m.PreLogicType,
				Value:         m.Value,
				SearchRule:    m.SearchRule,
			}
			mResult = append(mResult, i)
		}
		t.Monitor = &monitor{
			Column:      mc.Monitor.Column,
			Timeout:     mc.Monitor.Timeout,
			Interval:    mc.Monitor.Interval,
			MConditions: mResult,
		}
	}
	return t
}
