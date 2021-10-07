package service

import (
	"fagin/app/caches"
	"fagin/app/enums"
	"fagin/app/errno"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user_role"
	"fagin/pkg/cache"
	"fagin/pkg/casbins"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type adminRoleService struct{}

var AdminRoleService adminRoleService

func (*adminRoleService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]admin_role.AdminRole, error) {
	var roles []admin_role.AdminRole

	err := admin_role.NewDao().Query(params, columns, with).Paginate(&roles, p)
	if err != nil {
		return nil, err
	}

	return roles, err
}

func (*adminRoleService) Show(id uint, columns []string) (*admin_role.AdminRole, error) {
	return admin_role.NewDao().Show(id, columns)
}

func (*adminRoleService) Create(r *admin_role.AdminRole) (err error) {
	r.Menus, err = admin_role.GetMenuAll(r.Menus)
	if err != nil {
		return err
	}
	err = admin_role.NewDao().Create(r)
	if err != nil {
		return err
	}

	err = addPolicies(r)
	if err != nil {
		return err
	}

	return nil
}

func (*adminRoleService) Update(id uint, data gin.H) error {
	err := admin_role.NewDao().Update(id, data)
	if err != nil {
		return err
	}
	// 跳过超级管理员
	if id > 1 {
		params := gin.H{"id": id}
		with := gin.H{"Menus": nil}
		columns := []string{"*"}
		role := admin_role.New()
		err = role.Dao().Query(params, columns, with).First(role)
		if err != nil {
			return err
		}
		_, err = casbins.Casbin.DeletePoliciesByRole(role.Key)
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
	rules := make([][]string, 0, 30)
	for _, m := range role.Menus {
		// 过滤相同权限
		if m.Type == enums.AdminMenuTypePermission.Get() {
			rules = append(rules, []string{role.Key, m.Path, m.Method})
		}
	}
	_, err := casbins.Casbin.AddPolicies(rules)
	if err != nil {
		return err
	}
	return nil
}

func (*adminRoleService) Delete(id uint) error {
	// 检查角色是否还存在关联
	ok, err := admin_user_role.NewDao().RoleRelationExist(id)
	if err != nil {
		return err
	}
	if ok {
		return errno.SerRoleRelationExistErr
	}

	adminRole := admin_role.New()
	err = adminRole.Dao().FindById(id, []string{"id", "key"})
	if err != nil {
		return err
	}
	// 删除角色以及权限
	_, err = casbins.Casbin.DeleteRole(adminRole.Key)
	if err != nil {
		return err
	}
	// 删除关联菜单
	_ = db.ORM().Model(&adminRole).Association("Menus").Clear()
	return adminRole.Dao().Delete(id)
}

func (*adminRoleService) Deletes(ids []uint) error {
	var roles []admin_role.AdminRole

	err := admin_role.NewDao().
		Query(gin.H{"in_id": ids}, []string{"*"}, nil).
		Find(&roles)
	if err != nil {
		return err
	}
	// 批量删除角色
	for _, r := range roles {
		_, err = casbins.Casbin.DeleteRole(r.Key)
		if err != nil {
			return err
		}
	}

	return admin_role.NewDao().Deletes(ids)
}

func (*adminRoleService) UpdateStatus(id uint, status int) error {
	return admin_role.NewDao().Update(id, gin.H{
		"status": status,
	})
}

func (*adminRoleService) List(params gin.H, columns []string, with gin.H) ([]admin_role.AdminRole, error) {
	var roles []admin_role.AdminRole

	err := admin_role.NewDao().Query(params, columns, with).Find(&roles)
	if err != nil {
		return nil, err
	}

	return roles, err
}

// RemoveRoleMenusCache 清除角色关联菜单缓存
func (*adminRoleService) RemoveRoleMenusCache(roleID uint) error {
	var admins []admin_user_role.AdminUserRole
	// 获取角色关联的用户
	params := gin.H{
		"role_id": roleID,
	}
	err := admin_user_role.NewDao().
		Query(params, []string{"admin_id"}, nil).
		Find(&admins)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return nil
		}
		return err
	}

	ca := caches.NewAdminRoutesCache(nil)
	// 清除用户关联菜单的缓存
	for _, admin := range admins {
		_, err = ca.Remove(strconv.FormatUint(uint64(admin.AdminID), 10))
		if err != nil && err != cache.ErrNotOpen {
			return err
		}
	}
	return nil
}

// KeyExist key是否存在
func (*adminRoleService) KeyExist(t uint, ket string) bool {
	params := gin.H{
		"type": t,
		"key":  ket,
	}
	return admin_role.NewDao().
		Query(params, nil, nil).
		Exists()
}
