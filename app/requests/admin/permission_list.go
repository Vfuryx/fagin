package admin_request

import (
	"fagin/pkg/request"
)

type permissionList struct {
	Name               string `form:"name" json:"name" binding:"max=35"`
	request.Validation `binding:"-"`
}

func NewPermissionList() *permissionList {
	r := new(permissionList)
	r.SetRequest(r)
	return r
}
func (*permissionList) Message() map[string]string {
	return map[string]string{}
}

func (*permissionList) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
	}
}
