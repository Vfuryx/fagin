package user

import (
	"errors"
	"fagin/app"
	"fagin/pkg/db"
)

func New() *User {
	return &User{}
}

type Dao struct {
	db.Dao
}

var _ db.DAO = &Dao{}

func (u *User) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(u)

	return dao
}

func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())

	return dao
}

func (d *Dao) All(columns []string) ([]*User, error) {
	var model []*User
	err := db.ORM().Select(columns).Find(&model).Error

	return model, err
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
func (d *Dao) Create(data interface{}) error {
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

// Update 更新用户
func (d *Dao) Update(id uint, data map[string]interface{}) error {
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

func (d *Dao) GetUserByEmail(email string) error {
	params := map[string]interface{}{
		"email": email,
	}

	return d.Query(params, []string{"*"}, nil).First(d.GetModel())
}
