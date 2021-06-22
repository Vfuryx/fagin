package user

import (
	"fagin/app"
	"fagin/pkg/db"
	"github.com/pkg/errors"
)

func New() *User {
	return &User{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (u *User) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = u
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(columns []string) (*[]User, error) {
	var model []User
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

	if v, ok = params["username"]; ok {
		model = model.Where("username = ?", v)
	}

	if v, ok = params["status"]; ok {
		model = model.Where("status = ?", v)
	}

	if v, ok = params["email"]; ok {
		model = model.Where("email = ?", v)
	}

	d.DB = d.With(model, with)
	return d
}

// Create 创建用户
func (d *dao) Create(data interface{}) error {
	u, ok := data.(*User)
	if !ok {
		return errors.New("数据出错")
	}
	if u.Password != "" {
		// 加密密码
		p, err := app.Encrypt(u.Password)
		if err != nil {
			return err
		}
		u.Password = p
	}
	return d.Dao.Create(u)
}

// 更新用户
func (d *dao) Update(id uint, data map[string]interface{}) error {
	if v, ok := data["password"]; ok && v != "" {
		// 加密密码
		p, err := app.Encrypt(v.(string))
		if err != nil {
			return err
		}
		data["password"] = p
	}
	return d.Dao.Update(id, data)
}

func (d *dao) GetUserByEmail(email string) error {
	params := map[string]interface{}{
		"email": email,
	}
	return d.Query(params, []string{"*"}, nil).First(d.M)
}
