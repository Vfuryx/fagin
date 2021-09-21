package admin_role_menu

import (
	"fagin/pkg/db"
)

func New() *AdminRoleMenu {
	return &AdminRoleMenu{}
}

type Dao struct {
	db.Dao
}

var _ db.IDao = &Dao{}

func (m *AdminRoleMenu) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

func (d *Dao) All(columns []string) (*[]AdminRoleMenu, error) {
	var model []AdminRoleMenu
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
	if v, ok = params["menu_id"]; ok {
		model = model.Where("menu_id = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

// MenuRelationExist 是否存在菜单关联
func (d *Dao) MenuRelationExist(menuID uint) (bool, error) {
	var count int64
	err := db.ORM().Model(d.GetModel()).Where("menu_id = ?", menuID).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
