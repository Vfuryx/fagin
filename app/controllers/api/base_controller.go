package api

import (
	"fagin/app"
	"fagin/pkg/response"
	"github.com/gin-gonic/gin"
	"runtime/debug"
)

// BaseController 基础控制器
type BaseController struct {
}

func (bc BaseController) ResponseJsonOK(ctx *gin.Context, data interface{}) {
	response.JsonOK(ctx, data)
}

func (bc BaseController) ResponseJsonErr(ctx *gin.Context, err error, errors interface{}) {
	response.JsonErr(ctx, err, errors)
}

// ResponseJsonErrLog 处理错误并返回
func (bc BaseController) ResponseJsonErrLog(ctx *gin.Context, err error, errLog interface{}, errors interface{}) {
	go app.Log().Println(err, errLog, string(debug.Stack()))
	response.JsonErr(ctx, err, errors)
}
