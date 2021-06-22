package middleware

import "github.com/gin-gonic/gin"

// Limiter chan 实现的流控中间件
type Limiter struct {
	Token chan struct{}
	Max   int
}

func NewLimiter(num int) *Limiter {
	return &Limiter{
		Token: make(chan struct{}, num),
		Max:   num,
	}
}

func (l *Limiter) GetToken() bool {
	if len(l.Token) >= l.Max {
		return false
	}
	l.Token <- struct{}{}
	return true
}

func (l *Limiter) ReleaseToken() {
	<-l.Token
}

func LimitMiddleware(max int, fun gin.HandlerFunc) gin.HandlerFunc {
	l := NewLimiter(max)
	return func(ctx *gin.Context) {
		if ok := l.GetToken(); !ok {
			fun(ctx)
			ctx.Abort()
			return
		}
		ctx.Next()
		l.ReleaseToken()
		return
	}
}
