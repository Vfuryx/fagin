package enums

import (
	"fagin/pkg/enum"
)

// AdminPermissionGroup 枚举类型
type AdminPermissionGroup string

var _ enum.Enum = new(AdminPermissionGroup)

const (
	Admin     AdminPermissionGroup = "admin"     // 管理员管理
	Menu      AdminPermissionGroup = "menu"      // 菜单管理
	Operation AdminPermissionGroup = "operation" // 操作日志管理
)

func (code AdminPermissionGroup) String() string {
	switch code {
	case Admin:
		return "管理员管理"
	case Menu:
		return "菜单管理"
	case Operation:
		return "操作日志管理"
	}

	return ""
}

func (code AdminPermissionGroup) Get() string {
	return string(code)
}

// AllAdminPermissionGroup 所有
func AllAdminPermissionGroup() map[string]string {
	return map[string]string{
		Admin.Get(): Admin.String(),
		Menu.Get():  Menu.String(),
	}
}

func AdminPermissionGroupExists(code string) bool {
	_, ok := AllAdminPermissionGroup()[code]
	return ok
}
