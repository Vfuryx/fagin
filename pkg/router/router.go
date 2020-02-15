package router

import (
	"fagin/app"
	"fagin/app/middleware"
	"fagin/config"
	"fagin/pkg/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

// 路由引擎
var Engine *gin.Engine = New()

func New() *gin.Engine {
	e := gin.New()

	// 是否正式环境
	if ok := app.IsProd(); !ok {
		e.Use(gin.Logger(), gin.Recovery())
	}

	// 全局实例Session
	e.Use(session.Sessions())
	// 支持跨域
	if config.DefaultRouter.IsCors {
		conf := cors.DefaultConfig()
		conf.AllowOrigins = config.DefaultRouter.CorsConf.AllowOrigins
		conf.AllowHeaders = config.DefaultRouter.CorsConf.AllowHeaders
		conf.AllowMethods = config.DefaultRouter.CorsConf.AllowMethods
		e.Use(cors.New(conf))
	}
	// 收集日志
	e.Use(middleware.RecoveryLog())
	// 设置静态资源
	e.Static(config.Template.StaticRouter, config.Template.Static)
	// 加载顺序 1 - 3 顺序不能变
	// 1 定界符
	e.Delims(config.Template.DelimitersL, config.Template.DelimitersR)
	// 2 注册模版函数
	e.SetFuncMap(config.Template.FuncMap)
	// 3 设置模版目录
	e.LoadHTMLGlob(config.Template.LoadHTMLGlob)
	// 限定表单占用内存
	//engine.MaxMultipartMemory = 32 << 20

	// 配置 404 模块
	e.NoRoute(noRoute)

	return e
}

func Group(name string) *gin.RouterGroup {
	return Engine.Group(name)
}

// 存入 NoRouteMap
func NoRoute(basePath string, handle gin.HandlerFunc) {
	err := rootNode.SetPathAndFunc(basePath, handle)
	if err != nil {
		panic(err)
	}
}

// 404处理
func noRoute(ctx *gin.Context) {
	f := rootNode.GetPathFunc(ctx.Request.RequestURI)
	if f != nil {
		f(ctx)
	}
}
