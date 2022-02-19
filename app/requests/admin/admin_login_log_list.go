package request

import (
	"fagin/pkg/request"
	"time"
)

type AdminLoginLogList struct {
	LoginName string      `form:"login_name" json:"login_name" binding:"omitempty,max=255"`
	Time      []time.Time `form:"time[]" json:"time" binding:"omitempty,dive,required" time_format:"2006-01-02 15:04:05"`

	request.Validation `binding:"-"`
}

func NewAdminLoginLogList() *AdminLoginLogList {
	r := new(AdminLoginLogList)
	r.SetRequest(r)

	return r
}

func (*AdminLoginLogList) Message() map[string]string {
	return map[string]string{}
}

func (*AdminLoginLogList) Attributes() map[string]string {
	return map[string]string{
		"LoginName": "登录名称",
		"Time":      "时间",
	}
}

// func (r *AdminLoginLogList) Validate(ctx *gin.Context) (map[string]string, bool) {
//	return request.Validated(r, ctx)
//}
