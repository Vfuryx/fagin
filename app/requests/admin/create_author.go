package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type CreateAuthor struct {
	Name   string `form:"name" json:"name" binding:"required,max=32"`
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

var _ request.Request = CreateAuthor{}

func (CreateAuthor) Message() map[string]string {
	return map[string]string{}
}

func (CreateAuthor) Attributes() map[string]string {
	return map[string]string{
		"Name":   "名称",
		"Status": "状态",
	}
}

func (r CreateAuthor) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[CreateAuthor](r, ctx)
}
