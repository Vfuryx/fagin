package admin_role_menu

import (
	"fagin/pkg/db"
)

func New() *AdminRoleMenu {
	return &AdminRoleMenu{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (m *AdminRoleMenu) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = m
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(columns []string) (*[]AdminRoleMenu, error) {
	var model []AdminRoleMenu
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
	if v, ok = params["menu_id"]; ok {
		model = model.Where("menu_id = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

// MenuRelationExist 是否存在菜单关联
func (d *dao) MenuRelationExist(menuID uint) (bool, error) {
	var count int64
	err := db.ORM().Model(d.M).Where("menu_id = ?", menuID).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
