package enums

import "fagin/pkg/enum"

// AdminMenuType 枚举类型
type AdminMenuType int

var _ enum.Enum = new(AdminMenuType)

const (
	AdminMenuTypeDir        AdminMenuType = 0
	AdminMenuTypeMenu       AdminMenuType = 1
	AdminMenuTypePermission AdminMenuType = 2
)

func (code AdminMenuType) String() string {
	switch code {
	case AdminMenuTypeDir:
		return "目录"
	case AdminMenuTypeMenu:
		return "菜单"
	case AdminMenuTypePermission:
		return "权限"
	}
	return ""
}

func (code AdminMenuType) Get() int {
	return int(code)
}

// AllAdminMenuType 所有
func AllAdminMenuType() map[int]string {
	return map[int]string{
		AdminMenuTypeDir.Get():        AdminMenuTypeDir.String(),
		AdminMenuTypeMenu.Get():       AdminMenuTypeMenu.String(),
		AdminMenuTypePermission.Get(): AdminMenuTypePermission.String(),
	}
}
