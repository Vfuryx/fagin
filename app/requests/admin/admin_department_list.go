package request

import (
	"fagin/pkg/request"
)

type AdminDepartmentList struct {
	Name   string `form:"name" json:"name" binding:"max=128"`
	Status *uint8 `form:"status" json:"status" binding:""`

	request.Validation `binding:"-"`
}

func NewAdminDepartmentList() *AdminDepartmentList {
	r := new(AdminDepartmentList)
	r.SetRequest(r)

	return r
}

func (*AdminDepartmentList) Message() map[string]string {
	return map[string]string{}
}

func (*AdminDepartmentList) Attributes() map[string]string {
	return map[string]string{
		"Name":   "部门名称",
		"Status": "部门状态",
	}
}

// func (r *AdminDepartmentList) Validate(ctx *gin.Context) (map[string]string, bool) {
//	return request.Validated(r, ctx)
//}
