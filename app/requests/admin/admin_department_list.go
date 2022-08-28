package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type AdminDepartmentList struct {
	Name   string `form:"name" json:"name" binding:"max=128"`
	Status *uint8 `form:"status" json:"status" binding:""`
}

var _ request.Request = AdminDepartmentList{}

func (AdminDepartmentList) Message() map[string]string {
	return map[string]string{}
}

func (AdminDepartmentList) Attributes() map[string]string {
	return map[string]string{
		"Name":   "部门名称",
		"Status": "部门状态",
	}
}

func (r AdminDepartmentList) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[AdminDepartmentList](r, ctx)
}
