package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type RolesRoute struct {
	Roles []uint `form:"roles[]" json:"roles" binding:"required,dive,required"`
}

var _ request.Request = RolesRoute{}

func (RolesRoute) Message() map[string]string {
	return map[string]string{}
}

func (RolesRoute) Attributes() map[string]string {
	return map[string]string{
		"Roles": "角色组",
	}
}

func (r RolesRoute) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[RolesRoute](r, ctx)
}
