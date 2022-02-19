package responses

import (
	"fagin/app"
	"fagin/app/models/video_info"
	"fagin/pkg/response"
)

type videoList struct {
	ms []*video_info.VideoInfo

	response.Collect
}

func NewVideoList(models ...*video_info.VideoInfo) response.Response {
	res := videoList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *videoList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		m := map[string]interface{}{
			"id":          res.ms[i].ID,
			"title":       res.ms[i].Title,
			"status":      res.ms[i].Status,
			"path":        res.ms[i].Path,
			"description": res.ms[i].Description,
			"duration":    res.ms[i].Duration,
			"created_at":  app.TimeToStr(res.ms[i].CreatedAt),
		}
		sm = append(sm, m)
	}

	return sm
}
