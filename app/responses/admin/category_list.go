package admin_responses

import (
	"fagin/app/models/category"
	"fagin/pkg/response"
)

type categoryList struct {
	Ms []category.Category
	response.Collect
}

var _ response.Response = &categoryList{}

func CategoryList(models ...category.Category) *categoryList {
	res := categoryList{Ms: models}
	res.SetCollect(&res)
	return &res
}

func (res *categoryList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"id":     model.ID,
			"name":   model.Name,
			"sort":   model.Sort,
			"status": model.Status,
		}
		sm = append(sm, m)
	}
	return sm
}
