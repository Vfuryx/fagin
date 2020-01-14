package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type UpdateAdminUser struct {
	Email       string `form:"email" json:"email" binding:"required,email"`
	OldPassword string `form:"old_password" json:"old_password" binding:"required_with=NewPassword,omitempty,min=8"`
	// 跳过空值验证
	NewPassword string `form:"new_password" json:"new_password" binding:"required_with=OldPassword,omitempty,min=8"`
}

var _ request.Request = &UpdateAdminUser{}

func (UpdateAdminUser) FieldMap() map[string]string {
	return map[string]string{
		"Email":       "email",
		"OldPassword": "old_password",
		"NewPassword": "new_password",
	}
}

func (UpdateAdminUser) Message() map[string]string {
	return map[string]string{
		"Email.required":            "邮件不能为空",
		"Email.email":               "邮件格式不正确",
		"OldPassword.required_with": "旧密码不能为空",
		"OldPassword.min":           "旧密码不能小于8位",
		"NewPassword.required_with": "新密码不能为空",
		"NewPassword.min":           "新密码不能小于8位",
	}
}

func (UpdateAdminUser) Attributes() map[string]string {
	return map[string]string{
		"Email":       "邮件",
		"OldPassword": "旧密码",
		"NewPassword": "新密码",
	}
}

func (r *UpdateAdminUser) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
