package admin_response

import (
	"fagin/app/models/video_info"
	"fagin/pkg/response"
)

type videoList struct {
	Ms []video_info.VideoInfo
	response.Collect
}

var _ response.Response = &videoList{}

func VideoList(models ...video_info.VideoInfo) *videoList {
	res := videoList{Ms: models}
	res.Collect.Response = &res
	return &res
}

func (res *videoList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"id":          model.ID,
			"title":       model.Title,
			"status":      model.Status,
			"path":        model.Path,
			"description": model.Description,
		}
		sm = append(sm, m)
	}
	return sm
}
