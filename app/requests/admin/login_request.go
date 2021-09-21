package admin_request

import (
	"fagin/pkg/request"
)

type loginRequest struct {
	Name               string `form:"username" json:"username" binding:"required"`
	Password           string `form:"password" json:"password" binding:"required"`
	ID                 string `form:"id" json:"id" binding:"required"`
	Code               string `form:"code" json:"code" binding:"required"`
	request.Validation `binding:"-"`
}

func NewLoginRequest() *loginRequest {
	r := new(loginRequest)
	r.SetRequest(r)
	return r
}
func (*loginRequest) Message() map[string]string {
	return map[string]string{}
}

func (*loginRequest) Attributes() map[string]string {
	return map[string]string{
		"Name":     "用户名",
		"Password": "密码",
		"ID":       "ID",
		"Code":     "验证码",
	}
}
