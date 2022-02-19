package middleware

import (
	"context"
	"fagin/app"
	"fagin/app/errno"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// ContextTimeout 超时控制（测试） （不可用不能用于生成）
func ContextTimeout(t time.Duration) func(c *gin.Context) {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(c.Request.Context(), t)
		defer cancel()

		c.Request = c.Request.WithContext(ctx)

		go c.Next()

		select {
		case <-ctx.Done():
			cancel()

			if err := ctx.Err(); err != nil {
				app.ResponseJSONWithStatus(c, http.StatusOK, errno.ReqRequestTimeoutErr, nil, nil)
				c.Abort()

				return
			}
		case <-c.Done():
		}
	}
}
