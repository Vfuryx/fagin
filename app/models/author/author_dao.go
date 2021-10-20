package author

import (
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

func New() *Author {
	return &Author{}
}

type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

func (m *Author) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

func (d *Dao) All(columns []string) (*[]Author, error) {
	var model []Author
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
	if v, ok = params["status"]; ok {
		model = model.Where("status = ?", v)
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
