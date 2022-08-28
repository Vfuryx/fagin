package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type RoleKeyExistRequest struct {
	Key string `form:"key" json:"key" binding:"required,max=32"`
}

var _ request.Request = RoleKeyExistRequest{}

func (RoleKeyExistRequest) Message() map[string]string {
	return map[string]string{}
}

func (RoleKeyExistRequest) Attributes() map[string]string {
	return map[string]string{
		"Key": "角色关键字",
	}
}

func (r RoleKeyExistRequest) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[RoleKeyExistRequest](r, ctx)
}
