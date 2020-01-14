package middleware

import (
	"bytes"
	"fagin/app/models/admin_operation_log"
	"github.com/gin-gonic/gin"
	"io/ioutil"
)

type adminOperationLog struct{}

var AdminOperationLog adminOperationLog

func (adminOperationLog) Operation() func(*gin.Context) {
	return func(ctx *gin.Context) {
		data, _ := ctx.GetRawData()
		ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(data)) // 关键点

		ctx.Next()

		operation := ctx.GetString("operation")
		if operation == "" {
			return
		}

		aol := admin_operation_log.AdminOperationLog{
			User:      ctx.GetString("user_name"),
			Method:    ctx.Request.Method,
			Path:      ctx.Request.RequestURI,
			IP:        ctx.ClientIP(),
			Operation: operation,
			Input:     string(data),
		}

		go func() {
			_ = aol.Dao().Create(&aol)
		}()
	}
}

func (adminOperationLog) Log(operation string) func(*gin.Context) {
	return func(ctx *gin.Context) {
		ctx.Set("operation", operation)
	}
}
