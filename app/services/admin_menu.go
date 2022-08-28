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

func (*adminMenuService) All(params gin.H, columns []string) ([]admin_menu.AdminMenu, error) {
	menus, err := admin_menu.NewDao().All(params, columns)
	if err != nil {
		return nil, errorw.UP(err)
	}

	return menus, nil
}

func (*adminMenuService) Index(params gin.H, columns []string, with gin.H) ([]admin_menu.AdminMenu, error) {
	var menus []admin_menu.AdminMenu

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
	return errorw.UP(admin_menu.NewDao().CreateAndSetPaths(m))
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

//CREATE TABLE "public"."accident_reasons" (
//  "id" int8 NOT NULL DEFAULT nextval('accident_reasons_id_seq'::regclass),
//  "name" varchar(255) COLLATE "pg_catalog"."default" NOT NULL,
//  "accident_reason_category_id" int8 NOT NULL,
//  "active" bool NOT NULL,
//  "created_at" timestamp(0),
//  "updated_at" timestamp(0),
//  CONSTRAINT "accident_reasons_pkey" PRIMARY KEY ("id"),
//  CONSTRAINT "accident_reasons_accident_reason_category_id_foreign" FOREIGN KEY ("accident_reason_category_id") REFERENCES "public"."accident_reason_categories" ("id") ON DELETE CASCADE ON UPDATE NO ACTION
//)
// CREATE TABLE "public"."admin_menus" (
//  "id" int8 NOT NULL DEFAULT nextval('admin_menus_id_seq'::regclass),
//  "created_at" timestamptz(6),
//  "updated_at" timestamptz(6),
//  "deleted_at" timestamptz(6),
//  "parent_id" int8 NOT NULL DEFAULT 0,
//  "paths" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "type" int8 NOT NULL DEFAULT 0,
//  "name" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "permission" varchar(64) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "title" varchar(32) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "component" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "path" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "method" varchar(16) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "redirect" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "frame_src" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "current_active" varchar(255) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "icon" varchar(128) COLLATE "pg_catalog"."default" NOT NULL DEFAULT ''::character varying,
//  "is_show" int2 NOT NULL DEFAULT 1,
//  "is_hide_child" int2 NOT NULL DEFAULT 0,
//  "is_no_cache" int2 NOT NULL DEFAULT 0,
//  "sort" int8 NOT NULL DEFAULT 100,
//  "status" int2 NOT NULL DEFAULT 1,
//  CONSTRAINT "admin_menus_pkey" PRIMARY KEY ("id")
//)
//;
