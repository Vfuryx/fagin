package responses

import (
	"fagin/app/models/category"
	"fagin/pkg/response"
)

type categoryList struct {
	ms []*category.Category

	response.Collect
}

func NewCategoryList(models ...*category.Category) response.Response {
	res := categoryList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *categoryList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		m := map[string]interface{}{
			"id":     res.ms[i].ID,
			"name":   res.ms[i].Name,
			"sort":   res.ms[i].Sort,
			"status": res.ms[i].Status,
		}
		sm = append(sm, m)
	}

	return sm
}
