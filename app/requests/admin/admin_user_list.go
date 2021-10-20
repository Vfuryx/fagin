package admin_request

import (
	"fagin/pkg/request"
)

type adminUserList struct {
	Username string `form:"username" json:"username" binding:"max=64"`
	Phone    string `form:"phone" json:"phone" binding:"max=11"`
	Status   *uint8 `form:"status" json:"status" binding:""`

	request.Validation `binding:"-"`
}

// NewAdminUserList 后台管理员列表查询
func NewAdminUserList() *adminUserList {
	r := new(adminUserList)
	r.SetRequest(r)
	return r
}

func (*adminUserList) Message() map[string]string {
	return map[string]string{}
}

func (*adminUserList) Attributes() map[string]string {
	return map[string]string{
		"Username": "用户名称",
		"Phone":    "手机号码",
		"Status":   "状态",
	}
}
