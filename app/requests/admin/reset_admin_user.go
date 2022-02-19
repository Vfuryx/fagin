package request

import (
	"fagin/pkg/request"
)

type ResetAdminUser struct {
	Password string `form:"password" json:"password" binding:"required,min=6"`

	request.Validation `binding:"-"`
}

func NewResetAdminUser() *ResetAdminUser {
	r := new(ResetAdminUser)
	r.SetRequest(r)

	return r
}

func (*ResetAdminUser) Message() map[string]string {
	return map[string]string{}
}

func (*ResetAdminUser) Attributes() map[string]string {
	return map[string]string{
		"Password": "密码",
	}
}
