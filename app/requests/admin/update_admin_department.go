package request

import (
	"fagin/pkg/request"
)

type UpdateAdminDepartment struct {
	ParentID uint   `form:"parent_id" json:"parent_id" binding:"min=0"`
	Name     string `form:"name" json:"name" binding:"required,max=128"`
	Remark   string `form:"remark" json:"remark" binding:"max=255"`
	Sort     uint   `form:"sort" json:"sort" binding:""`
	Status   *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`

	request.Validation `binding:"-"`
}

func NewUpdateAdminDepartment() *UpdateAdminDepartment {
	r := new(UpdateAdminDepartment)
	r.SetRequest(r)

	return r
}

func (*UpdateAdminDepartment) Message() map[string]string {
	return map[string]string{}
}

func (*UpdateAdminDepartment) Attributes() map[string]string {
	return map[string]string{
		"ParentID": "父id",
		"Name":     "部门名称",
		"Remark":   "部门备注",
		"Sort":     "排序",
		"Status":   "状态",
	}
}

// func (r *UpdateAdminDepartment) Validate(ctx *gin.Context) (map[string]string, bool) {
//	return request.Validated(r, ctx)
//}
