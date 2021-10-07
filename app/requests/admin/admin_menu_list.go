package admin_request

import (
	"fagin/pkg/request"
)

type adminMenuList struct {
	Title              string `form:"title" json:"title" binding:"max=128"`
	Status             *uint8 `form:"status" json:"status" binding:""`
	request.Validation `binding:"-"`
}

func NewAdminMenuList() *adminMenuList {
	r := new(adminMenuList)
	r.SetRequest(r)
	return r
}

func (*adminMenuList) Message() map[string]string {
	return map[string]string{}
}

func (*adminMenuList) Attributes() map[string]string {
	return map[string]string{
		"Title":  "菜单名称",
		"Status": "状态",
	}
}
