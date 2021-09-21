package api_request

import (
	"fagin/pkg/request"
)

type addUserRequest struct {
	UserName           string `form:"username" json:"username"  binding:"required"`
	Password           string `form:"password" json:"password" binding:"required"`
	request.Validation `binding:"-"`
}

func NewAddUserRequest() *addUserRequest {
	r := new(addUserRequest)
	r.SetRequest(r)
	return r
}

func (*addUserRequest) Message() map[string]string {
	return map[string]string{
		//"UserName.required": ":attribute不能为空",
		//"Password.required": ":attribute不能为空",
	}
}

func (*addUserRequest) Attributes() map[string]string {
	return map[string]string{
		"UserName": "用户名",
		"Password": "密码",
	}
}
