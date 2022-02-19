package request

import (
	"fagin/pkg/request"
)

type CategoryList struct {
	Cate string `uri:"cate" binding:"required,max=32"`

	request.Validation `binding:"-"`
}

func NewCategoryList() *CategoryList {
	r := new(CategoryList)
	r.SetRequest(r)

	return r
}

func (*CategoryList) Message() map[string]string {
	return map[string]string{}
}

func (*CategoryList) Attributes() map[string]string {
	return map[string]string{
		"Cate": "名称",
	}
}
