package api_request

import (
	"fagin/pkg/request"
)

type indexRequest struct {
	Name               string `form:"name" json:"name"  binding:"required,max=2"`
	Age                int    `form:"age" json:"age"  binding:"required"`
	request.Validation `binding:"-"`
}

func NewIndexRequest() *indexRequest {
	r := new(indexRequest)
	r.SetRequest(r)
	return r
}

func (*indexRequest) Message() map[string]string {
	return map[string]string{
		"Name.required": ":attribute不能为空",
		"Age.required":  ":attribute不能为空",
		"Name.max":      ":attribute不能大于:max个字符",
	}
}

func (*indexRequest) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
		"Age":  "年龄",
	}
}
