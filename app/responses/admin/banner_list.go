package admin_responses

import (
	"fagin/app/models/banner"
	"fagin/pkg/response"
)

type bannerList struct {
	Ms []banner.Banner
	response.Collect
}

var _ response.IResponse = &bannerList{}

func BannerList(models ...banner.Banner) *bannerList {
	res := bannerList{Ms: models}
	res.Collect.IResponse = &res
	return &res
}

func (res *bannerList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"id":     model.ID,
			"title":  model.Title,
			"banner": model.Banner,
			"path":   model.Path,
			"sort":   model.Sort,
			"status": model.Status,
		}
		sm = append(sm, m)
	}
	return sm
}
