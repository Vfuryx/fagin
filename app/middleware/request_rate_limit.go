package middleware

import (
	"github.com/gin-gonic/gin"
	"sync"
	"time"
)

// 限流中间件
// 使用计数器实现请求限流
type requestRateLimit struct {
	sync.Mutex
	Duration 	time.Duration	// 时间间隔
	MaxCount 	int				// 最大请求
	Count		int				// 当前请求数
}

func newRequestRateLimit(max int, duration time.Duration) *requestRateLimit {
	rrs := &requestRateLimit{
		Duration: duration,
		MaxCount: max,
	}
	// 间隔清空请求数
	go func() {
		tick := time.NewTicker(rrs.Duration)
		for {
			<-tick.C
			rrs.Lock()
			rrs.Count = 0
			rrs.Unlock()
		}
	}()
	return rrs
}

// 请求数加一
func (rrs *requestRateLimit) Increase() {
	rrs.Lock()
	rrs.Count++
	rrs.Unlock()
}

// 是否通过
func (rrs *requestRateLimit) IsAvailable() bool {
	rrs.Lock()
	defer rrs.Unlock()
	return rrs.Count <= rrs.MaxCount
}

// 限流中间件
func RateLimitMiddleware(max int, duration time.Duration, fun gin.HandlerFunc) gin.HandlerFunc {
	rrs := newRequestRateLimit(max, duration)
	return func(ctx *gin.Context) {
		rrs.Increase()
		if !rrs.IsAvailable() {
			fun(ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}