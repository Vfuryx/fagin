package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type AdminMenuList struct {
	Title  string `form:"title" json:"title" binding:"max=128"`
	Visible *uint8 `form:"visible" json:"visible" binding:""`
}

var _ request.Request = &AdminMenuList{}

func (AdminMenuList) Message() map[string]string {
	return map[string]string{
		//"Title.max":    "名称不能大于128位",
		//"Visible.oneof": "状态参数不正确",
	}
}

func (AdminMenuList) Attributes() map[string]string {
	return map[string]string{
		"Title":  "菜单名称",
		"Visible": "状态",
	}
}

func (r *AdminMenuList) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
