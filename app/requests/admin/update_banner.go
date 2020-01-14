package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type UpdateBanner struct {
	Title  string `form:"title" json:"title" binding:"required,max=32"`
	Banner string `form:"banner" json:"banner" binding:"required,max=255"`
	Path   string `form:"path" json:"path" binding:"required,max=255"`
	Sort   uint   `form:"sort" json:"sort" binding:"required"`
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

var _ request.Request = &UpdateBanner{}

func (UpdateBanner) FieldMap() map[string]string {
	return map[string]string{
		"Title":  "title",
		"Banner": "banner",
		"Path":   "path",
		"Sort":   "sort",
		"Status": "status",
	}
}

func (UpdateBanner) Message() map[string]string {
	return map[string]string{
		"Title.required":  "标题不能为空",
		"Title.max":       "标题不能大于32位",
		"Banner.required": "轮播图不能为空",
		"Banner.max":      "轮播图不能大于255位",
		"Path.required":   "路径不能为空",
		"Path.max":        "路径不能大于255位",
		"Sort.required":   "排序不能为空",
		"Status.required": "状态不能为空",
		"Status.oneof":    "状态参数不正确",
	}
}

func (UpdateBanner) Attributes() map[string]string {
	return map[string]string{
		"Title":  "标题",
		"Banner": "轮播图",
		"Path":   "路径",
		"Sort":   "排序",
		"Status": "状态",
	}
}

func (r *UpdateBanner) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}