package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type UpdateAdminRole struct {
	Remark  string `form:"remark" json:"remark" binding:"required,max=255"`
	Sort    uint   `form:"sort" json:"sort" binding:""`
	Status  *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	MenuIDs []uint `form:"menu_ids" json:"menu_ids" binding:"required,dive,required"`
}

var _ request.Request = &UpdateAdminRole{}

func (UpdateAdminRole) Message() map[string]string {
	return map[string]string{
		//"Remark.required": "角色备注不能为空",
		//"Remark.max":      "角色备注不能大于128位",
		//"Sort.required":   "排序不能为空",
		//"Status.required": "状态不能为空",
		//"Status.oneof":    "状态参数不正确",
	}
}

func (UpdateAdminRole) Attributes() map[string]string {
	return map[string]string{
		"Remark":  "角色备注",
		"Sort":    "排序",
		"Status":  "状态",
		"MenuIDs": "菜单组",
	}
}

func (r *UpdateAdminRole) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
