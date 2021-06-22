package service

import (
	"fagin/app/models/article"
	"fagin/app/models/category"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type categoryService struct{}

var Category categoryService

func (categoryService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]category.Category, error) {
	var ms []category.Category
	err := article.Dao().Query(params, columns, with).Paginator(&ms, p)
	return ms, err
}

func (categoryService) Show(id uint, columns []string) (*category.Category, error) {
	b := category.New()
	err := b.Dao().FindById(id, columns)
	return b, err
}

func (categoryService) Create(m *category.Category) error {
	return category.Dao().Create(m)
}

func (categoryService) Update(id uint, data gin.H) error {
	return category.Dao().Update(id, data)
}

func (categoryService) Delete(id uint) error {
	return category.Dao().Destroy(id)
}

func (categoryService) Deletes(ids []uint) error {
	return category.Dao().Deletes(ids)
}

func (categoryService) All(params gin.H, columns []string, with gin.H) ([]category.Category, error) {
	var ms []category.Category
	err := category.Dao().Query(params, columns, with).Find(&ms)
	return ms, err
}
