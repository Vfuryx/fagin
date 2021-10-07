package admin_request

import (
	"fagin/pkg/request"
)

type createStructure struct {
	Type               string `form:"type" json:"type" binding:"required,max=16"`
	Code               string `form:"code" json:"code" binding:"required,max=64"`
	Name               string `form:"name" json:"name" binding:"required,max=32"`
	PermissionIDs      []uint `form:"permission_ids" json:"permission_ids" binding:"required,dive,required"`
	Sort               uint   `form:"sort" json:"sort" binding:""`
	request.Validation `binding:"-"`
}

func NewCreateStructure() *createStructure {
	r := new(createStructure)
	r.SetRequest(r)
	return r
}

func (*createStructure) Message() map[string]string {
	return map[string]string{}
}

func (*createStructure) Attributes() map[string]string {
	return map[string]string{
		"Type":          "类型",
		"Code":          "组件编码",
		"Name":          "组件名称",
		"PermissionIDs": "权限组",
		"Sort":          "排序",
	}
}