package admin_request

import (
	"fagin/pkg/request"
)

type roleKeyExistRequest struct {
	Key                string `form:"key" json:"key" binding:"required,max=32"`
	request.Validation `binding:"-"`
}

func NewRoleKeyExistRequest() *roleKeyExistRequest {
	r := new(roleKeyExistRequest)
	r.SetRequest(r)
	return r
}
func (*roleKeyExistRequest) Message() map[string]string {
	return map[string]string{}
}

func (*roleKeyExistRequest) Attributes() map[string]string {
	return map[string]string{
		"Key": "角色关键字",
	}
}
