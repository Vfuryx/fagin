package article

import (
	"fagin/app/models/tag"
	"fagin/pkg/db"
)

// New 实例
func New() *Article {
	return &Article{}
}

// Dao New 实例
type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

// Dao New 实例
func (m *Article) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

// NewDao New 实例
func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

// All New 实例
func (d *Dao) All(columns []string) (*[]Article, error) {
	var model []Article
	err := db.ORM().Select(columns).Find(&model).Error
	return &model, err
}

// Query New 实例
func (d *Dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.DAO {
	model := db.ORM().Select(columns)

	var (
		v  interface{}
		ok bool
	)
	if v, ok = params["id"]; ok {
		model = model.Where("id = ?", v)
	}

	if v, ok = params["id_in"]; ok {
		model = model.Where("id in (?)", v)
	}

	if v, ok = params["status"]; ok {
		model = model.Where("status = ?", v)
	}

	if v, ok = params["category_id"]; ok {
		model = model.Where("category_id = ?", v)
	}

	if v, ok = params["orderBy"]; ok {
		model = model.Order(v)
	}

	d.DB = d.With(model, with)
	return d
}

// Deletes New 实例
func (d *Dao) Deletes(ids []uint) error {
	var a Article
	return db.ORM().Where("id in (?)", ids).Delete(&a).Error
}

// Update New 实例
func (d *Dao) Update(id uint, data map[string]interface{}) error {
	err := d.FindById(id, []string{"id"})
	if err != nil {
		return err
	}
	if v, ok := data["Tags"]; ok {
		// 有点复杂后期修改
		// 清空
		_ = db.ORM().Model(&Article{ID: id}).Association("Tags").Clear()
		// 插入
		_ = db.ORM().Model(&Article{ID: id}).Association("Tags").Append(v.([]tag.Tag))
		delete(data, "Tags")
	}
	return db.ORM().Model(d.GetModel()).Where("id = ?", id).Updates(data).Error
}
