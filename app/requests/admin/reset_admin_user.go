package admin_request

import (
	"fagin/pkg/request"
)

type resetAdminUser struct {
	Password           string `form:"password" json:"password" binding:"required,min=8"`
	request.Validation `binding:"-"`
}

func NewResetAdminUser() *resetAdminUser {
	r := new(resetAdminUser)
	r.Request = r
	return r
}

func (resetAdminUser) Message() map[string]string {
	return map[string]string{
	}
}

func (resetAdminUser) Attributes() map[string]string {
	return map[string]string{
		"Password": "密码",
	}
}
