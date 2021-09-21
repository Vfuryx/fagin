package admin_permission_group

import (
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

func New() *AdminPermissionGroup {
	return &AdminPermissionGroup{}
}

type Dao struct {
	db.Dao
}

var _ db.IDao = &Dao{}

func (m *AdminPermissionGroup) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

func (d *Dao) All(params gin.H, columns []string) (*[]AdminPermissionGroup, error) {
	var model []AdminPermissionGroup
	m := db.ORM().Select(columns)
	if v, ok := params["type"]; ok {
		m = m.Where("type = ?", v)
	}
	err := m.Find(&model).Error
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
	if v, ok = params["name"]; ok {
		model = model.Where("name like ?", v)
	}
	if v, ok = params["type"]; ok {
		model = model.Where("type = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *Dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.GetModel()).Error
}
