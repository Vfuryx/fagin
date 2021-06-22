package service

import (
	"fagin/app/models/author"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type authorService struct{}

var Author authorService

func (authorService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]author.Author, error) {
	var ms []author.Author
	err := author.Dao().Query(params, columns, with).Paginator(&ms, p)
	return ms, err
}

func (authorService) Show(id uint, columns []string) (*author.Author, error) {
	b := author.New()
	err := b.Dao().FindById(id, columns)
	return b, err
}

func (authorService) Create(m *author.Author) error {
	return author.Dao().Create(m)
}

func (authorService) Update(id uint, data gin.H) error {
	return author.Dao().Update(id, data)
}

func (authorService) Delete(id uint) error {
	return author.Dao().Destroy(id)
}

func (authorService) Deletes(ids []uint) error {
	return author.Dao().Deletes(ids)
}

func (authorService) All(params gin.H, columns []string, with gin.H) ([]author.Author, error) {
	var ms []author.Author
	err := author.Dao().Query(params, columns, with).Find(&ms)
	return ms, err
}