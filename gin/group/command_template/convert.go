package command_template

import "command_parser_schedule/dal/model"

func CreateConvert(c []*CommandTemplateCreate) []*model.CommandTemplate {
	result := make([]*model.CommandTemplate, 0, len(c))
	for _, item := range c {
		i := model.CommandTemplate{
			Name: item.Name,
			Data: item.Data,
		}
		result = append(result, &i)
	}
	return result
}

func UpdateConvert(ht []*model.CommandTemplate, uMap map[int32]*CommandTemplateUpdate) []*model.CommandTemplate {
	for i := 0; i < len(uMap); i++ {
		u := uMap[ht[i].ID]
		if u.Name != nil {
			ht[i].Name = *u.Name
		}
		if u.Data != nil {
			ht[i].Data = *u.Data
		}
	}
	return ht
}
