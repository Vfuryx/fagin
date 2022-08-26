package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type PermissionGroupList struct {
	Name string `form:"name" json:"name" binding:"max=35"`
	Type uint8  `form:"type" json:"type" binding:""`
}

var _ request.Request = PermissionGroupList{}

func (PermissionGroupList) Message() map[string]string {
	return map[string]string{}
}

func (PermissionGroupList) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
		"Type": "类型",
	}
}
func (r PermissionGroupList) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[PermissionGroupList](r, ctx)
}
