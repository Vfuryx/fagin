package routes

import (
	"fagin/app/controllers/admin"
	"fagin/app/middleware"
	"fagin/config"
	"fagin/pkg/router/no_router"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

var adminRoute routeFunc = func(Admin *gin.RouterGroup) {
	// 404
	no_router.NoRoute(Admin.BasePath(), func(ctx *gin.Context) {
		ctx.JSON(http.StatusNotFound, gin.H{"m": "admin 404"})
	})

	// 配置跨域
	if config.AdminRouter().IsCors() {
		conf := cors.DefaultConfig()
		conf.AllowOrigins = config.AdminRouter().CorsConf().AllowOrigins
		conf.AllowHeaders = config.AdminRouter().CorsConf().AllowHeaders
		conf.AllowMethods = config.AdminRouter().CorsConf().AllowMethods
		Admin.Use(cors.New(conf))
	}

	// 后台页面
	Admin.GET("/", admin.IndexController.Index)
	// 后台api
	api := Admin.Group("/api")
	{
		api.POST("/login", admin.AuthController.Login)          // 登录
		api.POST("/captcha", admin.AuthController.GetCaptcha)   // 获取验证码
		api.GET("/play/av:id", admin.VideoController.PlayVideo) // 播放视频

		// v1
		apiAuth := api.Group("",
			middleware.AdminAuth.IsLogin(),
			middleware.AdminAuth.CheckRoleCache(), // 权限验证
			middleware.AdminOperationLogToDB(),    // 记录操作日志中间件
		)
		{
			// 公共模块
			common := apiAuth.Group("common")
			{
				common.GET("/auth/info", admin.AuthController.Info)               // 获取用户信息
				common.GET("/auth/permcode", admin.AuthController.PermissionCode) // 获取用户权限
				common.GET("/auth/routes", admin.AuthController.Routes)           // 获取菜单（新）
				common.POST("/auth/logout", admin.AuthController.Logout)          // 登出
				// apiAuth.PUT("/user/:id", admin.AuthController.UpdateAdminUser) 			// 修改用户信息
			}

			// 用户管理
			user := apiAuth.Group("/users")
			{
				user.GET("", admin.UserController.Index)                   // 用户查询
				user.GET("/:id", admin.UserController.Show)                // 用户详情
				user.POST("", admin.UserController.Store)                  // 新增用户
				user.PUT("/:id", admin.UserController.Update)              // 更新用户
				user.PUT("/:id/status", admin.UserController.UpdateStatus) // 用户更新状态
				user.PUT("/:id/pwd", admin.UserController.Reset)           // 用户更新密码
				user.DELETE("/:id", admin.UserController.Del)              // 删除用户
				user.DELETE("", admin.UserController.Dels)                 // 批量删除用户
			}

			// 部门管理
			departments := apiAuth.Group("/departments")
			{
				departments.GET("", admin.DepartmentController.Index)      // 部门列表
				departments.GET("/:id", admin.DepartmentController.Show)   // 部门详情
				departments.POST("", admin.DepartmentController.Store)     // 新增部门
				departments.PUT("/:id", admin.DepartmentController.Update) // 更新部门
				departments.DELETE("/:id", admin.DepartmentController.Del) // 删除部门
			}

			// 账号管理
			admins := apiAuth.Group("/accounts")
			{
				admins.GET("", admin.AdminsController.Index)                   // 账号查询
				admins.GET("/:id", admin.AdminsController.Show)                // 账号详情
				admins.POST("", admin.AdminsController.Store)                  // 新增账号
				admins.PUT("/:id", admin.AdminsController.Update)              // 更新账号
				admins.DELETE("/:id", admin.AdminsController.Delete)           // 删除账号
				admins.PUT("/:id/status", admin.AdminsController.UpdateStatus) // 更新状态
				admins.PUT("/:id/pwd", admin.AdminsController.ResetPassword)   // 重置密码
				admins.POST("/:id/logout", admin.AdminsController.Logout)      // 强制下线
				admins.GET("/exists", admin.AdminsController.UsernameExists)   // 账号名是否已存在
				admins.GET("/roles", admin.RoleController.Roles)               // 角色列表
				admins.GET("/departments", admin.DepartmentController.Index)   // 部门列表
				admins.GET("/role-route", admin.AdminsController.RolesRoute)   // 角色列表
			}

			// 菜单管理
			menu := apiAuth.Group("/menus")
			{
				menu.GET("", admin.MenuController.Index)         // 菜单查询
				menu.GET("/:id", admin.MenuController.Show)      // 菜单详情
				menu.POST("", admin.MenuController.Store)        // 新增菜单
				menu.PUT("/:id", admin.MenuController.Update)    // 更新菜单
				menu.DELETE("/:id", admin.MenuController.Delete) // 删除菜单
				menu.GET("/all", admin.MenuController.All)       // 所有菜单
			}

			// 角色管理
			role := apiAuth.Group("/roles")
			{
				role.GET("", admin.RoleController.Index)                   // 角色查询
				role.GET("/:id", admin.RoleController.Show)                // 角色详情
				role.POST("", admin.RoleController.Store)                  // 角色新增
				role.PUT("/:id", admin.RoleController.Update)              // 角色更新
				role.DELETE("/:id", admin.RoleController.Delete)           // 角色删除
				role.GET("/key", admin.RoleController.KeyExist)            // 角色 key 是否存在
				role.PUT("/:id/status", admin.RoleController.UpdateStatus) // 角色更新状态
				role.GET("/menus", admin.MenuController.All)               // 所有菜单

				// role.GET("/all", admin.RoleController.Roles)               		  // 所有角色
			}

			// 操作日志管理
			logs := apiAuth.Group("/operation/logs")
			{
				logs.GET("", admin.OperationLogController.Index)    // 操作日志查询
				logs.GET("/:id", admin.OperationLogController.Show) // 操作日志详情
			}

			// 登录日志管理
			loginLog := apiAuth.Group("/login-logs")
			{
				loginLog.GET("", admin.LoginLogController.Index)    // 登录日志查询
				loginLog.GET("/:id", admin.LoginLogController.Show) // 登录日志详情
			}

			// 轮播图管理
			banner := apiAuth.Group("/banner")
			{
				banner.GET("/", admin.BannerController.Index)            // 轮播图列表
				banner.POST("/", admin.BannerController.Store)           // 创建轮播图
				banner.GET("/:id", admin.BannerController.Show)          // 轮播图详情
				banner.PUT("/:id", admin.BannerController.Update)        // 更新轮播图
				banner.POST("/upload", admin.BannerController.Upload)    // 上传轮播图
				banner.DELETE("/:id", admin.BannerController.Del)        // 轮播图删除
				banner.DELETE("/", admin.BannerController.DeleteBanners) // 批量删除轮播图
			}

			// 视频管理
			video := apiAuth.Group("/videos")
			{
				video.GET("/", admin.VideoController.VideoList)          // 视频列表
				video.POST("/", admin.VideoController.CreateVideo)       // 创建视频
				video.PUT("/:id", admin.VideoController.UpdateVideo)     // 更新视频
				video.DELETE("/:id", admin.VideoController.DeleteVideo)  // 删除视频
				video.DELETE("/", admin.VideoController.DeleteVideos)    // 批量删除视频
				video.POST("/upload", admin.VideoController.UploadVideo) // 上传视频
			}

			// 网站管理
			website := apiAuth.Group("/website")
			{
				website.GET("/info", admin.WebsiteConfigController.Info)       // 查看网站设置
				website.PUT("/info", admin.WebsiteConfigController.UpdateInfo) // 更新网站设置
				website.POST("/upload", admin.WebsiteConfigController.Upload)  // 上传图片
			}

			// 关于我们
			company := apiAuth.Group("/company")
			{
				company.GET("/introduction", admin.CompanyIntroductionController.ShowCompanyIntroduction)   // 查看公司介绍
				company.PUT("/introduction", admin.CompanyIntroductionController.UpdateCompanyIntroduction) // 更新公司介绍
				company.POST("/upload", admin.CompanyIntroductionController.Upload)                         // 上传图片
			}

			// 文章管理
			articles := apiAuth.Group("/articles")
			{
				articles.GET("/", admin.ArticleController.Index)      // 列表
				articles.GET("/:id", admin.ArticleController.Show)    // 查看
				articles.POST("/", admin.ArticleController.Store)     // 新建
				articles.PUT("/:id", admin.ArticleController.Update)  // 更新
				articles.DELETE("/:id", admin.ArticleController.Del)  // 删除
				articles.DELETE("/", admin.ArticleController.Deletes) // 批量删除
			}

			// 分类管理
			categories := apiAuth.Group("/categories")
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
			authors := apiAuth.Group("/authors")
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
			tags := apiAuth.Group("/tags")
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
