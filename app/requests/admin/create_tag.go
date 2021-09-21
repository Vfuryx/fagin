package admin_request

import (
	"fagin/pkg/request"
)

type createTag struct {
	Name               string `form:"name" json:"name" binding:"required,max=32"`
	Status             *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	request.Validation `binding:"-"`
}

func NewCreateTag() *createTag {
	r := new(createTag)
	r.SetRequest(r)
	return r
}

func (*createTag) Message() map[string]string {
	return map[string]string{}
}

func (*createTag) Attributes() map[string]string {
	return map[string]string{
		"Name":   "名称",
		"Status": "状态",
	}
}
