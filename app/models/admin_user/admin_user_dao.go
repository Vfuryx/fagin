package admin_user

import (
	"fagin/pkg/db"
	"fmt"
	"gorm.io/gorm"
	"time"
)

func New() *AdminUser {
	return &AdminUser{}
}

type Dao struct {
	db.Dao
}

var _ db.IDao = &Dao{}

func (m *AdminUser) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

func (d *Dao) All(columns []string) (*[]AdminUser, error) {
	var model []AdminUser
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

	if v, ok = params["ne_id"]; ok {
		model = model.Where("id != ?", v)
	}

	if v, ok = params["username"]; ok {
		model = model.Where("username = ?", v)
	}

	if v, ok = params["status"]; ok {
		model = model.Where("status = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *Dao) Deletes(ids []uint) error {
	var admin AdminUser
	return db.ORM().Where("id in (?)", ids).Delete(&admin).Error
}

func (d *Dao) Login(id uint) {
	t := time.Now().Unix()
	err := d.SetLoginAt(id, t)
	fmt.Println(err)
	err = d.SetCheckInAt(id, t)
	fmt.Println(err)
}

// SetLoginAt 设置登录时间
func (d *Dao) SetLoginAt(id uint, time int64) error {
	return db.ORM().Model(d.GetModel()).Where("id = ?", id).Updates(map[string]interface{}{
		"login_at":      time,
		"last_login_at": gorm.Expr("`login_at`"),
	}).Error
}

// SetCheckInAt 设置签发时间
func (d *Dao) SetCheckInAt(id uint, time int64) error {
	return db.ORM().Model(d.GetModel()).Where("id = ?", id).Updates(map[string]interface{}{
		"check_in_at": time,
	}).Error
}
