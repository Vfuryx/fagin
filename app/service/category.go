package service

import (
	"fagin/app/models/article"
	"fagin/app/models/category"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type categoryService struct{}

var Category categoryService

func (*categoryService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]category.Category, error) {
	var ms []category.Category
	err := article.NewDao().Query(params, columns, with).Paginate(&ms, p)
	return ms, err
}

func (*categoryService) Show(id uint, columns []string) (*category.Category, error) {
	b := category.New()
	err := b.Dao().FindById(id, columns)
	return b, err
}

func (*categoryService) Create(m *category.Category) error {
	return category.NewDao().Create(m)
}

func (*categoryService) Update(id uint, data gin.H) error {
	return category.NewDao().Update(id, data)
}

func (*categoryService) Delete(id uint) error {
	return category.NewDao().Destroy(id)
}

func (*categoryService) Deletes(ids []uint) error {
	return category.NewDao().Deletes(ids)
}

func (*categoryService) All(params gin.H, columns []string, with gin.H) ([]category.Category, error) {
	var ms []category.Category
	err := category.NewDao().Query(params, columns, with).Find(&ms)
	return ms, err
}
