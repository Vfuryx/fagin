package responses

import (
	"fagin/app/models/author"
	"fagin/pkg/response"
)

type authorList []author.Author

func NewAuthorList(models ...author.Author) *response.Collect[authorList] {
	return new(response.Collect[authorList]).SetCollect(models)
}

func (res authorList) Serialize() []map[string]interface{} {
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
