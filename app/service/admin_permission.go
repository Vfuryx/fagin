package service

import (
	"fagin/app/errno"
	"fagin/app/models/admin_permission"
	"fagin/app/models/admin_permission_group"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type adminPermissionService struct{}

var AdminPermissionService adminPermissionService

func (adminPermissionService) PermissionGroupList(
	params gin.H, columns []string, with gin.H, p *db.Paginator) ([]admin_permission_group.AdminPermissionGroup, error) {
	var groups []admin_permission_group.AdminPermissionGroup

	err := admin_permission_group.Dao().Query(params, columns, with).Paginator(&groups, p)
	if err != nil {
		return nil, err
	}

	return groups, err
}

func (adminPermissionService) GroupPermissions(params gin.H, columns []string, with gin.H) ([]admin_permission_group.AdminPermissionGroup, error) {
	var groups []admin_permission_group.AdminPermissionGroup
	err := admin_permission_group.Dao().Query(params, columns, with).Find(&groups)
	if err != nil {
		return nil, err
	}
	return groups, err
}

func (adminPermissionService) PermissionGroupAll(
	params gin.H, columns []string, ) (*[]admin_permission_group.AdminPermissionGroup, error) {
	return admin_permission_group.Dao().All(params, columns)
}

func (adminPermissionService) PermissionGroupShow(id uint, columns []string) (*admin_permission_group.AdminPermissionGroup, error) {
	m := admin_permission_group.New()
	err := m.Dao().FindById(id, columns)
	return m, err
}

func (adminPermissionService) PermissionGroupCreate(m *admin_permission_group.AdminPermissionGroup) error {
	return admin_permission_group.Dao().Create(m)
}

func (adminPermissionService) PermissionGroupUpdate(id uint, data gin.H) error {
	return admin_permission_group.Dao().Update(id, data)
}

func (adminPermissionService) PermissionList(
	params gin.H, columns []string, with gin.H, p *db.Paginator) ([]admin_permission.AdminPermission, error) {
	var groups []admin_permission.AdminPermission

	err := admin_permission.Dao().Query(params, columns, with).Paginator(&groups, p)
	if err != nil {
		return nil, err
	}

	return groups, err
}

func (adminPermissionService) PermissionShow(id uint, columns []string) (*admin_permission.AdminPermission, error) {
	m := admin_permission.New()
	err := m.Dao().FindById(id, columns)
	return m, err
}

func (adminPermissionService) PermissionCreate(m *admin_permission.AdminPermission) error {
	return admin_permission.Dao().Create(m)
}

func (adminPermissionService) PermissionUpdate(id uint, data gin.H) error {
	return admin_permission.Dao().Update(id, data)
}

func (adminPermissionService) FindByPath(
	method, path string, column []string) (admin_permission.AdminPermission, error) {
	params := gin.H{"path": path, "method": method}
	var m admin_permission.AdminPermission
	err := admin_permission.Dao().Query(params, column, nil).First(&m)
	return m, err
}

func (adminPermissionService) PermissionDelete(id uint) error {
	// 检查权限是否还存在关联
	ok, err := admin_permission.Dao().PermissionRelationExist(id)
	if err != nil {
		return err
	}
	if ok {
		return errno.Serve.PermissionRelationExistErr
	}
	return admin_permission.Dao().Destroy(id)
}
