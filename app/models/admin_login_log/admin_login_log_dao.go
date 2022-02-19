package admin_login_log

import (
	"fagin/pkg/db"
)

// New 实例化
func New() *AdminLoginLog {
	return &AdminLoginLog{}
}

// Dao DAO
type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

// Dao DAO
func (m *AdminLoginLog) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)

	return dao
}

// NewDao 实例化DAO
func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())

	return dao
}

// All 所有
func (d *Dao) All(columns []string) ([]*AdminLoginLog, error) {
	var model []*AdminLoginLog
	err := db.ORM().Select(columns).Find(&model).Error

	return model, err
}

// Query 查询
func (d *Dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.DAO {
	model := db.ORM().Select(columns)

	var (
		v  interface{}
		ok bool
	)

	if v, ok = params["id"]; ok {
		model = model.Where("id = ?", v)
	}

	if v, ok = params["like_login_name"]; ok && v.(string) != "" {
		model = model.Where("`login_name` LIKE ?", v)
	}

	if v, ok = params["start_time"]; ok {
		model = model.Where("`created_at` >= ?", v)
	}

	if v, ok = params["end_time"]; ok {
		model = model.Where("`created_at` < ?", v)
	}

	if v, ok = params["orderBy"]; ok {
		model = model.Order(v)
	}

	d.DB = d.With(model, with)

	return d
}

// Deletes 批量删除
func (d *Dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.GetModel()).Error
}
