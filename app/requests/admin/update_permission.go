package admin_request

import (
	"fagin/pkg/request"
)

type updatePermission struct {
	Type               uint8  `form:"type" json:"type" binding:"oneof=0 1 2"`
	GID                uint   `form:"gid" json:"gid" binding:"required,min=1"`
	Name               string `form:"name" json:"name" binding:"max=35"`
	Path               string `form:"paths" json:"paths" binding:"max=255"`
	Method             string `form:"method" json:"method" binding:"max=255"`
	Sort               uint   `form:"sort" json:"sort" binding:""`
	Status             *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	request.Validation `binding:"-"`
}

func NewUpdatePermission() *updatePermission {
	r := new(updatePermission)
	r.Request = r
	return r
}

func (updatePermission) Message() map[string]string {
	return map[string]string{

	}
}

func (updatePermission) Attributes() map[string]string {
	return map[string]string{
		"Name":   "名称",
		"GID":    "权限分组",
		"Path":   "路径",
		"Method": "请求方法",
		"Status": "状态",
		"Sort":   "排序",
	}
}
