package responses

import (
	"fagin/app/models/category"
	"fagin/pkg/response"
)

type categoryList []category.Category

func NewCategoryList(models ...category.Category) *response.Collect[categoryList] {
	return new(response.Collect[categoryList]).SetCollect(models)
}

func (res categoryList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"id":     res[i].ID,
			"name":   res[i].Name,
			"sort":   res[i].Sort,
			"status": res[i].Status,
		}
		sm = append(sm, m)
	}

	return sm
}
