package request

import (
	"fagin/pkg/request"
)

type UpdateAdminMenu struct {
	ParentID      uint   `form:"parent_id" json:"parent_id" binding:""`
	Type          int    `form:"type" json:"type" binding:"oneof=0 1 2"`
	Path          string `form:"path" json:"path" binding:"required,max=255"`
	Name          string `form:"name" json:"name" binding:"max=32"`
	Component     string `form:"component" json:"component" binding:"max=255"`
	Redirect      string `form:"redirect" json:"redirect" binding:"max=255"`
	FrameSrc      string `form:"frame_src" json:"frame_src" binding:"max=255"`
	CurrentActive string `form:"current_active" json:"current_active" binding:"max=255"`
	Title         string `form:"title" json:"title" binding:"required,max=32"`
	Icon          string `form:"icon" json:"icon" binding:"max=128"`
	Method        string `form:"method" json:"method" binding:"omitempty,oneof=POST GET PUT PATCH DELETE"`
	Permission    string `form:"permission" json:"permission" binding:"max=32"`
	Sort          uint   `form:"sort" json:"sort" binding:""`
	IsShow        *uint8 `form:"is_show" json:"is_show" binding:"required,oneof=0 1"`
	IsHideChild   *uint8 `form:"is_hide_child" json:"is_hide_child" binding:"required,oneof=0 1"`
	IsNoCache     *uint8 `form:"is_no_cache" json:"is_no_cache" binding:"required,oneof=0 1"`
	Status        *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`

	request.Validation `binding:"-"`
}

func NewUpdateAdminMenu() *UpdateAdminMenu {
	r := new(UpdateAdminMenu)
	r.SetRequest(r)

	return r
}

func (*UpdateAdminMenu) Message() map[string]string {
	return map[string]string{}
}

func (*UpdateAdminMenu) Attributes() map[string]string {
	return map[string]string{
		"Type":          "类型",
		"ParentId":      "父ID",
		"Name":          "名称",
		"Component":     "组件",
		"Title":         "标题",
		"Icon":          "icon",
		"Path":          "路径",
		"Sort":          "排序",
		"Permission":    "权限",
		"IsShow":        "状态",
		"Redirect":      "重定向",
		"FrameSrc":      "内嵌iframe的地址",
		"CurrentActive": "当前激活的菜单",
		"Status":        "状态",
		"IsHideChild":   "是否隐藏子菜单",
		"IsNoCache":     "是否无缓存",
		"Method":        "菜单请求方法",
	}
}
