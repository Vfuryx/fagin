package services

import (
	"fagin/app"
	"fagin/app/enums"
	"fagin/app/errno"
	"fagin/app/models/admin_menu"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user"
	"fagin/pkg/casbins"
	"fagin/pkg/db"
	"fagin/pkg/errorw"
	"fagin/utils"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type adminUserService struct{}

// AdminUserService 后台用户服务
var AdminUserService adminUserService

func (*adminUserService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]*admin_user.AdminUser, error) {
	var users []*admin_user.AdminUser
	err := admin_user.NewDao().Query(params, columns, with).Paginate(&users, p)

	return users, errorw.UP(err)
}

func (*adminUserService) Show(id uint, columns []string, with gin.H) (*admin_user.AdminUser, error) {
	m := admin_user.New()
	err := m.Dao().Query(gin.H{"id": id}, columns, with).First(&m)

	return m, errorw.UP(err)
}

// Create 创建管理员
func (*adminUserService) Create(u *admin_user.AdminUser, roles []*admin_role.AdminRole) error {
	err := admin_user.NewDao().Create(u)
	if err != nil {
		return errorw.UP(err)
	}

	for _, v := range roles {
		_, err = casbins.Casbin.AddUserRole(utils.Uint64ToStr(uint64(u.ID)), v.Key)
		if err != nil {
			return errorw.UP(err)
		}
	}

	return nil
}

// Update 编辑管理员
func (*adminUserService) Update(id uint, data gin.H) (err error) {
	if v, ok := data["role_ids"]; ok {
		params := gin.H{"in_id": v, "type": 0}
		// 获取权限组
		var roles []admin_role.AdminRole
		err = admin_role.NewDao().
			Query(params, []string{"*"}, nil).
			Find(&roles)
		// 角色不存在
		if err != nil && gorm.ErrRecordNotFound != err {
			return errorw.UP(err)
		}

		uid := utils.Uint64ToStr(uint64(id))
		// 清除角色
		_, err = casbins.Casbin.DeleteRolesForUser(uid)
		if err != nil {
			return errorw.UP(err)
		}
		// 更换角色
		for i := range roles {
			// 添加用户角色
			_, err = casbins.Casbin.AddUserRole(uid, roles[i].Key)
			if err != nil {
				return errorw.UP(err)
			}
		}
		// 关联角色
		err = db.ORM().Model(&admin_user.AdminUser{ID: id}).
			Association("Roles").Replace(roles)
		if err != nil {
			return errorw.UP(err)
		}

		delete(data, "role_ids")
	}

	err = admin_user.NewDao().Update(id, data)

	return errorw.UP(err)
}

func (*adminUserService) Delete(id uint) error {
	uid := utils.Uint64ToStr(uint64(id))
	// 删除角色
	_, err := casbins.Casbin.DeleteRolesForUser(uid)
	if err != nil {
		return errorw.UP(err)
	}

	err = admin_user.NewDao().Destroy(id)

	return errorw.UP(err)
}

func (*adminUserService) Deletes(ids []uint) error {
	uIDs := make([]string, 0)

	for _, id := range ids {
		uIDs = append(uIDs, utils.Uint64ToStr(uint64(id)))
	}
	// 删除角色
	_, err := casbins.Casbin.DeleteRolesForUsers(uIDs)
	if err != nil {
		return errorw.UP(err)
	}

	err = admin_user.NewDao().Deletes(ids)

	return errorw.UP(err)
}

func (*adminUserService) UpdateStatus(id uint, status int) error {
	err := admin_user.NewDao().Update(id, gin.H{
		"status": status,
	})

	return errorw.UP(err)
}

// UserInfoByID 根据ID获取用户信息
func (*adminUserService) UserInfoByID(id uint, columns []string) (*admin_user.AdminUser, error) {
	params := map[string]interface{}{
		"id":     id,
		"status": enums.StatusActive.Get(),
	}
	au := admin_user.New()
	err := au.Dao().Query(params, columns, nil).First(au)

	return au, errorw.UP(err)
}

