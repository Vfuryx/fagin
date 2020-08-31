package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type CreateAdminMenu struct {
	ParentId   uint   `form:"parent_id" json:"parent_id" binding:""`
	Name       string `form:"name" json:"name" binding:"required,max=32"`
	Title      string `form:"title" json:"title" binding:"required,max=32"`
	Icon       string `form:"icon" json:"icon" binding:"required,max=128"`
	Type       *uint8 `form:"type" json:"type" binding:"required,oneof=0 1 2 3"`
	Path       string `form:"path" json:"path" binding:"max=128"`
	Component  string `form:"component" json:"component" binding:"max=64"`
	Method     string `form:"method" json:"method" binding:"max=16"`
	Permission string `form:"permission" json:"permission" binding:"max=32"`
	Visible    *uint8 `form:"visible" json:"visible" binding:"required"`
	IsLink     *uint8 `form:"is_link" json:"is_link" binding:"required"`
	Sort       uint   `form:"sort" json:"sort" binding:""`
}

var _ request.Request = &CreateAdminMenu{}

func (CreateAdminMenu) Message() map[string]string {
	return map[string]string{
		//"Name.max":           "名称不能大于32位",
		//"Title.required":     "标题不能为空",
		//"Title.max":          "标题不能大于32位",
		//"Icon.required":      "图标不能为空",
		//"Icon.max":           "图标不能大于128位",
		//"Type.required":      "类型不能为空",
		//"Type.oneof":         "类型参数不正确",
		//"Path.required":      "路径不能为空",
		//"Path.max":           "路径不能大于128位",
		//"component.required": "组件路径不能为空",
		//"component.max":      "组件路径不能大于64位",
		//"Method.max":         "菜单请求方法不能大于16位",
		//"Permission.max":     "菜单请求方法不能大于32位",
		//"Visible.required":   "状态不能为空",
		//"IsLink.required":    "是否外链不能为空",
		//"Sort.required":      "排序不能为空",
	}
}

func (CreateAdminMenu) Attributes() map[string]string {
	return map[string]string{
		"ParentId":   "父ID",
		"Name":       "名称",
		"Title":      "标题",
		"Icon":       "icon",
		"Type":       "类型",
		"Path":       "路径",
		"Sort":       "排序",
		"Component":  "组件路径",
		"Method":     "菜单请求方法",
		"Permission": "权限",
		"Visible":    "状态",
		"IsLink":     "是否外链",
	}
}

func (r *CreateAdminMenu) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
