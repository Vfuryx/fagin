package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type UpdateVideo struct {
	Title       string `form:"title" json:"title" binding:"required"`
	Path        string `form:"path" json:"path" binding:"required"`
	Description string `form:"description" json:"description" binding:"required"`
	Status      *uint8 `form:"status" json:"status" binding:"min=0,max=1"`
}

var _ request.Request = &UpdateVideo{}

func (UpdateVideo) FieldMap() map[string]string {
	return map[string]string{
		"Title":       "title",
		"Path":        "path",
		"Description": "description",
		"Status":      "status",
	}
}

func (UpdateVideo) Message() map[string]string {
	return map[string]string{
		"Title.required":       "标题不能为空",
		"Path.required":        "路径不能为空",
		"Description.required": "详情不能唯恐",
		"Status.min":           "状态不正确",
		"Status.max":           "状态不正确",
	}
}

func (UpdateVideo) Attributes() map[string]string {
	return map[string]string{
		"Title":       "标题",
		"Path":        "路径",
		"Description": "详情",
		"Status":      "状态",
	}
}

func (r *UpdateVideo) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
