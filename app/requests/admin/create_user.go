package admin_request

import (
	"fagin/pkg/request"
)

type createUser struct {
	NickName           string `form:"nick_name" json:"nick_name" binding:"required,max=64"`
	Phone              string `form:"phone" json:"phone" binding:"required,max=64"`
	Email              string `form:"email" json:"email" binding:"required,max=64"`
	Username           string `form:"username" json:"username" binding:"required,max=64"`
	Sex                *uint8 `form:"sex" json:"sex" binding:"required,oneof=0 1 2"`
	Remark             string `form:"remark" json:"remark" binding:"required,max=255"`
	RoleID             uint   `form:"role_id" json:"role_id" binding:"required,min=1"`
	Status             *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
	Password           string `form:"password" json:"password" binding:"required,min=8"`
	request.Validation `binding:"-"`
}

func NewCreateUser() *createUser {
	r := new(createUser)
	r.Request = r
	return r
}

func (createUser) Message() map[string]string {
	return map[string]string{
		//"NickName.required": "昵称不能为空",
		//"NickName.max":      "昵称不能大于64位",
	}
}

func (createUser) Attributes() map[string]string {
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
