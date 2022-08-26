package request

import (
	"fagin/pkg/request"
	"time"

	"github.com/gin-gonic/gin"
)

type AdminOperationLogList struct {
	Path   string      `form:"path" json:"path" binding:"omitempty,max=255"`
	Method string      `form:"method" json:"method" binding:"omitempty,oneof= ALL GET POST PUT DELETE"`
	Time   []time.Time `form:"time[]" json:"time" binding:"omitempty,dive,required" time_format:"2006-01-02 15:04:05"`
}

var _ request.Request = AdminOperationLogList{}

func (AdminOperationLogList) Message() map[string]string {
	return map[string]string{}
}

func (AdminOperationLogList) Attributes() map[string]string {
	return map[string]string{
		"Path":   "路径",
		"Method": "方法",
		"Time":   "时间",
	}
}

func (r AdminOperationLogList) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[AdminOperationLogList](r, ctx)
}
