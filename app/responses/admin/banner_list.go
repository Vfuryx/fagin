package responses

import (
	"fagin/app/models/banner"
	"fagin/pkg/response"
)

type bannerList struct {
	ms []*banner.Banner

	response.Collect
}

func NewBannerList(models ...*banner.Banner) response.Response {
	res := bannerList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *bannerList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		m := map[string]interface{}{
			"id":     res.ms[i].ID,
			"title":  res.ms[i].Title,
			"banner": res.ms[i].Banner,
			"path":   res.ms[i].Path,
			"sort":   res.ms[i].Sort,
			"status": res.ms[i].Status,
		}
		sm = append(sm, m)
	}

	return sm
}
