package admin

import (
	"fagin/pkg/logger"
	"fagin/pkg/response"
	"github.com/gin-gonic/gin"
)

// BaseController 基础控制器
type BaseController struct {
}

func (ctr BaseController) ResponseJsonOK(ctx *gin.Context, data interface{}) {
	response.JsonOK(ctx, data)
}

func (ctr BaseController) ResponseJsonErr(ctx *gin.Context, err error, errs interface{}) {
	response.JsonErr(ctx, err, errs)
}

// ResponseJsonErrLog 处理错误并返回
func (ctr BaseController) ResponseJsonErrLog(ctx *gin.Context, err error, log interface{}) {
	response.JsonErr(ctx, err, nil).Log(logger.AdminMode, "\n", log)
}
