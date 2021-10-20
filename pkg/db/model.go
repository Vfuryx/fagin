package db

import (
	"errors"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type DAO interface {
	FindById(id uint, columns []string) error
	Create(data interface{}) error
	Update(id uint, data map[string]interface{}) error
	Destroy(id uint) error
	Query(params map[string]interface{}, columns []string, with map[string]interface{}) DAO
	Find(model interface{}) error
	First(model interface{}) error
	Count() (int64, error)
	Exists() bool
	Paginate(model interface{}, p *Paginator) error
}

type Dao struct {
	m  interface{}
	DB *gorm.DB
}

func (d *Dao) SetModel(model interface{}) {
	d.m = model
}
func (d *Dao) GetModel() interface{} {
	return d.m
}

func (d *Dao) Find(model interface{}) error {
	if d.DB == nil {
		return ORM().Find(model).Error
	}
	return d.DB.Find(model).Error
}

func (d *Dao) First(model interface{}) error {
	if d.DB == nil {
		return ORM().First(model).Error
	}
	return d.DB.First(model).Error
}

func (d *Dao) FindById(id uint, columns []string) error {
	return ORM().Select(columns).Where("id = ?", id).First(d.GetModel()).Error
}

func (d *Dao) Create(data interface{}) error {
	return ORM().Create(data).Error
}

func (d *Dao) Update(id uint, data map[string]interface{}) error {
	err := d.FindById(id, []string{"id"})
	if err != nil {
		return err
	}
	return ORM().Model(d.GetModel()).Where("id = ?", id).Updates(data).Error
}

func (d *Dao) Destroy(id uint) error {
	return ORM().Where("id = ?", id).Delete(d.GetModel()).Error
}

func (d *Dao) With(db *gorm.DB, with map[string]interface{}) *gorm.DB {
	for index, value := range with {
		if value != nil {
			switch value.(type) {
			case gin.H:
				for k, v := range value.(gin.H) {
					db = db.Preload(index, k, v)
				}
			case func(db *gorm.DB) *gorm.DB:
				db = db.Preload(index, value)
			}

		} else {
			db = db.Preload(index)
		}
	}
	return db
}

func (d *Dao) Count() (count int64, err error) {
	err = d.DB.Model(d.GetModel()).Select([]string{}).Count(&count).Error
	return count, err
}

func (d *Dao) Exists() bool {
	err := d.DB.Model(d.GetModel()).
		Select([]string{"1"}).
		First(&struct {
			A bool `gorm:"column:1;"`
		}{}).
		Error

	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return false
		}
		return true
	}
	return true
}

// Paginate 分页处理
func (d *Dao) Paginate(model interface{}, p *Paginator) error {
	if d.DB == nil {
		d.DB = ORM()
	}
	return p.Paginate(d.DB, model)
}
