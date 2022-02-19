package admin_user

import (
	"fagin/pkg/db"
	"time"

	"gorm.io/gorm"
)

// New 实例化
func New() *AdminUser {
	return &AdminUser{}
}

// Dao Dao
type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

// Dao 获取Dao
func (m *AdminUser) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)

	return dao
}

// NewDao 实例化
func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())

	return dao
}

// All 所有数据
func (d *Dao) All(columns []string) ([]*AdminUser, error) {
	var model []*AdminUser
	err := db.ORM().Select(columns).Find(&model).Error

	return model, err
}

// Query 通用查询
func (d *Dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.DAO {
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

// Deletes 批量删除
func (d *Dao) Deletes(ids []uint) error {
	var admin AdminUser
	return db.ORM().Where("id in (?)", ids).Delete(&admin).Error
}

// Login 登录
func (d *Dao) Login(id uint) {
	t := time.Now().Unix()
	_ = d.SetLoginAt(id, t)
	_ = d.SetCheckInAt(id, t)
}

// SetLoginAt 设置登录时间
func (d *Dao) SetLoginAt(id uint, t int64) error {
	return db.ORM().Model(d.GetModel()).Where("id = ?", id).Updates(map[string]interface{}{
		"login_at":      t,
		"last_login_at": gorm.Expr("`login_at`"),
	}).Error
}

// SetCheckInAt 设置签发时间
func (d *Dao) SetCheckInAt(id uint, t int64) error {
	return db.ORM().Model(d.GetModel()).Where("id = ?", id).Updates(map[string]interface{}{
		"check_in_at": t,
	}).Error
}
