package request

import (
	"fagin/pkg/request"
)

type AddUserRequest struct {
	UserName string `form:"username" json:"username"  binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`

	request.Validation `binding:"-"`
}

func NewAddUserRequest() *AddUserRequest {
	r := new(AddUserRequest)
	r.SetRequest(r)

	return r
}

func (*AddUserRequest) Message() map[string]string {
	return map[string]string{
		// "UserName.required": ":attribute不能为空",
		// "Password.required": ":attribute不能为空",
	}
}

func (*AddUserRequest) Attributes() map[string]string {
	return map[string]string{
		"UserName": "用户名",
		"Password": "密码",
	}
}
