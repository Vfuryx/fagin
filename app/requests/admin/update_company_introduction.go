package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type UpdateCompanyIntroduction struct {
	Name    string `form:"name" json:"name" binding:"required,max=32"`
	Image   string `form:"image" json:"image" binding:"required,max=255"`
	Content string `form:"content" json:"content" binding:"required"`
	Status  *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

var _ request.Request = &UpdateCompanyIntroduction{}

func (UpdateCompanyIntroduction) FieldMap() map[string]string {
	return map[string]string{
		"Name":    "name",
		"Image":   "image",
		"Content": "content",
		"Status":  "status",
	}
}

func (UpdateCompanyIntroduction) Message() map[string]string {
	return map[string]string{
		"Name.required":    "公司名称不能为空",
		"Name.max":         "公司名称不能大于32位",
		"Image.required":   "图片不能为空",
		"Image.max":        "公司名称不能大于255位",
		"Content.required": "内容不能为空",
		"Status.required":  "状态不能为空",
		"Status.oneof":     "状态参数不正确",
	}
}

func (UpdateCompanyIntroduction) Attributes() map[string]string {
	return map[string]string{
		"Name":    "公司名称",
		"Image":   "图片",
		"Content": "内容",
		"Status":  "状态",
	}
}

func (r *UpdateCompanyIntroduction) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
