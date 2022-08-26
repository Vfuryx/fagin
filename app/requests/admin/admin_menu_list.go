package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type AdminMenuList struct {
	Title  string `form:"title" json:"title" binding:"max=128"`
	Status *uint8 `form:"status" json:"status" binding:""`
}

var _ request.Request = AdminMenuList{}

func (AdminMenuList) Message() map[string]string {
	return map[string]string{}
}

func (AdminMenuList) Attributes() map[string]string {
	return map[string]string{
		"Title":  "菜单名称",
		"Status": "状态",
	}
}

func (r AdminMenuList) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[AdminMenuList](r, ctx)
}
