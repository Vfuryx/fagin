package api_request

import (
	"fagin/pkg/request"
)

type updateUserRequest struct {
	// 为啥要 * 号呢，防止零值覆盖，现在是空为 nil
	Password           *string `form:"password" json:"password" binding:""`
	Status             *uint8  `form:"status" json:"status" binding:"min=0,max=1"`
	request.Validation `binding:"-"`
}

func NewUpdateUserRequest() *updateUserRequest {
	r := new(updateUserRequest)
	r.Request = r
	return r
}

func (updateUserRequest) Message() map[string]string {
	return map[string]string{
		//"Status.min": "状态参数错误",
		//"Status.max": "状态参数错误",
	}
}

func (updateUserRequest) Attributes() map[string]string {
	return map[string]string{
		"Password": "密码",
		"Status":   "状态",
	}
}
