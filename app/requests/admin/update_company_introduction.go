package admin_request

import (
	"fagin/pkg/request"
)

type updateCompanyIntroduction struct {
	Name               string `form:"name" json:"name" binding:"required,max=32"`
	Image              string `form:"image" json:"image" binding:"required,max=255"`
	Content            string `form:"content" json:"content" binding:"required"`
	Status             *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	request.Validation `binding:"-"`
}

func NewUpdateCompanyIntroduction() *updateCompanyIntroduction {
	r := new(updateCompanyIntroduction)
	r.SetRequest(r)
	return r
}

func (*updateCompanyIntroduction) Message() map[string]string {
	return map[string]string{}
}

func (*updateCompanyIntroduction) Attributes() map[string]string {
	return map[string]string{
		"Name":    "公司名称",
		"Image":   "图片",
		"Content": "内容",
		"Status":  "状态",
	}
}
