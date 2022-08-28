package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type AdminRoleList struct {
	Type   uint8  `form:"type" json:"type" binding:""`
	Name   string `form:"name" json:"name" binding:"max=128"`
	Key    string `form:"key" json:"key" binding:"max=32"`
	Status *uint8 `form:"status" json:"status" binding:""`
}

var _ request.Request = AdminRoleList{}

func (AdminRoleList) Message() map[string]string {
	return map[string]string{}
}

func (AdminRoleList) Attributes() map[string]string {
	return map[string]string{
		"Type":   "类型",
		"Name":   "角色名称",
		"Key":    "角色关键字",
		"Status": "状态",
	}
}

func (r AdminRoleList) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[AdminRoleList](r, ctx)
}
