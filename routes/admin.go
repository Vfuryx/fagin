package routes

import (
	"fagin/app/controllers/admin"
	"fagin/app/middleware"
	"fagin/config"
	"fagin/pkg/router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Admin = router.Group("admin")

func init() {
	// 记录操作日志中间件
	Admin.Use(middleware.AdminOperationLog.LoggerToDB())

	router.NoRoute(Admin.BasePath(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"m": "admin 404"})
	})

	// 配置跨域
	if config.AdminRouter.IsCors {
		conf := cors.DefaultConfig()
		conf.AllowOrigins = config.AdminRouter.CorsConf.AllowOrigins
		conf.AllowHeaders = config.AdminRouter.CorsConf.AllowHeaders
		conf.AllowMethods = config.AdminRouter.CorsConf.AllowMethods
		Admin.Use(cors.New(conf))
	}

	// 后台页面
	Admin.GET("/", admin.IndexController.Index)
	// 后台api
	api := Admin.Group("/api")
	{
		// 登录
		api.POST("/login", admin.AuthController.Login)
		// 获取验证码
		api.POST("/captcha", admin.AuthController.GetCaptcha)

		// 播放视频
		api.GET("/play/av:id", admin.VideoController.PlayVideo)

		// v1 													// 权限验证
		apiV1 := api.Group("/v1", middleware.AdminAuth.IsLogin(), middleware.AdminAuth.AuthCheckRole())
		{
			// 获取用户信息
			apiV1.GET("/us/info", admin.AuthController.Info)
			// 获取角色菜单
			apiV1.GET("/us/routes", admin.AuthController.Routes)
			// 修改用户信息
			//apiV1.PUT("/user/:id", admin.AuthController.UpdateAdminUser)
			// 登出
			apiV1.POST("/user/logout", admin.AuthController.Logout)

			// 用户管理
			user := apiV1.Group("/user")
			{
				// 用户列表
				user.GET("/", admin.UserController.Index)
				// 展示用户
				user.GET("/:id", admin.UserController.Show)
				// 新增用户
				user.POST("/", admin.UserController.Store)
				// 更新用户
				user.PUT("/:id", admin.UserController.Update)
				// 用户更新状态
				user.PUT("/:id/status/", admin.UserController.UpdateStatus)
				// 用户更新密码
				user.PUT("/:id/reset/", admin.UserController.Reset)
				// 删除用户
				user.DELETE("/:id", admin.UserController.Del)
				// 批量删除用户
				user.DELETE("/", admin.UserController.Dels)
			}

			// 菜单管理
			menu := apiV1.Group("/menu")
			{
				// 菜单列表
				menu.GET("/", admin.MenuController.Index)
				// 菜单展示
				menu.GET("/:id", admin.MenuController.Show)
				// 菜单新增
				menu.POST("/", admin.MenuController.Store)
				// 菜单更新
				menu.PUT("/:id", admin.MenuController.Update)
				// 菜单删除
				menu.DELETE("/:id", admin.MenuController.Del)
			}

			// 角色管理
			role := apiV1.Group("/role")
			{
				// 角色列表
				role.GET("/", admin.RoleController.Index)
				// 角色列表
				role.POST("/list", admin.RoleController.Roles)
				// 角色展示
				role.GET("/:id", admin.RoleController.Show)
				// 角色新增
				role.POST("/", admin.RoleController.Store)
				// 角色更新
				role.PUT("/:id", admin.RoleController.Update)
				// 角色更新状态
				role.PUT("/:id/status/", admin.RoleController.UpdateStatus)
				// 角色删除
				role.DELETE("/:id", admin.RoleController.Del)
				// 角色组删除
				role.DELETE("/", admin.RoleController.Dels)

			}

			// 操作日志管理
			logs := apiV1.Group("/operation/logs")
			{
				// 操作日志列表
				logs.GET("/", admin.OperationLogController.Index)
				// 操作日志详情
				logs.GET("/:id", admin.OperationLogController.Show)
			}

			// 轮播图管理
			banner := apiV1.Group("/banner")
			{
				// 轮播图列表
				banner.GET("/", admin.BannerController.Index)
				// 创建轮播图
				banner.POST("/", admin.BannerController.Store)
				// 轮播图详情
				banner.GET("/:id", admin.BannerController.Show)
				// 更新轮播图
				banner.PUT("/:id", admin.BannerController.Update)
				// 上传轮播图
				banner.POST("/upload", admin.BannerController.Upload)
				// 轮播图删除
				banner.DELETE("/:id", admin.BannerController.Del)
				// 批量删除轮播图
				banner.DELETE("/", admin.BannerController.DeleteBanners)
			}

			// 视频管理
			video := apiV1.Group("/video")
			{
				// 视频列表
				video.GET("/list", admin.VideoController.VideoList)
				// 创建视频
				video.POST("/", admin.VideoController.CreateVideo)
				// 更新视频
				video.PUT("/:id", admin.VideoController.UpdateVideo)
				// 删除视频
				video.DELETE("/:id", admin.VideoController.DeleteVideo)
				// 批量删除视频
				video.DELETE("/", admin.VideoController.DeleteVideos)
				// 上传视频
				video.POST("/upload", admin.VideoController.UploadVideo)
			}

			// 网站管理
			website := apiV1.Group("/website")
			{
				// 查看网站设置
				website.GET("/info", admin.WebsiteConfigController.Info)
				// 更新网站设置
				website.PUT("/info", admin.WebsiteConfigController.UpdateInfo)
				// 上传图片
				website.POST("/upload", admin.WebsiteConfigController.Upload)
			}

			// 关于我们
			company := apiV1.Group("/company")
			{
				// 查看公司介绍
				company.GET("/introduction", admin.CompanyIntroductionController.ShowCompanyIntroduction)
				// 更新公司介绍
				company.PUT("/introduction", admin.CompanyIntroductionController.UpdateCompanyIntroduction)
				// 上传图片
				company.POST("/upload", admin.CompanyIntroductionController.Upload)
			}
		}
	}
}
