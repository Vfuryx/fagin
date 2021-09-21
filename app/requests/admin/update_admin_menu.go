package admin_request

import (
	"fagin/pkg/request"
)

type updateAdminMenu struct {
	Type               *uint8 `form:"type" json:"type" binding:"required,oneof=0 1"`
	ParentId           uint   `form:"parent_id" json:"parent_id" binding:""`
	Path               string `form:"path" json:"path" binding:"required,max=128"`
	Name               string `form:"name" json:"name" binding:"required,max=32"`
	Component          string `form:"component" json:"component" binding:"required,max=32"`
	Redirect           string `form:"redirect" json:"redirect" binding:"max=128"`
	Target             string `form:"target" json:"target" binding:"max=128"`
	Title              string `form:"title" json:"title" binding:"required,max=32"`
	Icon               string `form:"icon" json:"icon" binding:"max=128"`
	IsShow             *uint8 `form:"is_show" json:"is_show" binding:"required,oneof=0 1"`
	Sort               uint   `form:"sort" json:"sort" binding:""`
	Status             *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	IsHideChild        *uint8 `form:"is_hide_child" json:"is_hide_child" binding:"required,oneof=0 1"`
	Method             string `form:"method" json:"method" binding:"max=16"`
	Permission         string `form:"permission" json:"permission" binding:"max=32"`
	request.Validation `binding:"-"`
}

func NewUpdateAdminMenu() *updateAdminMenu {
	r := new(updateAdminMenu)
	r.SetRequest(r)
	return r
}

func (*updateAdminMenu) Message() map[string]string {
	return map[string]string{}
}

func (*updateAdminMenu) Attributes() map[string]string {
	return map[string]string{
		"Type":        "类型",
		"ParentId":    "父ID",
		"Name":        "名称",
		"Component":   "组件",
		"Title":       "标题",
		"Icon":        "icon",
		"Path":        "路径",
		"Sort":        "排序",
		"Permission":  "权限",
		"IsShow":      "状态",
		"Redirect":    "重定向",
		"Target":      "目标",
		"Status":      "状态",
		"IsHideChild": "是否隐藏子菜单",
		"Method":      "菜单请求方法",
	}
}
