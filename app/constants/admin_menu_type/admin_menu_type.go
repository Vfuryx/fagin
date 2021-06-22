package admin_menu_type

import "fagin/pkg/constant"

const (
	// TypeAdmin 总后台
	TypeAdmin uint = 0
	// TypeSeller 商家后台
	TypeSeller uint = 1
)

type adminMenuType struct {
	TypeAdmin  uint
	TypeSeller uint
	constant.Constant
}

func AdminMenuType() *adminMenuType {
	mt := adminMenuType{
		TypeAdmin:  TypeAdmin,
		TypeSeller: TypeSeller,
	}
	mt.Constant.C = &mt
	return &mt
}

func (s *adminMenuType) All() map[interface{}]string {
	return map[interface{}]string{
		s.TypeAdmin:  "总后台",
		s.TypeSeller: "商家后台",
	}
}
