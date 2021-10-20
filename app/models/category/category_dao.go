package category

import (
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

func New() *Category {
	return &Category{}
}

type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

func (m *Category) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

func (d *Dao) All(columns []string) (*[]Category, error) {
	var model []Category
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
	if v, ok = params["name"]; ok {
		model = model.Where("name = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *Dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.GetModel()).Error
}

func (d *Dao) ExistsByID(id uint) bool {
	return d.Query(gin.H{"id": id}, nil, nil).
		Exists()
}
