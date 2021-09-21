package admin_request

import (
	"fagin/pkg/request"
)

type userList struct {
	Username           string `form:"name" json:"username" binding:"max=64"`
	Phone              string `form:"key" json:"phone" binding:"max=64"`
	Status             *uint8 `form:"status" json:"status" binding:""`
	request.Validation `binding:"-"`
}

func NewUserList() *userList {
	r := new(userList)
	r.SetRequest(r)
	return r
}

func (*userList) Message() map[string]string {
	return map[string]string{
		//"Username.max": "用户名称不能大于128位",
		//"Phone.max":    "手机号码不能大于128位",
	}
}

func (*userList) Attributes() map[string]string {
	return map[string]string{
		"Username": "用户名称",
		"Phone":    "手机号码",
		"Status":   "状态",
	}
}
