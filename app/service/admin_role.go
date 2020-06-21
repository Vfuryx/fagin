package service

import (
	"fagin/app/constants/admin_menu_type"
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

func (adminRoleService) Create(r *admin_role.AdminRole) error {
	err := admin_role.Dao().Create(r)
	if err != nil {
		return err
	}

	err = addPolicies(r)
	if err != nil {
		return err
	}

	return nil
}

func (adminRoleService) Update(id uint, data gin.H) error {
	err := admin_role.Dao().Update(id, data)
	if id != 1 {
		if err != nil {
			return err
		}
		params := gin.H{"id": id}
		with := gin.H{"Menus": nil}
		columns := []string{"*"}
		role := admin_role.New()
		err = role.Dao().Query(params, columns, with).First(role)
		if err != nil {
			return err
		}
		_, err = Canbin.DeletePoliciesByRole(role.Key)
		if err != nil {
			return err
		}

		err = addPolicies(role)
		if err != nil {
			return err
		}
	}

	return nil
}

// 添加权限
func addPolicies(role *admin_role.AdminRole) error {
	rules := make([][]string, 0, 20)
	for _, m := range role.Menus {
		if admin_menu_type.TypeApi == m.Type {
			rules = append(rules, []string{role.Key, m.Path, m.Action})
		}
	}
	_, err := Canbin.AddPolicies(rules)
	if err != nil {
		return err
	}
	return nil
}

func (adminRoleService) Delete(id uint) error {
	role := admin_role.New()
	err := role.Dao().FindById(id, []string{"*"})
	if err != nil {
		return err
	}
	// 删除角色以及权限
	_, err = Canbin.DeleteRole(role.Key)
	if err != nil {
		return err
	}
	return role.Dao().Delete(id)
}

func (adminRoleService) Deletes(ids []uint) error {
	var roles []admin_role.AdminRole

	err := admin_role.Dao().Query(gin.H{"in_id": ids}, []string{"*"}, nil).Find(&roles)
	if err != nil {
		return err
	}
	// 批量删除角色
	for _, r := range roles {
		_, err = Canbin.DeleteRole(r.Key)
		if err != nil {
			return err
		}
	}

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
