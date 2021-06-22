package api_request

import (
	"fagin/pkg/request"
)

type indexRequest struct {
	Name               string `form:"name" json:"name"  binding:"required"`
	Age                int    `form:"age" json:"age"  binding:""`
	request.Validation `binding:"-"`
}

func NewIndexRequest() *indexRequest {
	r := new(indexRequest)
	r.Request = r
	return r
}

func (indexRequest) Message() map[string]string {
	return map[string]string{
		//"Name.required": "名称不能为空",
		//"Age.required":  "年龄不能为空",
	}
}

func (indexRequest) Attributes() map[string]string {
	return map[string]string{
		"Name": "名称",
		"Age":  "年龄",
	}
}
