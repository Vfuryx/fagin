package request

import (
	"fagin/pkg/request"
)

type UpdateAdminUser struct {
	NickName string `form:"nick_name" json:"nick_name" binding:"required,max=64"`
	Username string `form:"username" json:"username" binding:"required,max=64"`
	Phone    string `form:"phone" json:"phone" binding:"required,max=64"`
	Email    string `form:"email" json:"email" binding:"required,max=64"`
	Sex      *uint8 `form:"sex" json:"sex" binding:"required,oneof=0 1 2"`
	Roles    []uint `form:"roles" json:"roles" binding:"dive,required"`
	Status   *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	Remark   string `form:"remark" json:"remark" binding:"max=255"`
	HomePath string `form:"home_path" json:"home_path" binding:"required,max=255"`

	request.Validation `binding:"-"`
}

func NewUpdateAdminUser() *UpdateAdminUser {
	r := new(UpdateAdminUser)
	r.SetRequest(r)

	return r
}

func (*UpdateAdminUser) Message() map[string]string {
	return map[string]string{}
}

func (*UpdateAdminUser) Attributes() map[string]string {
	return map[string]string{
		"NickName": "昵称",
		"Phone":    "电话",
		"Email":    "邮箱",
		"Username": "用户名",
		"Sex":      "性别",
		"Remark":   "备注",
		"Roles":    "角色组",
		"Status":   "状态",
		"Password": "密码",
	}
}