func (*adminUserService) UpdateAdminUser(id uint, data map[string]interface{}) error {
	err := admin_user.NewDao().Update(id, data)
	return errorw.UP(err)
}

func (*adminUserService) CheckPassword(id uint, old string) error {
	au := admin_user.New()
	err := au.Dao().FindByID(id, []string{"id", "password"})

	if err != nil {
		return errorw.UP(err)
	}

	if err = app.Compare(au.Password, old); err != nil {
		return errorw.UP(errno.SerAccountOrPasswordErr)
	}

	return nil
}

func (ser *adminUserService) GetRoutes(userID uint) ([]admin_menu.AdminMenu, error) {
	var err error

	params := gin.H{"id": userID}
	with := gin.H{"Roles": func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", enums.StatusActive.Get()).
			Where("type = ?", 0)
	}}
	// 获取用户角色id
	user := admin_user.New()

	err = user.Dao().Query(params, []string{"id"}, with).Find(&user)
	if err != nil {
		return nil, errorw.UP(err)
	}

	var roleIds = make([]uint, 0)
	for _, r := range user.Roles {
		roleIds = append(roleIds, r.ID)
	}

	// 根据角色id获取菜单
	return ser.GetRoutesByRoleIDs(roleIds)
}

func (*adminUserService) GetRoutesByRoleIDs(roleIDs []uint) ([]admin_menu.AdminMenu, error) {
	// 根据角色id获取菜单
	params := gin.H{"in_id": roleIDs}
	with := gin.H{
		"Menus": func(db *gorm.DB) *gorm.DB {
			return db.
				Where("type in (?)", []int{enums.AdminMenuTypeDir.Get(), enums.AdminMenuTypeMenu.Get()}).
				Where("status = ?", enums.StatusActive.Get()).
				Order("sort desc, id asc")
		},
	}
	columns := []string{"*"}

	var roles []*admin_role.AdminRole

	err := admin_role.NewDao().Query(params, columns, with).Find(&roles)
	if err != nil {
		return nil, errorw.UP(err)
	}

	var (
		// 遍历出菜单
		ids   = make(map[uint]uint)
		menus []admin_menu.AdminMenu
	)

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

func (*adminUserService) GetPermissionCode(userID uint) ([]string, error) {
	var err error

	params := gin.H{"id": userID}
	with := gin.H{"Roles": func(db *gorm.DB) *gorm.DB {
		return db.Where("status = ?", enums.StatusActive.Get()).
			Where("type = ?", 0)
	}}
	// 获取用户角色id
	user := admin_user.New()

	err = user.Dao().Query(params, []string{"id"}, with).Find(&user)
	if err != nil {
		return nil, errorw.UP(err)
	}

	var roleIds = make([]uint, 0)
	for _, r := range user.Roles {
		roleIds = append(roleIds, r.ID)
	}

	// 根据角色id获取菜单
	params = gin.H{"in_id": roleIds}
	with = gin.H{
		"Menus": func(db *gorm.DB) *gorm.DB {
			return db.
				Select([]string{"id", "permission"}).
				Where("type = ?", enums.AdminMenuTypePermission.Get()).
				Where("status = ?", enums.StatusActive.Get()).
				Order("sort desc, id asc")
		},
	}

	columns := []string{"*"}

	var roles []admin_role.AdminRole

	err = admin_role.NewDao().Query(params, columns, with).Find(&roles)
	if err != nil {
		return nil, errorw.UP(err)
	}

	var (
		// 遍历出菜单
		ids   = make(map[uint]uint)
		menus = make([]string, 0)
	)

	for i := range roles {
		for index := range roles[i].Menus {
			if _, ok := ids[roles[i].Menus[index].ID]; !ok {
				ids[roles[i].Menus[index].ID] = roles[i].Menus[index].ID
				menus = append(menus, roles[i].Menus[index].Permission)
			}
		}
	}

	return menus, nil
}

// UsernameExists 用户名是否已存在
func (*adminUserService) UsernameExists(username string, uid uint) bool {
	params := gin.H{
		"username": username,
	}
	if uid > 0 {
		params["ne_id"] = uid
	}

	return admin_user.NewDao().Query(params, nil, nil).Exists()
}
