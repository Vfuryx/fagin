package banner

import (
	"fagin/pkg/db"
)

func New() *Banner {
	return &Banner{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (m *Banner) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = m
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(columns []string) (*[]Banner, error) {
	var model []Banner
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

	if v, ok = params["orderBy"]; ok {
		model = model.Order(v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.M).Error
}