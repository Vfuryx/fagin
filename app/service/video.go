package service

import (
	"fagin/app/models/video_info"
	"fagin/pkg/db"
	"fagin/pkg/errorw"
)

type videoInfoService struct{}

var VideoInfo videoInfoService

func (*videoInfoService) CreateVideo(m *video_info.VideoInfo) error {
	err := video_info.NewDao().Create(m)
	return errorw.UP(err)
}

func (*videoInfoService) VideoList(params map[string]interface{}, columns []string, with map[string]interface{}, p *db.Paginator) ([]video_info.VideoInfo, error) {
	var vs []video_info.VideoInfo
	err := video_info.NewDao().Query(params, columns, with).Paginate(&vs, p)
	return vs, errorw.UP(err)
}

func (*videoInfoService) ShowVideo(id uint, columns []string) (*video_info.VideoInfo, error) {
	v := video_info.New()
	err := v.Dao().FindById(id, columns)
	return v, errorw.UP(err)
}

func (*videoInfoService) UpdateVideo(id uint, data map[string]interface{}) error {
	err := video_info.NewDao().Update(id, data)
	return errorw.UP(err)
}

func (*videoInfoService) DeleteVideo(id uint) error {
	err := video_info.NewDao().Destroy(id)
	return errorw.UP(err)
}

func (*videoInfoService) DeleteVideos(ids []uint) error {
	err := video_info.NewDao().Deletes(ids)
	return errorw.UP(err)
}

func (*videoInfoService) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.DAO {
	return video_info.NewDao().Query(params, columns, with)
}
