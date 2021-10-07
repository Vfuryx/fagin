package service

import (
	"fagin/app/models/banner"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type bannerService struct{}

var Banner bannerService

func (*bannerService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]banner.Banner, error) {
	var banners []banner.Banner
	err := banner.NewDao().Query(params, columns, with).Paginate(&banners, p)
	return banners, err
}

func (*bannerService) Show(id uint, columns []string) (*banner.Banner, error) {
	b := banner.New()
	err := b.Dao().FindById(id, columns)
	return b, err
}

func (*bannerService) Create(m *banner.Banner) error {
	return banner.NewDao().Create(m)
}

func (*bannerService) Update(id uint, data gin.H) error {
	return banner.NewDao().Update(id, data)
}

func (*bannerService) Delete(id uint) error {
	return banner.NewDao().Destroy(id)
}

func (*bannerService) DeleteBanners(ids []uint) error {
	return banner.NewDao().Deletes(ids)
}
