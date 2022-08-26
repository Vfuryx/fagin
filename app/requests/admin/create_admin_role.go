package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type CreateAdminRole struct {
	// Type    uint8  `form:"type" json:"type" binding:"required,min=1"`
	Name    string `form:"name" json:"name" binding:"required,max=128"`
	Key     string `form:"key" json:"key" binding:"required,max=32"`
	Remark  string `form:"remark" json:"remark" binding:"max=255"`
	Sort    uint   `form:"sort" json:"sort" binding:""`
	Status  *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	MenuIDs []uint `form:"menu_ids" json:"menu_ids" binding:"required,dive,required"`
}

var _ request.Request = CreateAdminRole{}

func (CreateAdminRole) Message() map[string]string {
	return map[string]string{}
}

func (CreateAdminRole) Attributes() map[string]string {
	return map[string]string{
		"Type":    "类型",
		"Name":    "角色名称",
		"Key":     "角色关键字",
		"Remark":  "角色备注",
		"Sort":    "排序",
		"Status":  "状态",
		"MenuIDs": "菜单组",
	}
}

func (r CreateAdminRole) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[CreateAdminRole](r, ctx)
}
