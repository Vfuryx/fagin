package admin_request

import (
	"fagin/pkg/request"
)

type adminRoleList struct {
	Type               uint8  `form:"type" json:"type" binding:""`
	Name               string `form:"name" json:"name" binding:"max=128"`
	Key                string `form:"key" json:"key" binding:"max=32"`
	Status             *uint8 `form:"status" json:"status" binding:""`
	request.Validation `binding:"-"`
}

func NewAdminRoleList() *adminRoleList {
	r := new(adminRoleList)
	r.SetRequest(r)
	return r
}

func (*adminRoleList) Message() map[string]string {
	return map[string]string{}
}

func (*adminRoleList) Attributes() map[string]string {
	return map[string]string{
		"Type":   "类型",
		"Name":   "角色名称",
		"Key":    "角色关键字",
		"Status": "状态",
	}
}
