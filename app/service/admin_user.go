package service

import (
	"fagin/app"
	"fagin/app/constants/admin_menu_type"
	"fagin/app/constants/status"
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user"
	"fagin/pkg/casbins"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
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

func (adminUserService) Show(id uint, columns []string, with gin.H) (*admin_user.AdminUser, error) {
	m := admin_user.New()
	err := m.Dao().Query(gin.H{"id": id}, columns, with).First(&m)
	return m, err
}

// Create 创建管理员
func (adminUserService) Create(u *admin_user.AdminUser, roles []admin_role.AdminRole) error {
	err := admin_user.Dao().Create(u)
	if err != nil {
		return err
	}

	for _, v := range roles {
		_, err = casbins.Casbin.AddUserRole(strconv.FormatUint(uint64(u.ID), 10), v.Key)
		if err != nil {
			return err
		}
	}

	return nil
}

// Update 编辑管理员
func (adminUserService) Update(id uint, data gin.H) (err error) {
	if v, ok := data["role_ids"]; ok {
		params := gin.H{"in_id": v, "type": 0}
		// 获取权限组
		var roles []admin_role.AdminRole
		err = admin_role.Dao().
			Query(params, []string{"*"}, nil).
			Find(&roles)
		// 角色不存在
		if err != nil && gorm.ErrRecordNotFound != err {
			return err
		}
		uid := strconv.FormatUint(uint64(id), 10)
		// 清除角色
		ok, err = casbins.Casbin.DeleteRolesForUser(uid)
		if err != nil {
			return err
		}
		// 更换角色
		for _, r := range roles {
			// 添加用户角色
			_, err = casbins.Casbin.AddUserRole(uid, r.Key)
			if err != nil {
				return err
			}
		}
		// 关联角色
		err = db.ORM().Model(&admin_user.AdminUser{ID: id}).
			Association("Roles").Replace(roles)
		if err != nil {
			return err
		}
		delete(data, "role_ids")
	}
	return admin_user.Dao().Update(id, data)
}

func (adminUserService) Delete(id uint) error {
	uid := strconv.FormatUint(uint64(id), 10)
	// 删除角色
	_, err := casbins.Casbin.DeleteRolesForUser(uid)
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
	_, err := casbins.Casbin.DeleteRolesForUsers(uIDs)
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

func (adminUserService) UserInfoById(id uint, columns []string) (*admin_user.AdminUser, error) {
	params := map[string]interface{}{
		"id":     id,
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
		return errno.Serve.ErrPasswordIncorrect
	}
	return nil
}

func (adminUserService) GetNavs(userID uint) ([]admin_menu.AdminMenu, error) {
	var err error
	params := gin.H{"id": userID}
	with := gin.H{"Roles": func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", 1).Where("type = ?", 0)
	}}
	// 获取用户角色id
	user := admin_user.New()
	err = user.Dao().Query(params, []string{"id"}, with).Find(&user)
	if err != nil {
		return nil, err
	}
	var roleIds []uint
	for _, r := range user.Roles {
		roleIds = append(roleIds, r.ID)
	}

	// 根据角色id获取菜单
	params = gin.H{"in_id": roleIds}
	with = gin.H{
		"Menus": func(db *gorm.DB) *gorm.DB {
			return db.
				Where("type = ?", admin_menu_type.TypeAdmin).
				Where("status = ?", status.Active).
				Order("sort desc, id asc")
		},
	}
	columns := []string{"*"}
	var roles []admin_role.AdminRole
	err = admin_role.Dao().Query(params, columns, with).Find(&roles)
	if err != nil {
		return nil, err
	}
	// 遍历出菜单
	var ids = make(map[uint]uint, 20)
	var menus []admin_menu.AdminMenu
	for _, adminRole := range roles {
		for index := range adminRole.Menus {
			if _, ok := ids[adminRole.Menus[index].ID]; !ok {
				ids[adminRole.Menus[index].ID] = adminRole.Menus[index].ID
				menus = append(menus, adminRole.Menus[index])
			}
		}
	}
	return menus, nil
}

// UsernameExist 用户名是否已存在
func (adminUserService) UsernameExist(username string, uid uint) bool {
	params := gin.H{
		"username":  username,
	}
	if uid > 0 {
		params["ne_id"] = uid
	}
	err := admin_user.Dao().Query(params, []string{"1"}, nil).First(&admin_user.AdminUser{})
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			return true
		}
		return false
	}
	return false
}
