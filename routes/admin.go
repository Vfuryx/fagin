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
	var log = middleware.AdminOperationLog.Log
	// 初始化操作日志中间件
	Admin.Use(middleware.AdminOperationLog.Operation())

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
		api.POST("/login", admin.AuthController.Login, log("后台登录"))
		// 播放视频
		api.GET("/play/av:id", admin.VideoController.PlayVideo)

		// v1 													// 权限验证
		apiV1 := api.Group("/v1", middleware.AdminAuth.IsLogin())
		{
			// 获取用户信息
			apiV1.GET("/user/info", admin.AuthController.Info)
			// 修改用户信息
			apiV1.PUT("/user/:id", admin.AuthController.UpdateAdminUser, log("修改用户信息"))
			// 登出
			apiV1.POST("/user/logout", admin.AuthController.Logout)

			// 操作日志列表
			logs := apiV1.Group("/operation/logs")
			{
				// 查看操作日志列表
				logs.GET("/", admin.OperationLogController.Index)
				// 查看操作日志
				logs.GET("/:id", admin.OperationLogController.Show)
			}

			// 轮播图管理
			banner := apiV1.Group("/banner")
			{
				// 轮播图列表
				banner.GET("/", admin.BannerController.Index)
				// 创建轮播图
				banner.POST("/", admin.BannerController.Store, log("创建轮播图"))
				// 轮播图详情
				banner.GET("/:id", admin.BannerController.Show)
				// 更新轮播图
				banner.PUT("/:id", admin.BannerController.Update, log("更新轮播图"))
				// 上传轮播图
				banner.POST("/upload", admin.BannerController.Upload)
			}

			// 视频管理
			video := apiV1.Group("/video")
			{
				// 视频列表
				video.GET("/list", admin.VideoController.VideoList)
				// 创建视频
				video.POST("/", admin.VideoController.CreateVideo, log("创建视频"))
				// 更新视频
				video.PUT("/:id", admin.VideoController.UpdateVideo, log("更新视频"))
				// 删除视频
				video.DELETE("/:id", admin.VideoController.DeleteVideo, log("删除视频"))
				// 上传视频
				video.POST("/upload", admin.VideoController.UploadVideo)
			}

			// 网站管理
			website := apiV1.Group("/website")
			{
				// 查看网站设置
				website.GET("/info", admin.WebsiteConfigController.Info)
				// 更新网站设置
				website.PUT("/info", admin.WebsiteConfigController.UpdateInfo, log("更新网站设置"))
				// 上传图片
				website.POST("/upload", admin.WebsiteConfigController.Upload)
			}

			// 关于我们
			company := apiV1.Group("/company")
			{
				// 查看公司介绍
				company.GET("/introduction", admin.CompanyIntroductionController.ShowCompanyIntroduction)
				// 更新公司介绍
				company.PUT("/introduction", admin.CompanyIntroductionController.UpdateCompanyIntroduction, log("更新公司介绍"))
				// 上传图片
				company.POST("/upload", admin.CompanyIntroductionController.Upload)
			}
		}
	}
}
