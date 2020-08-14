package admin_menu_type

import "fagin/pkg/constant"

const (
	// 主目录类型
	TypeMain uint8 = 0
	// 菜单类型
	TypeMenu uint8 = 1
	// 接口类型
	TypeApi uint8 = 2
	// 按钮类型
	TypeButton uint8 = 3
)

type adminMenuType struct {
	TypeMain   uint8
	TypeMenu   uint8
	TypeApi    uint8
	TypeButton uint8
	constant.Constant
}

func AdminMenuType() *adminMenuType {
	mt := adminMenuType{
		TypeMain:   TypeMain,
		TypeMenu:   TypeMenu,
		TypeApi:    TypeApi,
		TypeButton: TypeButton,
	}
	mt.Constant.C = &mt
	return &mt
}

func (s *adminMenuType) All() map[interface{}]string {
	return map[interface{}]string{
		s.TypeMain:   "主目录",
		s.TypeMenu:   "菜单",
		s.TypeApi:    "接口",
		s.TypeButton: "按钮",
	}
}
