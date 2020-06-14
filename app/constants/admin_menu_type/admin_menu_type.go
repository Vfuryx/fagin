package admin_menu_type

import "fagin/pkg/constant"

const (
	// 主目录类型
	TypeMain int = 0
	// 菜单类型
	TypeMenu int = 1
	// 接口类型
	TypeApi int = 2
	// 按钮类型
	TypeButton int = 3
)

type adminMenuType struct {
	TypeMain   int
	TypeMenu   int
	TypeApi    int
	TypeButton int
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
