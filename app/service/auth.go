package service

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/user"
)

type authService struct{}

var Auth authService

/**
 * 登录
 */
func (authService) Login(email string, password string) (user.User, error) {
	// 查找用户
	u := user.New()
	err := u.Dao().GetUserByEmail(email)
	// 不存在用户
	if err != nil {
		return *u, errno.ErrUserNotFound
	}

	//匹配密码
	if err := app.Compare(u.Password, password); err != nil {
		return *u, errno.ErrPasswordIncorrect
	}

	return *u, nil
}
