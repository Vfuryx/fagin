package admin_responses

import (
	"fagin/app/models/author"
	"fagin/pkg/response"
)

type authorList struct {
	Ms []author.Author
	response.Collect
}

var _ response.IResponse = &authorList{}

func AuthorList(models ...author.Author) *authorList {
	res := authorList{Ms: models}
	res.SetCollect(&res)
	return &res
}

func (res *authorList) Serialize() []map[string]interface{} {
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
