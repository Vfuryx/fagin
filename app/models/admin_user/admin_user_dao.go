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

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (m *AdminUser) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = m
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(columns []string) (*[]AdminUser, error) {
	var model []AdminUser
	err := db.ORM().Select(columns).Find(&model).Error
	return &model, err
}

func (d *dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.IDao {
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

func (d *dao) Deletes(ids []uint) error {
	var admin AdminUser
	return db.ORM().Where("id in (?)", ids).Delete(&admin).Error
}

func (d *dao) Login(id uint) {
	t := time.Now().Unix()
	err := d.SetLoginAt(id, t)
	fmt.Println(err)
	err = d.SetCheckInAt(id, t)
	fmt.Println(err)
}

// SetLoginAt 设置登录时间
func (d *dao) SetLoginAt(id uint, time int64) error {
	return db.ORM().Model(d.Dao.M).Where("id = ?", id).Updates(map[string]interface{}{
		"login_at":      time,
		"last_login_at": gorm.Expr("`login_at`"),
	}).Error
}

// SetCheckInAt 设置签发时间
func (d *dao) SetCheckInAt(id uint, time int64) error {
	return db.ORM().Model(d.Dao.M).Where("id = ?", id).Updates(map[string]interface{}{
		"check_in_at": time,
	}).Error
}
