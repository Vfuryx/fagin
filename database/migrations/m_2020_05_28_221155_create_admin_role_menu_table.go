package migrations

import (
	"fagin/app/models/admin_menu"
	"fagin/pkg/migrate/migration"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	// 角色菜单关联表
	type AdminRoleMenu struct {
		RoleID uint `gorm:"type:int(10);not null;default:0;comment:'角色id';column:role_id;"`
		MenuID uint `gorm:"type:int(10);not null;default:0;comment:'菜单id';column:menu_id;"`
	}
	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
		ID: "m_2020_05_28_221155_create_admin_role_menu_table",
		Migrate: func(tx *gorm.DB) error {
			err := tx.AutoMigrate(&AdminRoleMenu{}).Error
			if err != nil {
				return err
			}

			menus := []admin_menu.AdminMenu{
				{
					ID:         1,
					ParentId:   0,
					Paths:      "0-1",
					Name:       "Admin",
					Title:      "系统管理",
					Icon:       "example",
					Type:       0,
					Path:       "/admin",
					Component:  "Layout",
					Action:     "",
					Permission: "",
					Sort:       1,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         2,
					ParentId:   1,
					Paths:      "0-1-2",
					Name:       "User",
					Title:      "用户管理",
					Icon:       "user",
					Type:       1,
					Path:       "user",
					Component:  "/admin/user",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         3,
					ParentId:   1,
					Paths:      "0-1-3",
					Name:       "Menu",
					Title:      "菜单管理",
					Icon:       "nested",
					Type:       1,
					Path:       "menu",
					Component:  "/menu/index",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         4,
					ParentId:   1,
					Paths:      "0-1-4",
					Name:       "Role",
					Title:      "角色管理",
					Icon:       "peoples",
					Type:       1,
					Path:       "role",
					Component:  "/role/index",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         5,
					ParentId:   1,
					Paths:      "0-1-5",
					Name:       "WebsitesSetup",
					Title:      "网站设置",
					Icon:       "component",
					Type:       1,
					Path:       "website",
					Component:  "/admin/website",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         6,
					ParentId:   0,
					Paths:      "0-6",
					Name:       "API管理",
					Title:      "API管理",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         7,
					ParentId:   6,
					Paths:      "0-6-7",
					Name:       "V1",
					Title:      "V1",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         8,
					ParentId:   7,
					Paths:      "0-6-7-8",
					Name:       "系统管理API",
					Title:      "系统管理API",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         9,
					ParentId:   7,
					Paths:      "0-6-7-9",
					Name:       "权限API",
					Title:      "权限API",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         10,
					ParentId:   9,
					Paths:      "0-6-7-9-10",
					Name:       "用户详情",
					Title:      "用户详情",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/us/info",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         11,
					ParentId:   9,
					Paths:      "0-6-7-9-11",
					Name:       "用户菜单列表",
					Title:      "用户菜单列表",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/us/menurole",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         12,
					ParentId:   9,
					Paths:      "0-6-7-9-12",
					Name:       "前端路由",
					Title:      "前端路由",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/us/routes",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         13,
					ParentId:   9,
					Paths:      "0-6-7-9-13",
					Name:       "",
					Title:      "用户退出",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/user/logout",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         14,
					ParentId:   8,
					Paths:      "0-6-7-8-14",
					Name:       "",
					Title:      "用户管理",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         15,
					ParentId:   14,
					Paths:      "0-6-7-8-14-15",
					Name:       "",
					Title:      "用户列表",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/user/",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         16,
					ParentId:   14,
					Paths:      "0-6-7-8-14-16",
					Name:       "",
					Title:      "展示用户",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/user/:id",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         17,
					ParentId:   14,
					Paths:      "0-6-7-8-14-17",
					Name:       "",
					Title:      "新增用户",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/user/",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         18,
					ParentId:   14,
					Paths:      "0-6-7-8-14-18",
					Name:       "",
					Title:      "更新用户",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/user/:id",
					Component:  "",
					Action:     "PUT",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         19,
					ParentId:   14,
					Paths:      "0-6-7-8-14-19",
					Name:       "",
					Title:      "更新用户状态",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/user/:id/status/",
					Component:  "",
					Action:     "PUT",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         20,
					ParentId:   14,
					Paths:      "0-6-7-8-14-20",
					Name:       "",
					Title:      "用户重置密码",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/user/:id/reset/",
					Component:  "",
					Action:     "PUT",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         21,
					ParentId:   14,
					Paths:      "0-6-7-8-14-21",
					Name:       "",
					Title:      "用户删除",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/user/:id",
					Component:  "",
					Action:     "DELETE",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         22,
					ParentId:   14,
					Paths:      "0-6-7-8-14-22",
					Name:       "",
					Title:      "用户批量删除",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/user/",
					Component:  "",
					Action:     "DELETE",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         23,
					ParentId:   8,
					Paths:      "0-6-7-8-23",
					Name:       "",
					Title:      "菜单管理",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         24,
					ParentId:   23,
					Paths:      "0-6-7-8-23-24",
					Name:       "",
					Title:      "菜单列表",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/menu/",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         25,
					ParentId:   23,
					Paths:      "0-6-7-8-23-25",
					Name:       "",
					Title:      "菜单详情",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/menu/:id",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         26,
					ParentId:   23,
					Paths:      "0-6-7-8-23-26",
					Name:       "",
					Title:      "菜单新增",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/menu/",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         27,
					ParentId:   23,
					Paths:      "0-6-7-8-23-27",
					Name:       "",
					Title:      "菜单更新",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/menu/:id",
					Component:  "",
					Action:     "PUT",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         28,
					ParentId:   23,
					Paths:      "0-6-7-8-23-28",
					Name:       "",
					Title:      "菜单删除",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/menu/:id",
					Component:  "",
					Action:     "DELETE",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         29,
					ParentId:   8,
					Paths:      "0-6-7-8-29",
					Name:       "",
					Title:      "角色管理",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         30,
					ParentId:   29,
					Paths:      "0-6-7-8-29-30",
					Name:       "",
					Title:      "角色列表",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/role/",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         31,
					ParentId:   29,
					Paths:      "0-6-7-8-29-31",
					Name:       "",
					Title:      "角色列表",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/role/list",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         32,
					ParentId:   29,
					Paths:      "0-6-7-8-29-32",
					Name:       "",
					Title:      "角色展示",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/role/:id",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         33,
					ParentId:   29,
					Paths:      "0-6-7-8-29-33",
					Name:       "",
					Title:      "角色新增",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/role/",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         34,
					ParentId:   29,
					Paths:      "0-6-7-8-29-34",
					Name:       "",
					Title:      "角色更新",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/role/:id",
					Component:  "",
					Action:     "PUT",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         35,
					ParentId:   29,
					Paths:      "0-6-7-8-29-35",
					Name:       "",
					Title:      "角色状态更新",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/role/:id/status/",
					Component:  "",
					Action:     "PUT",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         36,
					ParentId:   29,
					Paths:      "0-6-7-8-29-36",
					Name:       "",
					Title:      "角色删除",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/role/:id",
					Component:  "",
					Action:     "DELETE",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         37,
					ParentId:   29,
					Paths:      "0-6-7-8-29-37",
					Name:       "",
					Title:      "批量角色删除",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/role/",
					Component:  "",
					Action:     "DELETE",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         38,
					ParentId:   8,
					Paths:      "0-6-7-8-38",
					Name:       "",
					Title:      "网站设置",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         39,
					ParentId:   38,
					Paths:      "0-6-7-8-38-39",
					Name:       "",
					Title:      "网站设置详情",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/website/info",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         40,
					ParentId:   38,
					Paths:      "0-6-7-8-38-40",
					Name:       "",
					Title:      "网站设置更新",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/website/info",
					Component:  "",
					Action:     "PUT",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         41,
					ParentId:   38,
					Paths:      "0-6-7-8-38-41",
					Name:       "",
					Title:      "网站设置上传",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/website/upload",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         42,
					ParentId:   1,
					Paths:      "0-1-42",
					Name:       "",
					Title:      "操作日志",
					Icon:       "log",
					Type:       1,
					Path:       "operation_log",
					Component:  "/admin/operation_log",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         43,
					ParentId:   0,
					Paths:      "0-43",
					Name:       "",
					Title:      "关于我们",
					Icon:       "introduction",
					Type:       0,
					Path:       "/about",
					Component:  "Layout",
					Action:     "",
					Permission: "",
					Sort:       20,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         44,
					ParentId:   43,
					Paths:      "0-43-44",
					Name:       "",
					Title:      "公司介绍",
					Icon:       "introduction",
					Type:       1,
					Path:       "company",
					Component:  "/about/company_introduction",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         45,
					ParentId:   0,
					Paths:      "0-45",
					Name:       "",
					Title:      "轮播图管理",
					Icon:       "banner",
					Type:       0,
					Path:       "/banner",
					Component:  "Layout",
					Action:     "",
					Permission: "",
					Sort:       40,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         46,
					ParentId:   45,
					Paths:      "0-45-46",
					Name:       "",
					Title:      "轮播图管理",
					Icon:       "banner",
					Type:       1,
					Path:       "index",
					Component:  "/banner/index",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         47,
					ParentId:   0,
					Paths:      "0-47",
					Name:       "",
					Title:      "视频管理",
					Icon:       "eye-open",
					Type:       0,
					Path:       "/video",
					Component:  "Layout",
					Action:     "",
					Permission: "",
					Sort:       30,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         48,
					ParentId:   47,
					Paths:      "0-47-48",
					Name:       "",
					Title:      "视频",
					Icon:       "eye-open",
					Type:       1,
					Path:       "index",
					Component:  "/video/index",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    1,
					IsLink:     0,
				},
				{
					ID:         49,
					ParentId:   7,
					Paths:      "0-6-7-49",
					Name:       "",
					Title:      "轮播图管理API",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         50,
					ParentId:   7,
					Paths:      "0-6-7-50",
					Name:       "",
					Title:      "视频管理API",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         51,
					ParentId:   7,
					Paths:      "0-6-7-51",
					Name:       "",
					Title:      "关于我们API",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         52,
					ParentId:   49,
					Paths:      "0-6-7-49-52",
					Name:       "",
					Title:      "轮播图列表",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/banner/",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         53,
					ParentId:   49,
					Paths:      "0-6-7-49-53",
					Name:       "",
					Title:      "创建轮播图",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/banner/",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         54,
					ParentId:   49,
					Paths:      "0-6-7-49-54",
					Name:       "",
					Title:      "轮播图详情",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/banner/:id",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         55,
					ParentId:   49,
					Paths:      "0-6-7-49-55",
					Name:       "",
					Title:      "更新轮播图",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/banner/:id",
					Component:  "",
					Action:     "PUT",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         56,
					ParentId:   49,
					Paths:      "0-6-7-49-56",
					Name:       "",
					Title:      "上传轮播图",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/banner/upload",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         57,
					ParentId:   50,
					Paths:      "0-6-7-50-57",
					Name:       "",
					Title:      "视频列表",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/video/list",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         58,
					ParentId:   50,
					Paths:      "0-6-7-50-58",
					Name:       "",
					Title:      "创建视频",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/video/",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         59,
					ParentId:   50,
					Paths:      "0-6-7-50-59",
					Name:       "",
					Title:      "更新视频",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/video/:id",
					Component:  "",
					Action:     "PUT",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         60,
					ParentId:   50,
					Paths:      "0-6-7-50-60",
					Name:       "",
					Title:      "删除视频",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/video/:id",
					Component:  "",
					Action:     "DELETE",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         61,
					ParentId:   50,
					Paths:      "0-6-7-50-61",
					Name:       "",
					Title:      "上传视频",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/video/upload",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         62,
					ParentId:   51,
					Paths:      "0-6-7-51-62",
					Name:       "",
					Title:      "查看公司介绍",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/company/introduction",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         63,
					ParentId:   51,
					Paths:      "0-6-7-51-63",
					Name:       "",
					Title:      "更新公司介绍",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/company/introduction",
					Component:  "",
					Action:     "PUT",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         64,
					ParentId:   51,
					Paths:      "0-6-7-51-64",
					Name:       "",
					Title:      "上传图片",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/company/upload",
					Component:  "",
					Action:     "POST",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         65,
					ParentId:   8,
					Paths:      "0-6-7-8-65",
					Name:       "",
					Title:      "操作日志管理",
					Icon:       "example",
					Type:       0,
					Path:       "/",
					Component:  "/",
					Action:     "",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         66,
					ParentId:   65,
					Paths:      "0-6-7-8-65-66",
					Name:       "",
					Title:      "操作日志列表",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/operation/logs/",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
				{
					ID:         67,
					ParentId:   65,
					Paths:      "0-6-7-8-65-67",
					Name:       "",
					Title:      "操作日志详情",
					Icon:       "example",
					Type:       2,
					Path:       "/admin/api/v1/operation/logs/:id",
					Component:  "",
					Action:     "GET",
					Permission: "",
					Sort:       0,
					Visible:    0,
					IsLink:     0,
				},
			}

			for _, m := range menus {
				err := tx.Create(&m).Error
				if err != nil {
					return err
				}
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTableIfExists(&AdminRoleMenu{}).Error
		},
	})
}