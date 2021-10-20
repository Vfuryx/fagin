package service

import (
	"fagin/app/enums"
	"fagin/app/models/user"
	"fagin/pkg/db"
)

type userService struct{}

var User userService

func (*userService) UserList(params map[string]interface{}, columns []string, with map[string]interface{}, p *db.Paginator) ([]user.User, error) {
	var users []user.User
	err := user.NewDao().Query(params, columns, with).Paginate(&users, p)
	return users, err
}

func (*userService) AddUser(data *user.User) error {
	return user.NewDao().Create(data)
}

// ShowUser 获取可以显示的用户
func (*userService) ShowUser(id uint, columns []string) (*user.User, error) {
	params := map[string]interface{}{
		"id":     id,
		"status": enums.StatusActive.Get(),
	}
	u := user.New()
	err := u.Dao().Query(params, columns, nil).First(u)
	return u, err
}

func (*userService) UpdateUser(id uint, data map[string]interface{}) error {
	return user.NewDao().Update(id, data)
}

func (*userService) DestroyUser(id uint) error {
	return user.NewDao().Destroy(id)
}
