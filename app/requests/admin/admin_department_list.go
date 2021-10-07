package admin_request

import (
	"fagin/pkg/request"
)

type adminDepartmentList struct {
	Name   string `form:"name" json:"name" binding:"max=128"`
	Status *uint8 `form:"status" json:"status" binding:""`

	request.Validation `binding:"-"`
}

func NewAdminDepartmentList() *adminDepartmentList {
	r := new(adminDepartmentList)
	r.SetRequest(r)
	return r
}

func (*adminDepartmentList) Message() map[string]string {
	return map[string]string{}
}

func (*adminDepartmentList) Attributes() map[string]string {
	return map[string]string{
		"Name":   "部门名称",
		"Status": "部门状态",
	}
}

//func (r *adminDepartmentList) Validate(ctx *gin.Context) (map[string]string, bool) {
//	return request.Validated(r, ctx)
//}
