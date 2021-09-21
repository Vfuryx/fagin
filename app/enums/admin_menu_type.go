package enums

import "fagin/pkg/enum"

// AdminMenuType 枚举类型
type AdminMenuType int

var _ enum.Enum = new(AdminMenuType)

const (
	AdminMenuTypeAdmin  AdminMenuType = 0 // 总后台
	AdminMenuTypeSeller AdminMenuType = 1 // 商家后台
)

func (code AdminMenuType) String() string {
	switch code {
	case AdminMenuTypeAdmin:
		return "总后台"
	case AdminMenuTypeSeller:
		return "商家后台"
	}
	return ""
}

func (code AdminMenuType) Get() int {
	return int(code)
}

// AllAdminMenuType 所有
func AllAdminMenuType() map[string]int {
	return map[string]int{
		AdminMenuTypeAdmin.String():  AdminMenuTypeAdmin.Get(),
		AdminMenuTypeSeller.String(): AdminMenuTypeSeller.Get(),
	}
}
