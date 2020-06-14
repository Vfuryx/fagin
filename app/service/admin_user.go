package service

import (
	"fagin/app"
	"fagin/app/constants/status"
	"fagin/app/errno"
	"fagin/app/models/admin_user"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type adminUserService struct{}

var AdminUserService adminUserService

func (adminUserService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]admin_user.AdminUser, error) {
	var users []admin_user.AdminUser

	err := admin_user.Dao().Query(params, columns, with).Paginator(&users, p)
	if err != nil {
		return nil, err
	}

	return users, err
}

func (adminUserService) Show(id uint, columns []string) (*admin_user.AdminUser, error) {
	m := admin_user.New()
	err := m.Dao().FindById(id, columns)
	return m, err
}

func (adminUserService) Create(m *admin_user.AdminUser) error {
	return admin_user.Dao().Create(m)
}

func (adminUserService) Update(id uint, data gin.H) error {
	return admin_user.Dao().Update(id, data)
}

func (adminUserService) Delete(id uint) error {
	return admin_user.Dao().Destroy(id)
}

func (adminUserService) Deletes(ids []uint) error {
	return admin_user.Dao().Deletes(ids)
}

func (adminUserService) UpdateStatus(id uint, status int) error {
	return admin_user.Dao().Update(id, gin.H{
		"status": status,
	})
}


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
		return errno.Api.ErrPasswordIncorrect
	}
	return nil
}

