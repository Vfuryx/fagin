package service

import (
	"fagin/app/models/video_info"
	"fagin/pkg/db"
)

type videoInfoService struct{}

var VideoInfo videoInfoService

func (videoInfoService) CreateVideo(m *video_info.VideoInfo) error {
	return video_info.Dao().Create(m)
}

func (videoInfoService) VideoList(params map[string]interface{}, columns []string, with map[string]interface{}, p *db.Paginator) ([]video_info.VideoInfo, error) {
	var vs []video_info.VideoInfo
	err := video_info.Dao().Query(params, columns, with).Paginator(&vs, p)
	return vs, err
}

func (videoInfoService) ShowVideo(id uint, columns []string) (*video_info.VideoInfo, error) {
	v := video_info.New()
	err := v.Dao().FindById(id, columns)
	return v, err
}

func (videoInfoService) UpdateVideo(id uint, data map[string]interface{}) error {
	return video_info.Dao().Update(id, data)
}

func (videoInfoService) DeleteVideo(id uint) error {
	return video_info.Dao().Destroy(id)
}

func (videoInfoService) DeleteVideos(ids []uint) error {
	return video_info.Dao().Deletes(ids)
}

func (videoInfoService) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.IDao {
	return video_info.Dao().Query(params, columns, with)
}
