package request

import (
	"fagin/pkg/request"
)

type PermissionList struct {
	Name string `form:"name" json:"name" binding:"max=35"`

	request.Validation `binding:"-"`
}

func NewPermissionList() *PermissionList {
	r := new(PermissionList)
	r.SetRequest(r)

	return r
}
func (*PermissionList) Message() map[string]string {
	return map[string]string{}
}

func (*PermissionList) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
	}
}
