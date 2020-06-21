package middleware

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/service"
	"fagin/app/service/admin_auth"
	"github.com/gin-gonic/gin"
	"strconv"
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
		ctx.Set("admin_user_id", c.UserID)

		ctx.Next()
	}
}

//权限检查中间件
func (adminAuth) AuthCheckRole() gin.HandlerFunc {
	return func(c *gin.Context) {
		userID := c.GetInt64("admin_user_id")

		roles, err := service.Canbin.GetRolesForUser(strconv.FormatInt(userID, 10))
		if err != nil {
			app.JsonResponse(c, errno.Api.ErrAuthCheckRole, nil)
			c.Abort()
			return
		}

		ok, err := service.Canbin.CheckRoles(roles, c.Request.URL.Path, c.Request.Method)
		if ok && err == nil {
			c.Next()
		} else {
			app.JsonResponse(c, errno.Api.ErrAuthCheckRole, nil)
			c.Abort()
			return
		}
	}
}
