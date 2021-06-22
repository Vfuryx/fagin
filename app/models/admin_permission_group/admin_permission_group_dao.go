package admin_permission_group

import (
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

func New() *AdminPermissionGroup {
	return &AdminPermissionGroup{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (m *AdminPermissionGroup) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = m
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(params gin.H, columns []string) (*[]AdminPermissionGroup, error) {
	var model []AdminPermissionGroup
	m := db.ORM().Select(columns)
	if v, ok := params["type"]; ok {
		m = m.Where("type = ?", v)
	}
	err := m.Find(&model).Error
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
	if v, ok = params["name"]; ok {
		model = model.Where("name like ?", v)
	}
	if v, ok = params["type"]; ok {
		model = model.Where("type = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.M).Error
}
