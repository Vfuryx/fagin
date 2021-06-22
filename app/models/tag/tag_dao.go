package tag

import (
	"fagin/pkg/db"
)

func New() *Tag {
	return &Tag{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (m *Tag) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = m
	dao.DB = db.ORM()
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	dao.Dao.DB = db.ORM()
	return dao
}

func (dao) All(columns []string) (*[]Tag, error) {
	var model []Tag
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

	if v, ok = params["in_id"]; ok {
		model = model.Where("id in (?)", v)
	}
	if v, ok = params["status"]; ok {
		model = model.Where("status = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.M).Error
}
