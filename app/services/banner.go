package services

import (
	"fagin/app/models/banner"
	"fagin/pkg/db"
	"fagin/pkg/errorw"

	"github.com/gin-gonic/gin"
)

type bannerService struct{}

// Banner Banner服务
var Banner bannerService

func (*bannerService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]*banner.Banner, error) {
	var banners []*banner.Banner
	err := banner.NewDao().Query(params, columns, with).Paginate(&banners, p)

	return banners, errorw.UP(err)
}

func (*bannerService) Show(id uint, columns []string) (*banner.Banner, error) {
	b := banner.New()
	err := b.Dao().FindByID(id, columns)

	return b, errorw.UP(err)
}

func (*bannerService) Create(m *banner.Banner) error {
	err := banner.NewDao().Create(m)
	return errorw.UP(err)
}

func (*bannerService) Update(id uint, data gin.H) error {
	err := banner.NewDao().Update(id, data)
	return errorw.UP(err)
}

func (*bannerService) Delete(id uint) error {
	err := banner.NewDao().Destroy(id)
	return errorw.UP(err)
}

func (*bannerService) DeleteBanners(ids []uint) error {
	err := banner.NewDao().Deletes(ids)
	return errorw.UP(err)
}
