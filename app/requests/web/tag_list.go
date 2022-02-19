package request

import (
	"fagin/pkg/request"
)

type TagList struct {
	Tag string `uri:"tag" binding:"required,max=32"`

	request.Validation `binding:"-"`
}

func NewTagList() *TagList {
	r := new(TagList)
	r.SetRequest(r)

	return r
}

func (*TagList) Message() map[string]string {
	return map[string]string{}
}

func (*TagList) Attributes() map[string]string {
	return map[string]string{
		"Tag": "名称",
	}
}
