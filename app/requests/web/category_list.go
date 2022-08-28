package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type CategoryList struct {
	Cate string `uri:"cate" binding:"required,max=32"`
}

var _ request.Request = CategoryList{}

func (CategoryList) Message() map[string]string {
	return map[string]string{}
}

func (CategoryList) Attributes() map[string]string {
	return map[string]string{
		"Cate": "名称",
	}
}

func (r CategoryList) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[CategoryList](r, ctx)
}
