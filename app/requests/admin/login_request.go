package request

import (
	"fagin/pkg/request"
)

type LoginRequest struct {
	Name     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	ID       string `form:"id" json:"id" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`

	request.Validation `binding:"-"`
}

func NewLoginRequest() *LoginRequest {
	r := new(LoginRequest)
	r.SetRequest(r)

	return r
}
func (*LoginRequest) Message() map[string]string {
	return map[string]string{}
}

func (*LoginRequest) Attributes() map[string]string {
	return map[string]string{
		"Name":     "用户名",
		"Password": "密码",
		"ID":       "ID",
		"Code":     "验证码",
	}
}
