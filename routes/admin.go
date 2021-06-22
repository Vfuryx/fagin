package routes

import (
	"fagin/app/controllers/admin"
	"fagin/app/middleware"
	"fagin/config"
	"fagin/pkg/router/no_router"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"net/http"
)

var adminRoute = func(Admin *gin.RouterGroup) {
	// 记录操作日志中间件
	Admin.Use(middleware.AdminOperationLogToDB())

	no_router.NoRoute(Admin.BasePath(), func(ctx *gin.Context) {
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
		apiV1 := api.Group(
			"/v1",
			middleware.AdminAuth.IsLogin(),
			middleware.AdminAuth.AuthCheckRoleCache(),
		)
		{
			// 获取用户信息
			apiV1.GET("/us/info", admin.AuthController.Info)
			//// 获取角色菜单
			//apiV1.GET("/us/routes", admin.AuthController.Routes)
			// 获取菜单（新）
			apiV1.GET("/us/navs", admin.AuthController.Navs)
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

			// 管理员管理
			admins := apiV1.Group("/admins")
			{
				admins.GET("/", admin.AdminsController.Index)                   // 用户列表
				admins.GET("/:id", admin.AdminsController.Show)                 // 用户详情
				admins.POST("/", admin.AdminsController.Store)                  // 新增用户
				admins.PUT("/:id", admin.AdminsController.Update)               // 更新用户
				admins.PUT("/:id/status/", admin.AdminsController.UpdateStatus) // 更新状态
				admins.PUT("/:id/reset/", admin.AdminsController.ResetPassword) // 重置密码
				admins.DELETE("/:id", admin.AdminsController.Delete)            // 删除用户
				admins.POST("/logout/:id", admin.AdminsController.Logout)       // 强制下线
				admins.GET("/username/", admin.AdminsController.UsernameExist)  // 用户名是否已存在
			}

			// 菜单管理
			menu := apiV1.Group("/menu")
			{
				menu.GET("/all", admin.MenuController.All)       // 所有菜单
				menu.GET("/", admin.MenuController.Index)        // 菜单列表
				menu.GET("/:id", admin.MenuController.Show)      // 菜单展示
				menu.POST("/", admin.MenuController.Store)       // 菜单新增
				menu.PUT("/:id", admin.MenuController.Update)    // 菜单更新
				menu.DELETE("/:id", admin.MenuController.Delete) // 菜单删除
			}

			apiV1.GET("/group_permissions", admin.PermissionController.GroupPermissions)

			// 权限组列表
			groups := apiV1.Group("/permission/groups")
			{
				groups.GET("/all", admin.PermissionController.GAll)
				groups.GET("/", admin.PermissionController.GIndex)
				groups.GET("/:id", admin.PermissionController.GShow)
				groups.POST("/", admin.PermissionController.GStore)
				groups.PUT("/:id", admin.PermissionController.GUpdate)
			}
			// 权限管理
			permissions := apiV1.Group("/permissions")
			{
				permissions.GET("/", admin.PermissionController.Index)        // 列表
				permissions.GET("/:id", admin.PermissionController.Show)      // 展示
				permissions.POST("/", admin.PermissionController.Store)       // 新增
				permissions.PUT("/:id", admin.PermissionController.Update)    // 更新
				permissions.DELETE("/:id", admin.PermissionController.Delete) // 删除
			}

			// 角色管理
			role := apiV1.Group("/role")
			{
				role.GET("/key", admin.RoleController.KeyExist)             // 角色 key 是否存在
				role.GET("/all", admin.RoleController.Roles)                // 角色列表
				role.GET("/", admin.RoleController.Index)                   // 角色列表
				role.GET("/:id", admin.RoleController.Show)                 // 角色展示
				role.POST("/", admin.RoleController.Store)                  // 角色新增
				role.PUT("/:id", admin.RoleController.Update)               // 角色更新
				role.PUT("/:id/status/", admin.RoleController.UpdateStatus) // 角色更新状态
				role.DELETE("/:id", admin.RoleController.Delete)            // 角色删除
			}

			// 操作日志管理
			logs := apiV1.Group("/operation/logs")
			{
				logs.GET("/", admin.OperationLogController.Index)   // 操作日志列表
				logs.GET("/:id", admin.OperationLogController.Show) // 操作日志详情
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
			video := apiV1.Group("/videos")
			{
				video.GET("/", admin.VideoController.VideoList)          // 视频列表
				video.POST("/", admin.VideoController.CreateVideo)       // 创建视频
				video.PUT("/:id", admin.VideoController.UpdateVideo)     // 更新视频
				video.DELETE("/:id", admin.VideoController.DeleteVideo)  // 删除视频
				video.DELETE("/", admin.VideoController.DeleteVideos)    // 批量删除视频
				video.POST("/upload", admin.VideoController.UploadVideo) // 上传视频
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

			// 文章管理
			articles := apiV1.Group("/articles")
			{
				articles.GET("/", admin.ArticleController.Index)      // 列表
				articles.GET("/:id", admin.ArticleController.Show)    // 查看
				articles.POST("/", admin.ArticleController.Store)     // 新建
				articles.PUT("/:id", admin.ArticleController.Update)  // 更新
				articles.DELETE("/:id", admin.ArticleController.Del)  // 删除
				articles.DELETE("/", admin.ArticleController.Deletes) // 批量删除
			}

			// 分类管理
			categories := apiV1.Group("/categories")
			{
				categories.GET("/", admin.CategoryController.Index)      // 列表
				categories.GET("/:id", admin.CategoryController.Show)    // 查看
				categories.POST("/", admin.CategoryController.Store)     // 新建
				categories.POST("/all", admin.CategoryController.All)    // 所有
				categories.PUT("/:id", admin.CategoryController.Update)  // 更新
				categories.DELETE("/:id", admin.CategoryController.Del)  // 删除
				categories.DELETE("/", admin.CategoryController.Deletes) // 批量删除
			}

			// 作者管理
			authors := apiV1.Group("/authors")
			{
				authors.GET("/", admin.AuthorController.Index)      // 列表
				authors.GET("/:id", admin.AuthorController.Show)    // 查看
				authors.POST("/", admin.AuthorController.Store)     // 新建
				authors.POST("/all", admin.AuthorController.All)    // 所有
				authors.PUT("/:id", admin.AuthorController.Update)  // 更新
				authors.DELETE("/:id", admin.AuthorController.Del)  // 删除
				authors.DELETE("/", admin.AuthorController.Deletes) // 批量删除
			}

			// 标签管理
			tags := apiV1.Group("/tags")
			{
				tags.GET("/", admin.TagController.Index)      // 列表
				tags.GET("/:id", admin.TagController.Show)    // 查看
				tags.POST("/", admin.TagController.Store)     // 新建
				tags.POST("/all", admin.TagController.All)    // 所有
				tags.PUT("/:id", admin.TagController.Update)  // 更新
				tags.DELETE("/:id", admin.TagController.Del)  // 删除
				tags.DELETE("/", admin.TagController.Deletes) // 批量删除
			}

		}
	}
}
