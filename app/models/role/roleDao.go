package role

import "fagin/pkg/db"

func New() *Role {
	return &Role{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (r *Role) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = r
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(columns []string) (*[]Role, error) {
	var model []Role
	err := db.ORM.Select(columns).Find(&model).Error
	return &model, err
}

func (d *dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.IDao {
	query := db.ORM.Select(columns)

	var (
		v  interface{}
		ok bool
	)
	if v, ok = params["id"]; ok {
		query = query.Where("id = ?", v)
	}

	d.DB = query
	return d
}