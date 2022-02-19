package request

import (
	"fagin/pkg/request"
)

type AdminMenuList struct {
	Title  string `form:"title" json:"title" binding:"max=128"`
	Status *uint8 `form:"status" json:"status" binding:""`

	request.Validation `binding:"-"`
}

func NewAdminMenuList() *AdminMenuList {
	r := new(AdminMenuList)
	r.SetRequest(r)

	return r
}

func (*AdminMenuList) Message() map[string]string {
	return map[string]string{}
}

func (*AdminMenuList) Attributes() map[string]string {
	return map[string]string{
		"Title":  "菜单名称",
		"Status": "状态",
	}
}
