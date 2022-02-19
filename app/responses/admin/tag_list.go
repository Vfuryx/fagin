package responses

import (
	"fagin/app/models/tag"
	"fagin/pkg/response"
)

type tagList struct {
	ms []*tag.Tag

	response.Collect
}

func NewTagList(models ...*tag.Tag) response.Response {
	res := tagList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *tagList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		m := map[string]interface{}{
			"id":     res.ms[i].ID,
			"name":   res.ms[i].Name,
			"status": res.ms[i].Status,
		}
		sm = append(sm, m)
	}

	return sm
}
