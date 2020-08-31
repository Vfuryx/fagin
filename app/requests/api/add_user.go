package api_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type AddUserRequest struct {
	UserName string `form:"username" json:"username"  binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

var _ request.Request = &AddUserRequest{}

func (AddUserRequest) Message() map[string]string {
	return map[string]string{
		//"UserName.required": "用户名不能为空",
		//"Password.required": "密码不能为空",
	}
}

func (AddUserRequest) Attributes() map[string]string {
	return map[string]string{
		"UserName": "用户名",
		"Password": "密码",
	}
}

func (addUserRequest *AddUserRequest) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(addUserRequest, ctx)
}
