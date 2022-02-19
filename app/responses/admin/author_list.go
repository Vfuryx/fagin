package responses

import (
	"fagin/app/models/author"
	"fagin/pkg/response"
)

type authorList struct {
	ms []*author.Author

	response.Collect
}

func NewAuthorList(models ...*author.Author) response.Response {
	res := authorList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *authorList) Serialize() []map[string]interface{} {
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
