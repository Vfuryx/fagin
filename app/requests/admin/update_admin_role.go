package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type UpdateAdminRole struct {
	Type uint8  `gorm:"type:tinyint(4) unsigned;not null;default:0;comment:'菜单类型 0:总后台 1:商家后台 2:集团后台';column:type;"`
	Name string `form:"name" json:"name" binding:"required,max=128"`
	// Key                string `form:"key" json:"key" binding:"required,max=32"`
	Remark  string `form:"remark" json:"remark" binding:"max=255"`
	Sort    uint   `form:"sort" json:"sort" binding:""`
	Status  *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	MenuIDs []uint `form:"menu_ids" json:"menu_ids" binding:"required,dive,required"`
}

var _ request.Request = UpdateAdminRole{}

func (UpdateAdminRole) Message() map[string]string {
	return map[string]string{}
}

func (UpdateAdminRole) Attributes() map[string]string {
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

func (r UpdateAdminRole) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[UpdateAdminRole](r, ctx)
}
