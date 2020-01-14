package service

import (
	"fagin/app/constants/status"
	"fagin/app/models/user"
	"fagin/pkg/db"
)

type userService struct{}

var User userService

func (userService) UserList(params map[string]interface{}, columns []string, with map[string]interface{}, p *db.Paginator) ([]user.User, error) {
	var users []user.User
	err := user.Dao().Query(params, columns, with).Paginator(&users, p)
	return users, err
}

func (userService) AddUser(data *user.User) error {
	return user.Dao().Create(data)
}

// 获取可以显示的用户
func (userService) ShowUser(id uint, columns []string) (*user.User, error) {
	params := map[string]interface{}{
		"id":     id,
		"status": status.Active,
	}
	u := user.New()
	err := u.Dao().Query(params, columns, nil).First(u)
	return u, err
}

func (userService) UpdateUser(id uint, data map[string]interface{}) error {
	return user.Dao().Update(id, data)
}

func (userService) DestroyUser(id uint) error {
	return user.Dao().Destroy(id)
}
