package services

import (
	"fagin/app/models/tag"
	"fagin/pkg/db"
	"fagin/pkg/errorw"

	"github.com/gin-gonic/gin"
)

type tagService struct{}

// Tag 标签
var Tag tagService

func (*tagService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) (ms []*tag.Tag, err error) {
	err = tag.NewDao().Query(params, columns, with).Paginate(&ms, p)

	return ms, errorw.UP(err)
}

func (*tagService) Show(id uint, columns []string) (*tag.Tag, error) {
	b := tag.New()
	err := b.Dao().FindByID(id, columns)

	return b, errorw.UP(err)
}

func (*tagService) Create(m *tag.Tag) error {
	err := tag.NewDao().Create(m)
	return errorw.UP(err)
}

func (*tagService) Update(id uint, data gin.H) error {
	err := tag.NewDao().Update(id, data)
	return errorw.UP(err)
}

func (*tagService) Delete(id uint) error {
	err := tag.NewDao().Destroy(id)
	return errorw.UP(err)
}

func (*tagService) Deletes(ids []uint) error {
	err := tag.NewDao().Deletes(ids)
	return errorw.UP(err)
}

func (*tagService) All(params gin.H, columns []string, with gin.H) (ms []*tag.Tag, err error) {
	err = tag.NewDao().Query(params, columns, with).Find(&ms)

	return ms, errorw.UP(err)
}
