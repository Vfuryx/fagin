package request

import (
	"fagin/pkg/request"
)

type PermissionGroupList struct {
	Name string `form:"name" json:"name" binding:"max=35"`
	Type uint8  `form:"type" json:"type" binding:""`

	request.Validation `binding:"-"`
}

func NewPermissionGroupList() *PermissionGroupList {
	r := new(PermissionGroupList)
	r.SetRequest(r)

	return r
}

func (*PermissionGroupList) Message() map[string]string {
	return map[string]string{}
}

func (*PermissionGroupList) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
		"Type": "类型",
	}
}
