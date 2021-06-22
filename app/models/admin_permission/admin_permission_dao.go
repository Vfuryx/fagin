package admin_permission

import (
	"fagin/pkg/db"
)

func New() *AdminPermission {
	return &AdminPermission{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (m *AdminPermission) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = m
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(columns []string) (*[]AdminPermission, error) {
	var model []AdminPermission
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
	if v, ok = params["gid"]; ok {
		model = model.Where("gid = ?", v)
	}
	if v, ok = params["type"]; ok {
		model = model.Where("type = ?", v)
	}

	if v, ok = params["path"]; ok {
		model = model.Where("path = ?", v)
	}

	if v, ok = params["method"]; ok {
		model = model.Where("method = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.M).Error
}

// PermissionRelationExist 检查权限是否还存在关联
func (d *dao) PermissionRelationExist(permissionID uint) (bool, error) {
	var count int64
	err := db.ORM().Table("admin_role_permissions").Where("permission_id = ?", permissionID).Count(&count).Error
	if err != nil {
		return false, err
	}
	if count > 0 {
		return true, nil
	}
	return false, nil
}
