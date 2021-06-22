package web_request

import (
	"fagin/pkg/request"
)

type tagList struct {
	Tag                string `uri:"tag" binding:"required,max=32"`
	request.Validation `binding:"-"`
}

func NewTagList() *tagList {
	r := new(tagList)
	r.Request = r
	return r
}

func (tagList) Message() map[string]string {
	return map[string]string{
	}
}

func (tagList) Attributes() map[string]string {
	return map[string]string{
		"Tag": "名称",
	}
}
