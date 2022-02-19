package request

import (
	"fagin/pkg/request"
)

type AdminUsernameExistRequest struct {
	ID       uint   `form:"id" json:"id" binding:"min=0"`
	Username string `form:"username" json:"username" binding:"required,max=32"`

	request.Validation `binding:"-"`
}

func NewAdminUsernameExistRequest() *AdminUsernameExistRequest {
	r := new(AdminUsernameExistRequest)
	r.SetRequest(r)

	return r
}

func (*AdminUsernameExistRequest) Message() map[string]string {
	return map[string]string{}
}

func (*AdminUsernameExistRequest) Attributes() map[string]string {
	return map[string]string{
		"ID":       "ID",
		"Username": "用户名",
	}
}
