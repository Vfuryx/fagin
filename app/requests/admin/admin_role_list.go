package request

import (
	"fagin/pkg/request"
)

type AdminRoleList struct {
	Type   uint8  `form:"type" json:"type" binding:""`
	Name   string `form:"name" json:"name" binding:"max=128"`
	Key    string `form:"key" json:"key" binding:"max=32"`
	Status *uint8 `form:"status" json:"status" binding:""`

	request.Validation `binding:"-"`
}

func NewAdminRoleList() *AdminRoleList {
	r := new(AdminRoleList)
	r.SetRequest(r)

	return r
}

func (*AdminRoleList) Message() map[string]string {
	return map[string]string{}
}

func (*AdminRoleList) Attributes() map[string]string {
	return map[string]string{
		"Type":   "类型",
		"Name":   "角色名称",
		"Key":    "角色关键字",
		"Status": "状态",
	}
}
