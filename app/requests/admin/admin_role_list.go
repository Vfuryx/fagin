package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type AdminRoleList struct {
	Name   string `form:"name" json:"name" binding:"max=128"`
	Key    string `form:"key" json:"key" binding:"max=32"`
	Status *uint8 `form:"status" json:"status" binding:""`
}

var _ request.Request = &AdminRoleList{}

func (AdminRoleList) Message() map[string]string {
	return map[string]string{
		//"Name.max":     "角色名称不能大于128位",
		//"Key.max":      "角色关键字不能大于128位",
		//"Status.oneof": "状态参数不正确",
	}
}

func (AdminRoleList) Attributes() map[string]string {
	return map[string]string{
		"Name":   "角色名称",
		"Key":    "角色关键字",
		"Status": "状态",
	}
}

func (r *AdminRoleList) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
