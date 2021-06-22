package admin_responses

import (
	"fagin/app"
	"fagin/app/models/video_info"
	"fagin/pkg/response"
)

type videoList struct {
	Ms []video_info.VideoInfo
	response.Collect
}

var _ response.IResponse = &videoList{}

func VideoList(models ...video_info.VideoInfo) *videoList {
	res := videoList{Ms: models}
	res.Collect.IResponse = &res
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
			"duration":    model.Duration,
			"created_at":  model.CreatedAt.Format(app.TimeFormat),
		}
		sm = append(sm, m)
	}
	return sm
}
