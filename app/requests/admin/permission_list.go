package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type PermissionList struct {
	Name string `form:"name" json:"name" binding:"max=35"`
}

var _ request.Request = PermissionList{}

func (PermissionList) Message() map[string]string {
	return map[string]string{}
}

func (PermissionList) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
	}
}
func (r PermissionList) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[PermissionList](r, ctx)
}
