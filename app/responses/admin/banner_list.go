package responses

import (
	"fagin/app/models/banner"
	"fagin/pkg/response"
)

type bannerList []banner.Banner

func NewBannerList(models ...banner.Banner) *response.Collect[bannerList] {
	return new(response.Collect[bannerList]).SetCollect(models)
}

func (res bannerList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"id":     res[i].ID,
			"title":  res[i].Title,
			"banner": res[i].Banner,
			"path":   res[i].Path,
			"sort":   res[i].Sort,
			"status": res[i].Status,
		}
		sm = append(sm, m)
	}

	return sm
}
