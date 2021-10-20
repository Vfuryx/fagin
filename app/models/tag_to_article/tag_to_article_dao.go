package tag_to_article

import (
	"fagin/pkg/db"
)

func New() *TagToArticle {
	return &TagToArticle{}
}

type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

func (m *TagToArticle) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

func (d *Dao) All(columns []string) (*[]TagToArticle, error) {
	var model []TagToArticle
	err := db.ORM().Select(columns).Find(&model).Error
	return &model, err
}

func (d *Dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.DAO {
	model := db.ORM().Select(columns)

	var (
		v  interface{}
		ok bool
	)
	if v, ok = params["id"]; ok {
		model = model.Where("id = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}
