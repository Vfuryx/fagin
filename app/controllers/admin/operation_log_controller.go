package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/responses/admin_response"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/log"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type operationLogController struct{}

var OperationLogController operationLogController

func (operationLogController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	params := gin.H{
		"sort": "id desc",
	}
	columns := []string{"id", "user", "method", "path", "ip", "operation", "created_at"}

	logs, err := service.AdminOperationLog.List(params, columns, nil, &paginator)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.InternalServerError, nil)
		return
	}

	data := admin_response.OperationLog(logs...).Collection()
	app.JsonResponse(ctx, errno.OK, gin.H{
		"logs":      data,
		"paginator": paginator,
	})
}

func (operationLogController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrBind, nil)
		return
	}

	columns := []string{"id", "user", "method", "path", "ip", "operation", "input"}
	l, err := service.AdminOperationLog.ShowLog(id, columns)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.InternalServerError, nil)
		return
	}
	app.JsonResponse(ctx, errno.OK, gin.H{
		"id":        l.ID,
		"user":      l.User,
		"method":    l.Method,
		"path":      l.Path,
		"ip":        l.IP,
		"operation": l.Operation,
		"input":     l.Input,
	})
	return
}
