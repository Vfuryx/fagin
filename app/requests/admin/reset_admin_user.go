package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type ResetAdminUser struct {
	Password string `form:"password" json:"password" binding:"required,min=8"`
}

var _ request.Request = &ResetAdminUser{}

func (ResetAdminUser) FieldMap() map[string]string {
	return map[string]string{
		"Password": "password",
	}
}

func (ResetAdminUser) Message() map[string]string {
	return map[string]string{
		"Password.required": "密码不能为空",
		"Password.min":      "密码不能小于8位",
	}
}

func (ResetAdminUser) Attributes() map[string]string {
	return map[string]string{
		"Password": "密码",
	}
}

func (r *ResetAdminUser) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
