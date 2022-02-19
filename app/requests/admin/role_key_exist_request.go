package request

import (
	"fagin/pkg/request"
)

type RoleKeyExistRequest struct {
	Key string `form:"key" json:"key" binding:"required,max=32"`

	request.Validation `binding:"-"`
}

func NewRoleKeyExistRequest() *RoleKeyExistRequest {
	r := new(RoleKeyExistRequest)
	r.SetRequest(r)

	return r
}
func (*RoleKeyExistRequest) Message() map[string]string {
	return map[string]string{}
}

func (*RoleKeyExistRequest) Attributes() map[string]string {
	return map[string]string{
		"Key": "角色关键字",
	}
}
