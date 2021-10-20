package admin_request

import (
	"fagin/pkg/request"
)

type permissionGroupList struct {
	Name               string `form:"name" json:"name" binding:"max=35"`
	Type               uint8  `form:"type" json:"type" binding:""`
	request.Validation `binding:"-"`
}

func NewPermissionGroupList() *permissionGroupList {
	r := new(permissionGroupList)
	r.SetRequest(r)
	return r
}

func (*permissionGroupList) Message() map[string]string {
	return map[string]string{}
}

func (*permissionGroupList) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
		"Type": "类型",
	}
}
