package service

import (
	"fagin/app/models/tag"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type tagService struct{}

var Tag tagService

func (tagService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) (ms []tag.Tag, err error) {
	err = tag.Dao().Query(params, columns, with).Paginator(&ms, p)
	return
}

func (tagService) Show(id uint, columns []string) (*tag.Tag, error) {
	b := tag.New()
	err := b.Dao().FindById(id, columns)
	return b, err
}

func (tagService) Create(m *tag.Tag) error {
	return tag.Dao().Create(m)
}

func (tagService) Update(id uint, data gin.H) error {
	return tag.Dao().Update(id, data)
}

func (tagService) Delete(id uint) error {
	return tag.Dao().Destroy(id)
}

func (tagService) Deletes(ids []uint) error {
	return tag.Dao().Deletes(ids)
}


func (tagService) All(params gin.H, columns []string, with gin.H) (ms []tag.Tag, err error) {
	err = tag.Dao().Query(params, columns, with).Find(&ms)
	return
}