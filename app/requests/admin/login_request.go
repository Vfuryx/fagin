package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Name     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	ID       string `form:"id" json:"id" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

var _ request.Request = LoginRequest{}

func (LoginRequest) Message() map[string]string {
	return map[string]string{}
}

func (LoginRequest) Attributes() map[string]string {
	return map[string]string{
		"Name":     "用户名",
		"Password": "密码",
		"ID":       "ID",
		"Code":     "验证码",
	}
}

func (r LoginRequest) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[LoginRequest](r, ctx)
}
