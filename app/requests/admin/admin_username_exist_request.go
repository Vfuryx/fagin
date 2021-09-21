package admin_request

import (
	"fagin/pkg/request"
)

type adminUsernameExistRequest struct {
	ID                 uint   `form:"id" json:"id" binding:"min=0"`
	Username           string `form:"username" json:"username" binding:"required,max=32"`
	request.Validation `binding:"-"`
}

func NewAdminUsernameExistRequest() *adminUsernameExistRequest {
	r := new(adminUsernameExistRequest)
	r.SetRequest(r)
	return r
}

func (*adminUsernameExistRequest) Message() map[string]string {
	return map[string]string{}
}

func (*adminUsernameExistRequest) Attributes() map[string]string {
	return map[string]string{
		"ID":       "ID",
		"Username": "用户名",
	}
}
