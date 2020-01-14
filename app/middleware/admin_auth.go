package middleware

import (
	"fagin/app"
	"fagin/app/service/admin_auth"
	"github.com/gin-gonic/gin"
)

type adminAuth struct{}

var AdminAuth adminAuth

/**
 * 验证后台管理员是否登录
 */
func (adminAuth) IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		c, err := admin_auth.AdminAuth.ParseRequest(ctx)

		if err != nil {
			app.JsonResponse(ctx, err, nil)
			ctx.Abort()
			return
		}

		ctx.Set("user_name", c.Name)

		ctx.Next()
	}
}
