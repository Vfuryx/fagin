package service

import (
	"fagin/app"
	"fagin/app/constants/status"
	"fagin/app/errno"
	"fagin/app/models/admin_user"
)

type adminUserService struct{}

var AdminUser adminUserService

func (adminUserService) UserInfo(name string, columns []string) (*admin_user.AdminUser, error) {
	params := map[string]interface{}{
		"name":   name,
		"status": status.Active,
	}
	au := admin_user.New()
	err := au.Dao().Query(params, columns, nil).First(au)
	return au, err
}

func (adminUserService) UpdateAdminUser(id uint, data map[string]interface{}) error {
	return admin_user.Dao().Update(id, data)
}

func (adminUserService) CheckPassword(id uint, old string) error {
	au := admin_user.New()
	err := au.Dao().FindById(id, []string{"id", "password"})
	if err != nil {
		return err
	}
	if err = app.Compare(au.Password, old); err != nil {
		return errno.ErrPasswordIncorrect
	}
	return nil
}
