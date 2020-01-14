package api_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	// 为啥要 * 号呢，防止零值覆盖，现在是空为 nil
	Password *string `form:"password" json:"password" binding:""`
	Status   *uint8  `form:"status" json:"status" binding:"min=0,max=1"`
}

var _ request.Request = &UpdateUserRequest{}

func (UpdateUserRequest) FieldMap() map[string]string {
	return map[string]string{
		"Password": "password",
		"Status":   "status",
	}
}

func (UpdateUserRequest) Message() map[string]string {
	return map[string]string{
		"Status.min": "状态参数错误",
		"Status.max": "状态参数错误",
	}
}

func (UpdateUserRequest) Attributes() map[string]string {
	return map[string]string{
		"Password": "密码",
		"Status":   "状态",
	}
}

func (r *UpdateUserRequest) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
