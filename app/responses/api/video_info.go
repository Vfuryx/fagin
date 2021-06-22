package api_responses

import (
	"fagin/app/models/video_info"
	"fagin/pkg/response"
)

type videoInfoList struct {
	Ms []video_info.VideoInfo
	response.Collect
}

var _ response.IResponse = &videoInfoList{}

func VideoInfoList(models ...video_info.VideoInfo) *videoInfoList {
	res := videoInfoList{Ms: models}
	res.Collect.IResponse = &res
	return &res
}

func (res *videoInfoList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"id":    model.ID,
			"title": model.Title,
		}
		sm = append(sm, m)
	}
	return sm
}

type videoInfoShow struct {
	Ms []video_info.VideoInfo
	response.Collect
}

var _ response.IResponse = &videoInfoShow{}

func VideoInfoShow(models ...video_info.VideoInfo) *videoInfoShow {
	res := videoInfoShow{Ms: models}
	res.Collect.IResponse = &res
	return &res
}

func (res *videoInfoShow) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		m := map[string]interface{}{
			"id":    model.ID,
			"title": model.Title,
		}
		sm = append(sm, m)
	}
	return sm
}
