package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type UpdateStructure struct {
	Type          string `form:"type" json:"type" binding:"required,max=16"`
	Code          string `form:"code" json:"code" binding:"required,max=64"`
	Name          string `form:"name" json:"name" binding:"required,max=32"`
	PermissionIDs []uint `form:"permission_ids" json:"permission_ids" binding:"required,dive,required"`
	Sort          uint   `form:"sort" json:"sort" binding:""`
}

var _ request.Request = UpdateStructure{}

func (UpdateStructure) Message() map[string]string {
	return map[string]string{}
}

func (UpdateStructure) Attributes() map[string]string {
	return map[string]string{
		"SID":           "组件",
		"Type":          "类型",
		"Code":          "组件编码",
		"Name":          "组件名称",
		"PermissionIDs": "权限组",
		"Sort":          "排序",
	}
}

func (r UpdateStructure) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[UpdateStructure](r, ctx)
}

// UpdateStructureURI  获取 uint 类型的 id
type UpdateStructureURI struct {
	ID  uint `uri:"id" binding:"required,min=1"`
	SID uint `uri:"sid" binding:"required,min=1"`
}

var _ request.Request = UpdateStructureURI{}

func (UpdateStructureURI) Message() map[string]string {
	return map[string]string{}
}

func (UpdateStructureURI) Attributes() map[string]string {
	return map[string]string{
		"ID":  "菜单",
		"SID": "组件",
	}
}

func (r UpdateStructureURI) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.ValidateURI[UpdateStructureURI](r, ctx)
}
