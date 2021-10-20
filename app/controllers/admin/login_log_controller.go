package admin

import (
	"fagin/app"
	"fagin/app/errno"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type loginLogController struct {
	BaseController
}

// LoginLogController 登录日志控制器
var LoginLogController loginLogController

func (ctr *loginLogController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	r := adminRequest.NewAdminLoginLogList()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	params := gin.H{
		"orderBy": "id desc",
	}
	if r.LoginName != "" {
		params["like_login_name"] = "%" + r.LoginName + "%"
	}
	if len(r.Time) >= 2 {
		params["start_time"] = r.Time[0]
		params["end_time"] = r.Time[1]
	}

	columns := []string{"*"}

	logs, err := service.AdminLoginLog.List(params, columns, nil, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.InternalServerErr, err)
		return
	}

	data := response.LoginLog(logs...).Collection()
	ctr.ResponseJsonOK(ctx, gin.H{
		"items": data,
		"total": paginator.TotalCount,
	})
}

func (ctr *loginLogController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"*"}
	l, err := service.AdminLoginLog.ShowLog(id, columns)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxShowErr, err)
		return
	}
	ctr.ResponseJsonOK(ctx, gin.H{
		"id":         l.ID,
		"login_name": l.LoginName,
		"ip":         l.IP,
		"location":   l.Location,
		"browser":    l.Browser,
		"os":         l.OS,
		"created_at": app.TimeToStr(l.CreatedAt),
	})
	return
}
