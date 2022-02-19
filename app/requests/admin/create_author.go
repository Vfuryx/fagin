package request

import (
	"fagin/pkg/request"
)

type CreateAuthor struct {
	Name   string `form:"name" json:"name" binding:"required,max=32"`
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`

	request.Validation `binding:"-"`
}

func NewCreateAuthor() *CreateAuthor {
	r := new(CreateAuthor)
	r.SetRequest(r)

	return r
}

func (*CreateAuthor) Message() map[string]string {
	return map[string]string{}
}

func (*CreateAuthor) Attributes() map[string]string {
	return map[string]string{
		"Name":   "名称",
		"Status": "状态",
	}
}
