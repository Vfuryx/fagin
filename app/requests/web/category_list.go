package web_request

import (
	"fagin/pkg/request"
)

type categoryList struct {
	Cate               string `uri:"cate" binding:"required,max=32"`
	request.Validation `binding:"-"`
}

func NewCategoryList() *categoryList {
	r := new(categoryList)
	r.SetRequest(r)
	return r
}

func (*categoryList) Message() map[string]string {
	return map[string]string{}
}

func (*categoryList) Attributes() map[string]string {
	return map[string]string{
		"Cate": "名称",
	}
}
