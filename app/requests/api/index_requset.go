package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type IndexRequest struct {
	Name string `form:"name" json:"name"  binding:"required,max=2"`
	Age  int    `form:"age" json:"age"  binding:"required"`
}

var _ request.Request = IndexRequest{}

func (IndexRequest) Message() map[string]string {
	return map[string]string{
		"Name.required": ":attribute不能为空",
		"Age.required":  ":attribute不能为空",
		"Name.max":      ":attribute不能大于:max个字符",
	}
}

func (IndexRequest) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
		"Age":  "年龄",
	}
}

func (r IndexRequest) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[IndexRequest](r, ctx)
}
