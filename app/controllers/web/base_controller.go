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

func (ctr BaseController) ResponseJsonOK(ctx *gin.Context, data interface{}) {
	response.JsonOK(ctx, data)
}

func (ctr BaseController) ResponseJsonErr(ctx *gin.Context, err error, errors interface{}) {
	response.JsonErr(ctx, err, errors)
}

// ResponseJsonErrLog 处理错误并返回
func (ctr BaseController) ResponseJsonErrLog(ctx *gin.Context, err error, log interface{}) {
	go app.Log().Error("响应信息", err.Error(), "\n", log)

	if errors.Is(err, errno.ReqErr) {
		response.JsonErr(ctx, err, nil)
	} else {
		// 数据未找到
		app.View(ctx, "web.site.404", gin.H{})
	}
}
