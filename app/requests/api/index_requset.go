package api_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type IndexRequest struct {
	Name string `form:"name" json:"name"  binding:"required"`
	Age  int    `form:"age" json:"age"  binding:""`
}

var _ request.Request = &IndexRequest{}

func (IndexRequest) FieldMap() map[string]string {
	return map[string]string{
		"Name": "name",
		"Age":  "age",
	}
}

func (IndexRequest) Message() map[string]string {
	return map[string]string{
		"Name.required": "名称不能为空",
		"Age.required":  "年龄不能为空",
	}
}

func (IndexRequest) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
		"Age":  "年龄",
	}
}

func (indexRequest *IndexRequest) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(indexRequest, ctx)
}
