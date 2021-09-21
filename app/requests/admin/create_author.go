package admin_request

import (
	"fagin/pkg/request"
)

type createAuthor struct {
	Name               string `form:"name" json:"name" binding:"required,max=32"`
	Status             *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	request.Validation `binding:"-"`
}

func NewCreateAuthor() *createAuthor {
	r := new(createAuthor)
	r.SetRequest(r)
	return r
}

func (*createAuthor) Message() map[string]string {
	return map[string]string{}
}

func (*createAuthor) Attributes() map[string]string {
	return map[string]string{
		"Name":   "名称",
		"Status": "状态",
	}
}
