package middleware

import (
	"fagin/app/service"
	"fagin/pkg/response"
	"github.com/gin-gonic/gin"
)

type auth struct{}

var Auth auth

// IsLogin 验证用户是否登录
func (*auth) IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		c, err := service.Token.ParseRequest(ctx)

		if err != nil {
			response.JsonErr(ctx, err, nil)
			ctx.Abort()
			return
		}

		ctx.Set("user_id", c.ID)

		ctx.Next()
	}
}

// AuthCheckRole 验证用户是否有权限访问
func (*auth) AuthCheckRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		//role := strconv.Itoa(int(service.Login.ID))
		//
		//roles := service.Casbin.GetRolesForUser(role)
		//
		//ok := service.Casbin.CheckRoles(roles, ctx.Request.URL.Path, ctx.Request.Method)
		//
		//if ok {
		//	ctx.Next()
		//} else {
		//	app.ResponseJson(ctx, errno.MidPermissionDeniedErr, nil)
		//	ctx.Abort()
		//}
		//return
	}
}
