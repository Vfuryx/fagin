package service

import (
	"fagin/app/errno"
	"fagin/app/models/user"
	"golang.org/x/crypto/bcrypt"
)

type auth struct{}

var Auth = auth{}

/**
 * 登录
 */
func (this auth) Login(email string, password string) (user.User, error) {
	// 查找用户
	user, err := user.GetUserByEmail(email)
	// 不存在用户
	if err != nil {
		return user, errno.ErrUserNotFound
	}

	//匹配密码
	if err := Compare(user.Password, password); err != nil {
		return user, errno.ErrPasswordIncorrect
	}

	return user, nil
}

// Encrypt encrypts the plain text with bcrypt.
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}
