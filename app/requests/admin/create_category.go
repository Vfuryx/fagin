package request

import (
	"fagin/pkg/request"
)

type CreateCategory struct {
	Name   string `form:"name" json:"name" binding:"required,max=32"`
	Sort   uint   `form:"sort" json:"sort" binding:"required"`
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`

	request.Validation `binding:"-"`
}

func NewCreateCategory() *CreateCategory {
	r := new(CreateCategory)
	r.SetRequest(r)

	return r
}

func (*CreateCategory) Message() map[string]string {
	return map[string]string{}
}

func (*CreateCategory) Attributes() map[string]string {
	return map[string]string{
		"Name":   "名称",
		"Sort":   "排序",
		"Status": "状态",
	}
}
