package admin

import (
	"fagin/app/enums"
	"fagin/app/errno"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type operationLogController struct {
	BaseController
}

// OperationLogController 操作日志控制器
var OperationLogController operationLogController

func (ctr *operationLogController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	r := adminRequest.NewAdminOperationLogList()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	params := gin.H{
		"orderBy": "id desc",
	}
	if r.Path != "" {
		params["like_path"] = "%" + r.Path + "%"
	}
	if r.Method != "" && r.Method != "ALL" {
		params["method"] = r.Method
	}
	if len(r.Time) >= 2 {
		params["start_time"] = r.Time[0]
		params["end_time"] = r.Time[1]
	}

	columns := []string{"id", "user", "method", "path", "ip", "operation", "created_at", "module"}

	logs, err := service.AdminOperationLog.List(params, columns, nil, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.InternalServerErr, err)
		return
	}

	data := response.OperationLog(logs...).Collection()
	ctr.ResponseJsonOK(ctx, gin.H{
		"items": data,
		"total": paginator.TotalCount,
	})
}

func (ctr *operationLogController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{
		"id", "user", "method", "path", "ip", "operation", "input", "module",
		"created_at", "user_agent", "latency_time", "location",
	}
	l, err := service.AdminOperationLog.ShowLog(id, columns)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxShowErr, err)
		return
	}
	ctr.ResponseJsonOK(ctx, gin.H{
		"id":           l.ID,
		"user":         l.User,
		"method":       l.Method,
		"path":         l.Path,
		"ip":           l.IP,
		"operation":    l.Operation,
		"input":        l.Input,
		"module":       l.Module,
		"created_at":   l.CreatedAt.Format(enums.TimeFormatDef.Get()),
		"user_agent":   l.UserAgent,
		"latency_time": l.LatencyTime,
		"location":     l.Location,
	})
	return
}
