package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type CreateCategory struct {
	Name   string `form:"name" json:"name" binding:"required,max=32"`
	Sort   uint   `form:"sort" json:"sort" binding:"required"`
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

var _ request.Request = CreateCategory{}

func (CreateCategory) Message() map[string]string {
	return map[string]string{}
}

func (CreateCategory) Attributes() map[string]string {
	return map[string]string{
		"Name":   "名称",
		"Sort":   "排序",
		"Status": "状态",
	}
}

func (r CreateCategory) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[CreateCategory](r, ctx)
}
