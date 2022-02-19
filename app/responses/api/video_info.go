package api_responses

import (
	"fagin/app/models/video_info"
	"fagin/pkg/response"
)

type videoInfoList struct {
	ms []*video_info.VideoInfo

	response.Collect
}

func NewVideoInfoList(models ...*video_info.VideoInfo) response.Response {
	res := videoInfoList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *videoInfoList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		m := map[string]interface{}{
			"id":    res.ms[i].ID,
			"title": res.ms[i].Title,
		}
		sm = append(sm, m)
	}

	return sm
}

type VideoInfoShow struct {
	ms []*video_info.VideoInfo

	response.Collect
}

func NewVideoInfoShow(models ...*video_info.VideoInfo) response.Response {
	res := VideoInfoShow{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *VideoInfoShow) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		m := map[string]interface{}{
			"id":    res.ms[i].ID,
			"title": res.ms[i].Title,
		}
		sm = append(sm, m)
	}

	return sm
}
