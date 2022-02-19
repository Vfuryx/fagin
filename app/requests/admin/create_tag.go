package request

import (
	"fagin/pkg/request"
)

type CreateTag struct {
	Name   string `form:"name" json:"name" binding:"required,max=32"`
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`

	request.Validation `binding:"-"`
}

func NewCreateTag() *CreateTag {
	r := new(CreateTag)
	r.SetRequest(r)

	return r
}

func (*CreateTag) Message() map[string]string {
	return map[string]string{}
}

func (*CreateTag) Attributes() map[string]string {
	return map[string]string{
		"Name":   "名称",
		"Status": "状态",
	}
}
