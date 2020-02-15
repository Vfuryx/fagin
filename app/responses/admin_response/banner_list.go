package admin_response

import (
	"fagin/app/models/banner"
	"fagin/pkg/response"
)

type bannerList struct {
	Ms []banner.Banner
	response.Collect
}

var _ response.Response = &bannerList{}

func BannerList(models ...banner.Banner) *bannerList {
	res := bannerList{Ms: models}
	res.Collect.Response = &res
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
