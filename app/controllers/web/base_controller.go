package web

import (
	"errors"
	"fagin/app"
	"fagin/app/errno"
	"fagin/pkg/response"

	"github.com/gin-gonic/gin"
)

// BaseController 基础控制器
type BaseController struct {
}

const DefaultLimit = 15

func (ctr BaseController) ResponseJSONOK(ctx *gin.Context, data interface{}) {
	response.JSONSuccess(ctx, data)
}

func (ctr BaseController) ResponseJSONErr(ctx *gin.Context, err error, errs interface{}) {
	response.JSONErr(ctx, err, errs)
}

// ResponseJSONErrLog 处理错误并返回
func (ctr BaseController) ResponseJSONErrLog(ctx *gin.Context, err error, log interface{}) {
	go app.Log().Error("响应信息", err.Error(), "\n", log)

	if errors.Is(err, errno.ReqErr) {
		response.JSONErr(ctx, err, nil)
	} else {
		// 数据未找到
		app.View(ctx, "web.site.404", gin.H{})
	}
}
