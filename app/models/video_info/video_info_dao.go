package video_info

import (
	"fagin/pkg/db"
)

func New() *VideoInfo {
	return &VideoInfo{}
}

type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

func (m *VideoInfo) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

func (d *Dao) All(columns []string) (*[]VideoInfo, error) {
	var model []VideoInfo
	err := db.ORM().Select(columns).Find(&model).Error
	return &model, err
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

	if v, ok = params["author_id"]; ok {
		model = model.Where("author_id = ?", v)
	}
	if v, ok = params["title"]; ok {
		model = model.Where("Title like %?%", v)
	}
	if v, ok = params["status"]; ok {
		model = model.Where("status = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *Dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.GetModel()).Error
}
