package admin_user_role

import (
	"fagin/pkg/db"
)

func New() *AdminUserRole {
	return &AdminUserRole{}
}

type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

func (m *AdminUserRole) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)

	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())

	return dao
}

func (d *Dao) All(columns []string) ([]*AdminUserRole, error) {
	var model []*AdminUserRole
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

	if v, ok = params["in_id"]; ok {
		model = model.Where("id in (?)", v)
	}

	if v, ok = params["role_id"]; ok {
		model = model.Where("role_id = ?", v)
	}

	if v, ok = params["in_role_id"]; ok {
		model = model.Where("role_id in (?)", v)
	}

	if v, ok = params["in_user_id"]; ok {
		model = model.Where("user_id in (?)", v)
	}

	d.DB = d.With(model, with)

	return d
}

func (d *Dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.GetModel()).Error
}

// RoleRelationExist 是否存在角色关联
func (d *Dao) RoleRelationExist(roleID uint) (bool, error) {
	var count int64

	err := db.ORM().Model(d.GetModel()).Where("role_id = ?", roleID).Count(&count).Error
	if err != nil {
		return false, err
	}

	if count > 0 {
		return true, nil
	}

	return false, nil
}
