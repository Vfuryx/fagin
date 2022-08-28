package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type ResetAdminUser struct {
	Password string `form:"password" json:"password" binding:"required,min=6"`
}

var _ request.Request = ResetAdminUser{}

func (ResetAdminUser) Message() map[string]string {
	return map[string]string{}
}

func (ResetAdminUser) Attributes() map[string]string {
	return map[string]string{
		"Password": "密码",
	}
}

func (r ResetAdminUser) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[ResetAdminUser](r, ctx)
}
