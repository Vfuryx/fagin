package admin_responses

import (
	"fagin/app/models/tag"
	"fagin/pkg/response"
)

type tagList struct {
	Ms []tag.Tag
	response.Collect
}

var _ response.IResponse = &tagList{}

func TagList(models ...tag.Tag) *tagList {
	res := tagList{Ms: models}
	res.SetCollect(&res)
	return &res
}

func (res *tagList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"id":     model.ID,
			"name":   model.Name,
			"status": model.Status,
		}
		sm = append(sm, m)
	}
	return sm
}
