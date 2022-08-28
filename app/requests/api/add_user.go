package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type AddUserRequest struct {
	UserName string `form:"username" json:"username"  binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var _ request.Request = AddUserRequest{}

func (AddUserRequest) Message() map[string]string {
	return map[string]string{
		// "UserName.required": ":attribute不能为空",
		// "Password.required": ":attribute不能为空",
	}
}

func (AddUserRequest) Attributes() map[string]string {
	return map[string]string{
		"UserName": "用户名",
		"Password": "密码",
	}
}

func (r AddUserRequest) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[AddUserRequest](r, ctx)
}
