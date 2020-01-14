package middleware

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/service"
	"github.com/gin-gonic/gin"
	"strconv"
)

type auth struct{}

var Auth = auth{}

/**
 * 验证用户是否登录
 */
func (auth) IsLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		c, err := service.Token.ParseRequest(ctx)

		if err != nil {
			app.JsonResponse(ctx, err, nil)
			ctx.Abort()
			return
		}

		ctx.Set("user_id", c.ID)

		ctx.Next()
	}
}

/*
 * 验证用户是否有权限访问
 */
func (auth) AuthCheckRole() gin.HandlerFunc {
	return func(ctx *gin.Context) {

		role := strconv.Itoa(int(service.Login.ID))

		roles := service.Canbin.GetRolesForUser(role)

		ok := service.Canbin.CheckRoles(roles, ctx.Request.URL.Path, ctx.Request.Method)

		if ok {
			ctx.Next()
		} else {
			app.JsonResponse(ctx, errno.ErrPermissionDenied, nil)
			ctx.Abort()
		}
		return
	}
}
