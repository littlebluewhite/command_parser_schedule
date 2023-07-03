package header_template

import "command_parser_schedule/dal/model"

func CreateConvert(c []*HeaderTemplateCreate) []*model.HeaderTemplate {
	result := make([]*model.HeaderTemplate, 0, len(c))
	for _, item := range c {
		i := model.HeaderTemplate{
			Name: item.Name,
			Data: item.Data,
		}
		result = append(result, &i)
	}
	return result
}

func UpdateConvert(ht []*model.HeaderTemplate, uMap map[int32]*HeaderTemplateUpdate) []*model.HeaderTemplate {
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
