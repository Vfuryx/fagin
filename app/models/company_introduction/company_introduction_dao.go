package company_introduction

import (
	"fagin/pkg/db"
)

func New() *CompanyIntroduction {
	return &CompanyIntroduction{}
}

type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

func (m *CompanyIntroduction) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)

	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())

	return dao
}

func (d *Dao) All(columns []string) ([]*CompanyIntroduction, error) {
	var model []*CompanyIntroduction
	err := db.ORM().Select(columns).Find(&model).Error

	return model, err
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
