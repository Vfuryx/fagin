package admin_request

import (
	"fagin/pkg/request"
	"time"
)

type adminOperationLogList struct {
	Path   string      `form:"path" json:"path" binding:"omitempty,max=255"`
	Method string      `form:"method" json:"method" binding:"omitempty,oneof= ALL GET POST PUT DELETE"`
	Time   []time.Time `form:"time[]" json:"time" binding:"omitempty,dive,required" time_format:"2006-01-02 15:04:05"`

	request.Validation `binding:"-"`
}

func NewAdminOperationLogList() *adminOperationLogList {
	r := new(adminOperationLogList)
	r.SetRequest(r)
	return r
}

func (*adminOperationLogList) Message() map[string]string {
	return map[string]string{}
}

func (*adminOperationLogList) Attributes() map[string]string {
	return map[string]string{
		"Path":   "路径",
		"Method": "方法",
		"Time":   "时间",
	}
}
