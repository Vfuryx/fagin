package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type AdminUserList struct {
	Username string `form:"name" json:"username" binding:"max=64"`
	Phone    string `form:"key" json:"phone" binding:"max=64"`
	Status   *uint8 `form:"status" json:"status" binding:""`
}

var _ request.Request = &AdminUserList{}

func (AdminUserList) FieldMap() map[string]string {
	return map[string]string{
		"Username": "username",
		"Phone":    "phone",
		"Status":   "status",
	}
}

func (AdminUserList) Message() map[string]string {
	return map[string]string{
		"Username.max": "用户名称不能大于128位",
		"Phone.max":    "手机号码不能大于128位",
	}
}

func (AdminUserList) Attributes() map[string]string {
	return map[string]string{
		"Username": "用户名称",
		"Phone":    "手机号码",
		"Status":   "状态",
	}
}

func (r *AdminUserList) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
