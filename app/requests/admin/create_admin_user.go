package admin_request

import (
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type CreateAdminUser struct {
	NickName string `form:"nick_name" json:"nick_name" binding:"required,max=64"`
	Phone    string `form:"phone" json:"phone" binding:"required,max=64"`
	Email    string `form:"email" json:"email" binding:"required,max=64"`
	Username string `form:"username" json:"username" binding:"required,max=64"`
	Sex      *uint8 `form:"sex" json:"sex" binding:"required,oneof=0 1 2"`
	Remark   string `form:"remark" json:"remark" binding:"required,max=255"`
	RoleID   uint   `form:"role_id" json:"role_id" binding:"required,min=1"`
	Status   *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	Password string `form:"password" json:"password" binding:"required,min=8"`
}

var _ request.Request = &CreateAdminUser{}

func (CreateAdminUser) FieldMap() map[string]string {
	return map[string]string{
		"NickName": "nick_name",
		"Phone":    "phone",
		"Email":    "email",
		"Username": "username",
		"Sex":      "sex",
		"Remark":   "remark",
		"RoleID":   "role_id",
		"Status":   "status",
		"Password": "password",
	}
}

func (CreateAdminUser) Message() map[string]string {
	return map[string]string{
		"NickName.required": "昵称不能为空",
		"NickName.max":      "昵称不能大于64位",
		"Phone.required":    "手机号码不能为空",
		"Phone.max":         "手机号码不能大于64位",
		"Email.required":    "邮箱不能为空",
		"Email.max":         "邮箱不能大于64位",
		"Username.required": "用户名不能为空",
		"Username.max":      "用户名不能大于64位",
		"Sex.required":      "性别不能为空",
		"Sex.oneof":         "性别参数不正确",
		"Remark.required":   "备注不能为空",
		"Remark.max":        "备注不能大于255位",
		"RoleID.required":   "角色ID不能为空",
		"RoleID.min":        "角色ID不能小于1",
		"Status.required":   "状态不能为空",
		"Status.oneof":      "状态参数不正确",
		"Password.required": "密码不能为空",
		"Password.min":      "密码不能小于8位",
	}
}

func (CreateAdminUser) Attributes() map[string]string {
	return map[string]string{
		"NickName": "昵称",
		"Phone":    "电话",
		"Email":    "邮箱",
		"Username": "用户名",
		"Sex":      "性别",
		"Remark":   "备注",
		"RoleID":   "角色ID",
		"Status":   "状态",
		"Password": "密码",
	}
}

func (r *CreateAdminUser) Validate(ctx *gin.Context) (map[string]string, bool) {
	var v request.Validate
	return v.Validate(r, ctx)
}
