package service

import (
	"fagin/app"
	"fagin/app/constants/admin_menu_type"
	"fagin/app/constants/status"
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
	"strconv"
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

func (adminUserService) Create(u *admin_user.AdminUser, role *admin_role.AdminRole) error {
	err := admin_user.Dao().Create(u)
	if err != nil {
		return err
	}

	_, err = Canbin.AddUserRole(strconv.FormatUint(uint64(u.ID), 10), role.Key)
	if err != nil {
		return err
	}

	return nil
}

func (adminUserService) Update(id uint, data gin.H, role *admin_role.AdminRole) error {
	var user = admin_user.New()
	if role != nil {
		err := user.Dao().FindById(id, []string{"*"})
		if err != nil {
			return err
		}
		// 是否更换角色
		if role.ID != user.RoleID {
			uid := strconv.FormatUint(uint64(user.ID), 10)
			// 删除角色
			_, err = Canbin.DeleteRolesForUser(uid)
			if err != nil {
				return err
			}
			// 添加用户角色
			_, err = Canbin.AddUserRole(uid, role.Key)
			if err != nil {
				return err
			}
		}
	}

	err := user.Dao().Update(id, data)
	if err != nil {
		return err
	}

	return nil
}

func (adminUserService) Delete(id uint) error {
	uid := strconv.FormatUint(uint64(id), 10)
	// 删除角色
	_, err := Canbin.DeleteRolesForUser(uid)
	if err != nil {
		return err
	}
	return admin_user.Dao().Destroy(id)
}

func (adminUserService) Deletes(ids []uint) error {
	uIDs := make([]string, 0, 20)
	for _, id := range ids {
		uid := strconv.FormatUint(uint64(id), 10)
		uIDs = append(uIDs, uid)
	}
	// 删除角色
	_, err := Canbin.DeleteRolesForUsers(uIDs)
	if err != nil {
		return err
	}
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

func (adminUserService) GetRoutes(userID uint) ([]admin_menu.AdminMenu, error) {
	// 获取用户角色id
	user := admin_user.New()
	err := user.Dao().FindById(userID, []string{"id", "role_id"})
	if err != nil {
		return nil, err
	}

	var menus []admin_menu.AdminMenu
	// 判断是否超级管理员
	if user.RoleID == 1 {
		// 获取展示
		params := gin.H{
			"in_type": []int{int(admin_menu_type.TypeMenu), int(admin_menu_type.TypeMain)},
			"visible": status.Active,
			"sort":    "sort desc, id asc",
		}
		columns := []string{"*"}
		err = admin_menu.Dao().Query(params, columns, nil).Find(&menus)
		if err != nil {
			return nil, err
		}
	} else {
		// 根据角色id获取菜单
		params := gin.H{"id": user.RoleID}
		with := gin.H{
			"Menus": func(db *gorm.DB) *gorm.DB {
				return db.
					Where("type in (?)", []uint{uint(admin_menu_type.TypeMenu), uint(admin_menu_type.TypeMain)}).
					Where("visible = ?", status.Active).
					Order("sort desc, id asc")
			},
		}
		columns := []string{"*"}
		role := admin_role.New()
		err = role.Dao().Query(params, columns, with).First(role)
		if err != nil {
			return nil, err
		}

		menus = role.Menus
	}

	//组合菜单
	menus = getMenuTree(menus, 0)

	return menus, nil
}
