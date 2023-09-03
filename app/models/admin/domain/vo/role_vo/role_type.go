package role_vo

import "errors"

// TypeVO 菜单类型
type TypeVO uint8

const (
	// RoleTypeAdmin  总后台
	RoleTypeAdmin uint8 = 0
	// RoleTypeTenant  商家后台
	RoleTypeTenant uint8 = 1
	// RoleTypeGroup  集团后台
	RoleTypeGroup uint8 = 2
)

var TypeVOMap = map[uint8]string{
	RoleTypeAdmin:  "总后台",
	RoleTypeTenant: "商家后台",
	RoleTypeGroup:  "集团后台",
}

func (r *TypeVO) Parse(t uint8) error {
	if _, ok := TypeVOMap[t]; !ok {
		return errors.New("菜单类型错误，请输入正确的类型[总后台]、[商家后台]、[商家后台]")
	}

	*r = TypeVO(t)

	return nil
}
