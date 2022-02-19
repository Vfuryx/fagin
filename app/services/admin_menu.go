package services

import (
	"errors"
	"fagin/app/caches"
	"fagin/app/enums"
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_role_menu"
	"fagin/app/models/admin_user_role"
	"fagin/pkg/cache"
	"fagin/pkg/casbins"
	"fagin/pkg/db"
	"fagin/pkg/errorw"
	"fagin/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type adminMenuService struct{}

// AdminMenuService 后台菜单服务
var AdminMenuService adminMenuService

func (*adminMenuService) All(params gin.H, columns []string) ([]*admin_menu.AdminMenu, error) {
	menus, err := admin_menu.NewDao().All(params, columns)
	if err != nil {
		return nil, errorw.UP(err)
	}

	return menus, nil
}

func (*adminMenuService) Index(params gin.H, columns []string, with gin.H) ([]*admin_menu.AdminMenu, error) {
	var menus []*admin_menu.AdminMenu

	err := admin_menu.NewDao().Query(params, columns, with).Find(&menus)
	if err != nil {
		return nil, errorw.UP(err)
	}

	return menus, nil
}

func (*adminMenuService) Show(id uint, columns []string) (*admin_menu.AdminMenu, error) {
	m := admin_menu.New()
	err := m.Dao().FindByID(id, columns)

	return m, errorw.UP(err)
}

func (*adminMenuService) MenuExists(id uint) bool {
	return admin_menu.NewDao().ExistsByID(id)
}

func (*adminMenuService) Create(m *admin_menu.AdminMenu) error {
	err := admin_menu.NewDao().Create(m)
	if err != nil {
		return errorw.UP(err)
	}
	// 设置路径
	return setPaths(m)
}

// 设置路径
func setPaths(menus *admin_menu.AdminMenu) error {
	adminMenu := admin_menu.New()

	paths := ""

	// 判断父ID是否为0
	if menus.ParentID == 0 {
		paths = "0-" + utils.Uint64ToStr(uint64(menus.ID))
	} else {
		err := adminMenu.Dao().FindByID(menus.ParentID, []string{"id", "paths"})
		if err != nil {
			return errorw.UP(err)
		}
		paths = adminMenu.Paths + "-" + utils.Uint64ToStr(uint64(menus.ID))
	}

	return errorw.UP(menus.Dao().Update(menus.ID, gin.H{"paths": paths}))
}

func (*adminMenuService) Update(id uint, data gin.H) error {
	menu := admin_menu.New()

	err := menu.Dao().FindByID(id, []string{"*"})
	if err != nil {
		return errorw.UP(err)
	}

	// 获取path
	path := menu.Path

	v, ok := data["path"]
	if ok {
		p2, ok := v.(string)
		if ok {
			path = p2
		}
	}
	// 获取method
	method := menu.Method

	v, ok = data["method"]
	if ok {
		ac, ok := v.(string)
		if ok {
			method = ac
		}
	}

	//  判断是否api 是否有改变
	if path != menu.Path || method != menu.Method {
		// 获取关联角色
		var rms []admin_role_menu.AdminRoleMenu

		err = admin_role_menu.NewDao().Query(gin.H{"menu_id": id}, []string{"*"}, nil).Find(&rms)
		if err != nil {
			return errorw.UP(err)
		}

		ids := make([]uint, 0)

		if len(rms) > 0 {
			for _, rm := range rms {
				ids = append(ids, rm.RoleID)
			}
			// 获取角色
			var roles []admin_role.AdminRole

			err = admin_role.NewDao().Query(gin.H{"in_id": ids}, []string{"*"}, nil).Find(&roles)
			if err != nil {
				return errorw.UP(err)
			}

			for i := range roles {
				// 删除原来的角色权限
				_, err = casbins.Casbin.RemovePolicy(roles[i].Key, menu.Path, menu.Method)
				if err != nil {
					return errorw.UP(err)
				}

				// 新增角色权限
				_, err = casbins.Casbin.AddPolicyForRole(roles[i].Key, path, method)
				if err != nil {
					return errorw.UP(err)
				}
			}
		}
	}

	return admin_menu.NewDao().Update(id, data)
}

func (*adminMenuService) Delete(id uint) error {
	adminMenu := admin_menu.New()

	err := adminMenu.Dao().FindByID(id, []string{"id", "paths"})
	if err != nil {
		return errorw.UP(err)
	}

	if adminMenu.Paths == "" || adminMenu.Paths == "0" || adminMenu.Paths == "0-" {
		return errorw.UP(errno.DaoMenuPathsUnsafeErr)
	}

	params := gin.H{
		"like_paths": adminMenu.Paths + "-%",
	}
	// 是否存在下级
	ok := adminMenu.Dao().Query(params, nil, nil).Exists()
	if ok {
		return errorw.UP(errno.SerMenuSubExistErr)
	}

	// 是否存在关联
	ok = admin_role_menu.NewDao().MenuRelationExist(id)
	if ok {
		return errorw.UP(errno.SerMenuRelationExistErr)
	}

	err = admin_menu.NewDao().Delete(id)

	return errorw.UP(err)
}

// RemoveUserMenusCache 清除用户关联菜单缓存
func (*adminMenuService) RemoveUserMenusCache(menuID uint) error {
	var admins []admin_user_role.AdminUserRole

	err := db.ORM().Table("admin_role_menus as arm").
		Select([]string{"aur.admin_id", "aur.role_id"}).
		Joins("left join admin_user_roles as aur on aur.role_id = arm.role_id").
		Where("arm.menu_id = ?", menuID).
		Find(&admins).Error
	if err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil
		}

		return errorw.UP(err)
	}

	ca := caches.NewAdminRoutesCache(nil)
	// 清除用户关联菜单的缓存
	for _, admin := range admins {
		_, err = ca.Remove(utils.Uint64ToStr(uint64(admin.AdminID)))
		if err != nil && err != cache.ErrNotOpen {
			return errorw.UP(err)
		}
	}

	return nil
}

func (*adminMenuService) FindByPath(method, path string, column []string) (admin_menu.AdminMenu, error) {
	params := gin.H{
		"type":   enums.AdminMenuTypePermission.Get(),
		"path":   path,
		"method": method,
	}

	var m admin_menu.AdminMenu

	err := admin_menu.NewDao().Query(params, column, nil).First(&m)

	return m, errorw.UP(err)
}
