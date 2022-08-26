package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type CreateStructure struct {
	Type          string `form:"type" json:"type" binding:"required,max=16"`
	Code          string `form:"code" json:"code" binding:"required,max=64"`
	Name          string `form:"name" json:"name" binding:"required,max=32"`
	PermissionIDs []uint `form:"permission_ids" json:"permission_ids" binding:"required,dive,required"`
	Sort          uint   `form:"sort" json:"sort" binding:""`
}

var _ request.Request = CreateStructure{}

func (CreateStructure) Message() map[string]string {
	return map[string]string{}
}

func (CreateStructure) Attributes() map[string]string {
	return map[string]string{
		"Type":          "类型",
		"Code":          "组件编码",
		"Name":          "组件名称",
		"PermissionIDs": "权限组",
		"Sort":          "排序",
	}
}

func (r CreateStructure) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[CreateStructure](r, ctx)
}
