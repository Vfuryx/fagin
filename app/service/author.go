package service

import (
	"fagin/app/models/author"
	"fagin/pkg/db"
	"fagin/pkg/errorw"
	"github.com/gin-gonic/gin"
)

type authorService struct{}

var Author authorService

func (*authorService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]author.Author, error) {
	var ms []author.Author
	err := author.NewDao().Query(params, columns, with).Paginate(&ms, p)
	return ms, errorw.UP(err)
}

func (*authorService) Show(id uint, columns []string) (*author.Author, error) {
	b := author.New()
	err := b.Dao().FindById(id, columns)
	return b, errorw.UP(err)
}

func (*authorService) Create(m *author.Author) error {
	err := author.NewDao().Create(m)
	return errorw.UP(err)
}

func (*authorService) Update(id uint, data gin.H) error {
	err := author.NewDao().Update(id, data)
	return errorw.UP(err)
}

func (*authorService) Delete(id uint) error {
	err := author.NewDao().Destroy(id)
	return errorw.UP(err)
}

func (*authorService) Deletes(ids []uint) error {
	err := author.NewDao().Deletes(ids)
	return errorw.UP(err)
}

func (*authorService) All(params gin.H, columns []string, with gin.H) ([]author.Author, error) {
	var ms []author.Author
	err := author.NewDao().Query(params, columns, with).Find(&ms)
	return ms, errorw.UP(err)
}

func (*authorService) Exists(id uint) bool {
	return author.NewDao().ExistsByID(id)
}
