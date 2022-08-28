package responses

import (
	"fagin/app"
	"fagin/app/models/video_info"
	"fagin/pkg/response"
)

type videoList []video_info.VideoInfo

func NewVideoList(models ...video_info.VideoInfo) *response.Collect[videoList] {
	return new(response.Collect[videoList]).SetCollect(models)
}

func (res videoList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"id":          res[i].ID,
			"title":       res[i].Title,
			"status":      res[i].Status,
			"path":        res[i].Path,
			"description": res[i].Description,
			"duration":    res[i].Duration,
			"created_at":  app.TimeToStr(res[i].CreatedAt),
		}
		sm = append(sm, m)
	}

	return sm
}
