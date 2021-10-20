package middleware

import (
	"fagin/pkg/logger"
	"github.com/gin-gonic/gin"
	"net/http/httputil"
	"runtime/debug"
	"strings"
	"time"
)

// RecoveryLog 恢复日志
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
	prefix := "Authorization:"
	lines := strings.Split(string(httpRequest), "\r\n")
	for idx, header := range lines {
		if strings.HasPrefix(header, prefix) {
			lines[idx] = header[:len(prefix)] + " *"
			break
		}
	}

	go logger.Log().Errorf(
		"[Recovery] %s panic recovered:\n%s\n%s\n%s\n",
		time.Now(),
		strings.Join(lines, "\r\n"),
		r,
		s,
	)
}
