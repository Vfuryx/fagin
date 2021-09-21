package router

import (
	"fagin/app"
	"fagin/app/middleware"
	"fagin/config"
	"fagin/pkg/router/no_router"
	"fagin/pkg/view"
	"fagin/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"html/template"
	"net/http"
)

func New() *gin.Engine {
	e := gin.New()

	// 加载模版
	temp := loadHTMLGlobFS(e)

	// 是否正式环境
	if ok := app.IsProd(); !ok {
		view.DebugPrintLoadTemplate(temp)
		e.Use(gin.Logger(), gin.Recovery())
	}

	// 收集日志
	e.Use(middleware.RecoveryLog())

	// 加载路由
	routes.Handle(e)

	// 设置公开静态资源 （上传公开文件）
	e.Static(config.Template.PublicRouter, config.Template.Public)

	// 设置固定静态资源文件
	e.StaticFS(config.Template.StaticRouter, http.FS(config.Template.StaticEmbed))

	// 全局实例Session
	//e.Use(session.Sessions())

	// 支持跨域
	if config.DefaultRouter.IsCors {
		conf := cors.DefaultConfig()
		conf.AllowOrigins = config.DefaultRouter.CorsConf.AllowOrigins
		conf.AllowHeaders = config.DefaultRouter.CorsConf.AllowHeaders
		conf.AllowMethods = config.DefaultRouter.CorsConf.AllowMethods
		e.Use(cors.New(conf))
	}

	// 限定表单占用内存
	//e.MaxMultipartMemory = 32 << 20

	// 配置 404 模块
	e.NoRoute(no_router.NoRouteHandle)

	// 设置pprof
	setPprof(e)

	return e
}

// 设置 Pprof
func setPprof(e *gin.Engine) {
	adminGroup := e.Group("/debug", func(c *gin.Context) {
		if app.IsProd() {
			if c.Query("debug") != "debug" {
				c.AbortWithStatus(http.StatusForbidden)
				return
			}
		}
		c.Next()
	})
	pprof.RouteRegister(adminGroup, "pprof")
}

// 加载模版
func loadHTMLGlobFS(e *gin.Engine) *template.Template {
	temp := template.Must(template.New("").
		Delims(config.Template.DelimitersL, config.Template.DelimitersR).         // 1 设置定界符
		Funcs(config.Template.FuncMap).                                           // 2 注册模版函数
		ParseFS(config.Template.TemplatesEmbed, config.Template.TemplatePattern)) // 3 导入模版

	e.SetHTMLTemplate(temp)
	return temp
}
