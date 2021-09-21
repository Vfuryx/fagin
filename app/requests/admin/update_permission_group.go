package admin_request

import (
	"fagin/pkg/request"
)

type updatePermissionGroup struct {
	Name               string `form:"name" json:"name" binding:"max=35"`
	Type               uint8  `form:"type" json:"type" binding:""`
	Sort               uint   `form:"sort" json:"sort" binding:""`
	request.Validation `binding:"-"`
}

func NewUpdatePermissionGroup() *updatePermissionGroup {
	r := new(updatePermissionGroup)
	r.SetRequest(r)
	return r
}

func (*updatePermissionGroup) Message() map[string]string {
	return map[string]string{}
}

func (*updatePermissionGroup) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
		"Type": "类型",
		"Sort": "排序",
	}
}
