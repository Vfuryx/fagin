package admin_request

import (
	"fagin/pkg/request"
)

type adminMenuList struct {
	Title              string `form:"title" json:"title" binding:"max=128"`
	IsShow             *uint8 `form:"is_show" json:"is_show" binding:""`
	Type               uint8  `form:"type" json:"type" binding:""`
	request.Validation `binding:"-"`
}

func NewAdminMenuList() *adminMenuList {
	r := new(adminMenuList)
	r.Request = r
	return r
}

func (adminMenuList) Message() map[string]string {
	return map[string]string{
	}
}

func (adminMenuList) Attributes() map[string]string {
	return map[string]string{
		"Title":  "菜单名称",
		"IsShow": "展示",
		"Type":   "类型",
	}
}
