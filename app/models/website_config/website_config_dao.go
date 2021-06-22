package website_config

import (
	"fagin/pkg/db"
)

func New() *WebsiteConfig {
	return &WebsiteConfig{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (m *WebsiteConfig) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = m
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(columns []string) (*[]WebsiteConfig, error) {
	var model []WebsiteConfig
	err := db.ORM().Select(columns).Find(&model).Error
	return &model, err
}

func (d *dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.IDao {
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
