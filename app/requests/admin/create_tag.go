package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type CreateTag struct {
	Name   string `form:"name" json:"name" binding:"required,max=32"`
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

var _ request.Request = CreateTag{}

func (CreateTag) Message() map[string]string {
	return map[string]string{}
}

func (CreateTag) Attributes() map[string]string {
	return map[string]string{
		"Name":   "名称",
		"Status": "状态",
	}
}

func (r CreateTag) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[CreateTag](r, ctx)
}
