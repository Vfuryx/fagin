package admin_user

import (
	"fagin/pkg/db"
)

func New() *AdminUser {
	return &AdminUser{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (m *AdminUser) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = m
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(columns []string) (*[]AdminUser, error) {
	var model []AdminUser
	err := db.ORM.Select(columns).Find(&model).Error
	return &model, err
}

func (d *dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.IDao {
	model := db.ORM.Select(columns)

	var (
		v  interface{}
		ok bool
	)
	if v, ok = params["id"]; ok {
		model = model.Where("id = ?", v)
	}

	d.DB = model
	return d
}

func (d *dao) Deletes(ids []uint) error {
	var menu AdminUser
	return db.ORM.Where("id in (?)", ids).Delete(&menu).Error
}
