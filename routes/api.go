package routes

import (
	"fagin/app"
	"fagin/app/controllers/api"
	"fagin/app/errno"
	"fagin/app/middleware" // 中间件
	"fagin/pkg/router/no_router"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

var apiRoute = func(Api *gin.RouterGroup) {
	// 404
	no_router.NoRoute(Api.BasePath(), func(ctx *gin.Context) {
		app.ResponseJsonWithStatus(ctx, http.StatusNotFound, errno.NotFoundErr, nil, gin.H{"m": "api 404"})
	})

	// 测试限流中间件
	Api.Use(middleware.RateLimitMiddleware(1000, 60*time.Second, func(ctx *gin.Context) {
		ctx.String(http.StatusServiceUnavailable, "服务器繁忙")
	}))

	Api.GET("/", api.IndexController.Index)

	V1 := Api.Group("/v1")
	{
		V1.POST("/login", api.AuthController.Login)       // 用户登陆 获取token
		V1.POST("/register", api.UserController.Register) // 注册用户

		// 视频
		video := Api.Group("/video")
		{
			video.GET("/play/:id", api.VideoController.PlayVideo) // 观看视频
		}

		// 需要登陆才能调用
		V1.Use(middleware.Auth.IsLogin(), middleware.Auth.AuthCheckRole())
		{
			V1.GET("/", api.IndexController.Index)
			V1.POST("/", api.IndexController.Index)
			V1.DELETE("/", api.IndexController.Index)
			V1.GET("/logout", api.AuthController.Logout)
			V1.GET("/add", api.AuthController.CreateUser)

			//用户组
			user := V1.Group("/user")
			{
				user.GET("/", api.UserController.UserList) // 用户列表
				id := user.Group("/:id")
				{
					id.GET("/", api.UserController.ShowUser)       // 查看用户
					id.PATCH("/", api.UserController.UpdateUser)   // 更新用户
					id.DELETE("/", api.UserController.DestroyUser) // 删除用户
				}
			}
		}

	}
}
