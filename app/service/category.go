package service

import (
	"fagin/app/models/article"
	"fagin/app/models/category"
	"fagin/pkg/db"
	"fagin/pkg/errorw"
	"github.com/gin-gonic/gin"
)

type categoryService struct{}

// Category 分类服务
var Category categoryService

func (*categoryService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]category.Category, error) {
	var ms []category.Category
	err := article.NewDao().Query(params, columns, with).Paginate(&ms, p)
	return ms, errorw.UP(err)
}

func (*categoryService) Show(id uint, columns []string) (*category.Category, error) {
	b := category.New()
	err := b.Dao().FindById(id, columns)
	return b, errorw.UP(err)
}

func (*categoryService) Create(m *category.Category) error {
	err := category.NewDao().Create(m)
	return errorw.UP(err)
}

func (*categoryService) Update(id uint, data gin.H) error {
	err := category.NewDao().Update(id, data)
	return errorw.UP(err)
}

func (*categoryService) Delete(id uint) error {
	err := category.NewDao().Destroy(id)
	return errorw.UP(err)
}

func (*categoryService) Deletes(ids []uint) error {
	err := category.NewDao().Deletes(ids)
	return errorw.UP(err)
}

func (*categoryService) All(params gin.H, columns []string, with gin.H) ([]category.Category, error) {
	var ms []category.Category
	err := category.NewDao().Query(params, columns, with).Find(&ms)
	return ms, errorw.UP(err)
}

func (*categoryService) Exists(id uint) bool {
	return category.NewDao().ExistsByID(id)
}
