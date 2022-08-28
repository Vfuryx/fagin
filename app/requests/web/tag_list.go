package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type TagList struct {
	Tag string `uri:"tag" binding:"required,max=32"`
}

var _ request.Request = TagList{}

func (TagList) Message() map[string]string {
	return map[string]string{}
}

func (TagList) Attributes() map[string]string {
	return map[string]string{
		"Tag": "名称",
	}
}
func (r TagList) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[TagList](r, ctx)
}
