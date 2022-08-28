package request

import (
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type UpdateUserRequest struct {
	// 为啥要 * 号呢，防止零值覆盖，现在是空为 nil
	Password *string `form:"password" json:"password" binding:""`
	Status   *uint8  `form:"status" json:"status" binding:"min=0,max=1"`
}

var _ request.Request = UpdateUserRequest{}

func (UpdateUserRequest) Message() map[string]string {
	return map[string]string{
		// "Status.min": ":attribute参数错误",
		// "Status.max": ":attribute参数错误",
	}
}

func (UpdateUserRequest) Attributes() map[string]string {
	return map[string]string{
		"Password": "密码",
		"Status":   "状态",
	}
}

func (r UpdateUserRequest) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[UpdateUserRequest](r, ctx)
}
