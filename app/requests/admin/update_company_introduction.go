package request

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

var _ request.Request = UpdateCompanyIntroduction{}

func (UpdateCompanyIntroduction) Message() map[string]string {
	return map[string]string{}
}

func (UpdateCompanyIntroduction) Attributes() map[string]string {
	return map[string]string{
		"Name":    "公司名称",
		"Image":   "图片",
		"Content": "内容",
		"Status":  "状态",
	}
}

func (r UpdateCompanyIntroduction) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[UpdateCompanyIntroduction](r, ctx)
}
