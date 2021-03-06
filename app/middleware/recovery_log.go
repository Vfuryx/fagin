package middleware

import (
	"fagin/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"runtime/debug"
	"strings"
	"time"
)

func RecoveryLog() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			top := recover()
			if top != nil {
				httpRequest, _ := httputil.DumpRequest(ctx.Request, false)
				go writeLog(httpRequest, top, debug.Stack())
				// 重新抛出错误
				panic(top)
			}
		}()
		ctx.Next()
	}
}

// 写入日志
func writeLog(httpRequest []byte, r interface{}, s []byte) {
	headers := strings.Split(string(httpRequest), "\r\n")

	for idx, header := range headers {
		current := strings.Split(header, ":")
		if current[0] == "Authorization" {
			headers[idx] = current[0] + ": *"
		}
	}

	logger.Log.Panicf(
		"[Recovery] %s panic recovered:\n%s\n%s\n%s\n",
		time.Now(),
		strings.Join(headers, "\r\n"),
		r,
		s,
	)
}
