package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type AdminUserList struct {
	Username string `form:"username" json:"username" binding:"max=64"`
	Phone    string `form:"phone" json:"phone" binding:"max=11"`
	Status   *uint8 `form:"status" json:"status" binding:""`
}

var _ request.Request = AdminUserList{}

func (AdminUserList) Message() map[string]string {
	return map[string]string{}
}

func (AdminUserList) Attributes() map[string]string {
	return map[string]string{
		"Username": "用户名称",
		"Phone":    "手机号码",
		"Status":   "状态",
	}
}

func (r AdminUserList) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[AdminUserList](r, ctx)
}
