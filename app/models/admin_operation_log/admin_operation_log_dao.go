package admin_operation_log

import (
	"fagin/pkg/db"
)

func New() *AdminOperationLog {
	return &AdminOperationLog{}
}

type Dao struct {
	db.Dao
}

var _ db.IDao = &Dao{}

func (m *AdminOperationLog) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

func (d *Dao) All(columns []string) (*[]AdminOperationLog, error) {
	var model []AdminOperationLog
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

	if v, ok = params["like_path"]; ok && v.(string) != "" {
		model = model.Where("`path` LIKE ?", v)
	}

	if v, ok = params["method"]; ok && v.(string) != "" {
		model = model.Where("`method` = ?", v)
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
