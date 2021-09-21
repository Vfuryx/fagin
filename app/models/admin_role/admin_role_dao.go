package admin_role

import (
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_permission"
	"fagin/app/models/admin_role_menu"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

func New() *AdminRole {
	return &AdminRole{}
}

type Dao struct {
	db.Dao
}

var _ db.IDao = &Dao{}

func (m *AdminRole) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

func (d *Dao) All(columns []string) (*[]AdminRole, error) {
	var model []AdminRole
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

	if v, ok = params["in_id"]; ok {
		model = model.Where("id in (?)", v)
	}

	if v, ok = params["like_name"]; ok && v.(string) != "" {
		model = model.Where("`name` LIKE ?", v)
	}

	if v, ok = params["like_key"]; ok && v.(string) != "" {
		model = model.Where("`key` LIKE ?", v)
	}

	if v, ok = params["key"]; ok && v.(string) != "" {
		model = model.Where("`key` = ?", v)
	}

	if v, ok = params["status"]; ok {
		model = model.Where("`status` = ?", v)
	}

	if v, ok = params["type"]; ok {
		model = model.Where("`type` = ?", v)
	}

	if v, ok = params["orderBy"]; ok {
		model = model.Order(v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *Dao) Show(id uint, columns []string) (*AdminRole, error) {
	var role AdminRole
	err := db.ORM().Select(columns).Where("id = ?", id).
		Preload("Menus").Preload("Permissions").First(&role).Error
	return &role, err
}

func (d *Dao) Update(id uint, data map[string]interface{}) error {
	err := db.ORM().Model(d.GetModel()).Omit("menuIDs", "permissionsIDs").Where("id = ?", id).Updates(data).Error
	if err != nil {
		return err
	}

	// 获取菜单组
	var menus []admin_menu.AdminMenu
	if v, ok := data["menuIDs"]; ok {
		ids := v.([]uint)
		err := admin_menu.NewDao().Query(gin.H{"in_id": ids}, []string{"id"}, nil).Find(&menus)
		if err != nil {
			return err
		}
		delete(data, "menuIDs")
	}
	// 获取权限组
	var permissions []admin_permission.AdminPermission
	if v, ok := data["permissionsIDs"]; ok {
		ids := v.([]uint)
		err := admin_permission.NewDao().Query(gin.H{"in_id": ids}, []string{"id"}, nil).Find(&permissions)
		if err != nil {
			return err
		}
		delete(data, "permissionsIDs")
	}
	err = db.ORM().Model(&AdminRole{ID: id}).Association("Menus").Replace(menus)
	err = db.ORM().Model(&AdminRole{ID: id}).Association("Permissions").Replace(permissions)
	return err
}

func (d *Dao) Delete(id uint) error {
	var role AdminRole
	err := db.ORM().Where("id = ?", id).Delete(&role).Error
	if err != nil {
		return err
	}
	var roleMenu admin_role_menu.AdminRoleMenu
	return db.ORM().Where("role_id = ?", id).Delete(&roleMenu).Error
}

func (d *Dao) Deletes(ids []uint) error {
	var role AdminRole
	err := db.ORM().Where("id in (?)", ids).Delete(&role).Error
	if err != nil {
		return err
	}
	var roleMenu admin_role_menu.AdminRoleMenu
	return db.ORM().Where("role_id in (?)", ids).Delete(&roleMenu).Error
}
