package service

import (
	"fagin/app/models/admin_role"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type adminRoleService struct{}

var AdminRoleService adminRoleService

func (adminRoleService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]admin_role.AdminRole, error) {
	var roles []admin_role.AdminRole

	err := admin_role.Dao().Query(params, columns, with).Paginator(&roles, p)
	if err != nil {
		return nil, err
	}

	return roles, err
}

func (adminRoleService) Show(id uint, columns []string) (*admin_role.AdminRole, error) {
	return admin_role.Dao().Show(id, columns)
}

func (adminRoleService) Create(m *admin_role.AdminRole) error {
	return admin_role.Dao().Create(m)
}

func (adminRoleService) Update(id uint, data gin.H) error {
	return admin_role.Dao().Update(id, data)
}

func (adminRoleService) Delete(id uint) error {
	return admin_role.Dao().Delete(id)
}

func (adminRoleService) Deletes(ids []uint) error {
	return admin_role.Dao().Deletes(ids)
}

func (adminRoleService) UpdateStatus(id uint, status int) error {
	return admin_role.Dao().Update(id, gin.H{
		"status": status,
	})
}

func (adminRoleService) List(params gin.H, columns []string, with gin.H) ([]admin_role.AdminRole, error) {
	var roles []admin_role.AdminRole

	err := admin_role.Dao().Query(params, columns, with).Find(&roles)
	if err != nil {
		return nil, err
	}

	return roles, err
}
