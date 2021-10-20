package admin_request

import (
	"fagin/pkg/request"
)

type createVideo struct {
	Title              string `form:"title" json:"title" binding:"required"`
	Path               string `form:"path" json:"path" binding:"required"`
	Description        string `form:"description" json:"description" binding:"required"`
	Status             uint8  `form:"status" json:"status" binding:"min=0,max=1"`
	request.Validation `binding:"-"`
}

func NewCreateVideo() *createVideo {
	r := new(createVideo)
	r.SetRequest(r)
	return r
}

func (*createVideo) Message() map[string]string {
	return map[string]string{
		//"Title.required":       "标题不能为空",
		//"Path.required":        "路径不能为空",
		//"Description.required": "详情不能为空",
		//"Status.min":           "状态不正确",
		//"Status.max":           "状态不正确",
	}
}

func (*createVideo) Attributes() map[string]string {
	return map[string]string{
		"Title":       "标题",
		"Path":        "路径",
		"Description": "详情",
		"Status":      "状态",
	}
}
