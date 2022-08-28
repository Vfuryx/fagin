package api_responses

import (
	"fagin/app/models/video_info"
	"fagin/pkg/response"
)

type videoInfoList []video_info.VideoInfo

func NewVideoInfoList(models ...video_info.VideoInfo) *response.Collect[videoInfoList] {
	return new(response.Collect[videoInfoList]).SetCollect(models)
}

func (res videoInfoList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"id":    res[i].ID,
			"title": res[i].Title,
		}
		sm = append(sm, m)
	}

	return sm
}

type VideoInfoShow []video_info.VideoInfo

func NewVideoInfoShow(models ...video_info.VideoInfo) *response.Collect[videoInfoList] {
	return new(response.Collect[videoInfoList]).SetCollect(models)
}

func (res VideoInfoShow) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"id":    res[i].ID,
			"title": res[i].Title,
		}
		sm = append(sm, m)
	}

	return sm
}
