package responses

import (
	"fagin/app/models/tag"
	"fagin/pkg/response"
)

type tagList []tag.Tag

func NewTagList(models ...tag.Tag) *response.Collect[tagList] {
	return new(response.Collect[tagList]).SetCollect(models)
}

func (res tagList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"id":     res[i].ID,
			"name":   res[i].Name,
			"status": res[i].Status,
		}
		sm = append(sm, m)
	}

	return sm
}
