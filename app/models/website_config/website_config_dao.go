package website_config

import (
	"fagin/pkg/db"
)

func New() *WebsiteConfig {
	return &WebsiteConfig{}
}

type Dao struct {
	db.Dao
}

var _ db.IDao = &Dao{}

func (m *WebsiteConfig) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

func (d *Dao) All(columns []string) (*[]WebsiteConfig, error) {
	var model []WebsiteConfig
	err := db.ORM().Select(columns).Find(&model).Error
	return &model, err
}

func (d *Dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.IDao {
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
