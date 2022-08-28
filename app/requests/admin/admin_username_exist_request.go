package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type AdminUsernameExistRequest struct {
	ID       uint   `form:"id" json:"id" binding:"min=0"`
	Username string `form:"username" json:"username" binding:"required,max=32"`
}

var _ request.Request = AdminUsernameExistRequest{}

func (AdminUsernameExistRequest) Message() map[string]string {
	return map[string]string{}
}

func (AdminUsernameExistRequest) Attributes() map[string]string {
	return map[string]string{
		"ID":       "ID",
		"Username": "用户名",
	}
}

func (r AdminUsernameExistRequest) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[AdminUsernameExistRequest](r, ctx)
}
