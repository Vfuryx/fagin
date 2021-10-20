package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type updateStructure struct {
	Type               string `form:"type" json:"type" binding:"required,max=16"`
	Code               string `form:"code" json:"code" binding:"required,max=64"`
	Name               string `form:"name" json:"name" binding:"required,max=32"`
	PermissionIDs      []uint `form:"permission_ids" json:"permission_ids" binding:"required,dive,required"`
	Sort               uint   `form:"sort" json:"sort" binding:""`
	request.Validation `binding:"-"`
}

func NewUpdateStructure() *updateStructure {
	r := new(updateStructure)
	r.SetRequest(r)
	return r
}

func (*updateStructure) Message() map[string]string {
	return map[string]string{}
}

func (*updateStructure) Attributes() map[string]string {
	return map[string]string{
		"SID":           "组件",
		"Type":          "类型",
		"Code":          "组件编码",
		"Name":          "组件名称",
		"PermissionIDs": "权限组",
		"Sort":          "排序",
	}
}

// 获取 uint 类型的 id
type updateStructureUri struct {
	ID                 uint `uri:"id" binding:"required,min=1"`
	SID                uint `uri:"sid" binding:"required,min=1"`
	request.Validation `binding:"-"`
}

func NewUpdateStructureUri() *updateStructureUri {
	r := new(updateStructureUri)
	r.SetRequest(r)
	return r
}

func (*updateStructureUri) Message() map[string]string {
	return map[string]string{}
}

func (*updateStructureUri) Attributes() map[string]string {
	return map[string]string{
		"ID":  "菜单",
		"SID": "组件",
	}
}

func (r *updateStructureUri) Validate(ctx *gin.Context) (map[string]string, bool) {
	return request.ValidateUri(r, ctx)
}
