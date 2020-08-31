package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type CreateAdminRole struct {
	Name    string `form:"name" json:"name" binding:"required,max=128"`
	Key     string `form:"key" json:"key" binding:"required,max=32"`
	Remark  string `form:"remark" json:"remark" binding:"required,max=255"`
	Sort    uint   `form:"sort" json:"sort" binding:""`
	Status  *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	MenuIDs []uint `form:"menu_ids" json:"menu_ids" binding:"required,dive,required"`
}

var _ request.Request = &CreateAdminRole{}

func (CreateAdminRole) Message() map[string]string {
	return map[string]string{
		//"Name.required":   "角色名称不能为空",
		//"Name.max":        "角色名称不能大于128位",
		//"Key.required":    "角色关键字不能为空",
		//"Key.max":         "角色关键字不能大于128位",
		//"Remark.required": "角色备注不能为空",
		//"Remark.max":      "角色备注不能大于128位",
		//"Sort.required":   "排序不能为空",
		//"Status.required": "状态不能为空",
		//"Status.oneof":    "状态参数不正确",
	}
}

func (CreateAdminRole) Attributes() map[string]string {
	return map[string]string{
		"Name":    "角色名称",
		"Key":     "角色关键字",
		"Remark":  "角色备注",
		"Sort":    "排序",
		"Status":  "状态",
		"MenuIDs": "菜单组",
	}
}

func (r *CreateAdminRole) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
