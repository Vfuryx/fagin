package router

import (
	"fagin/app"
	"fagin/app/middleware"
	"fagin/config"
	"fagin/pkg/session"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"strings"
)

// 404 树 （有个坑，不能超过 32 个元素）
// 下标	- 	字符
// 0 	| 	["":{}, "abc"{}]
// 1 	|	["123":{}, "456":{}]
// /abc/123 能匹配到 0 => "abc":{} , 2 => "123":{} 	最后得到 abc/123
// /abc/789 能匹配到 0 => "abc":{} 					最后得到 abc
type noRouteTreeNode map[string]interface{}
type noRouteTree map[string]noRouteTreeNode

var noRouteTreeMap = make([]noRouteTree, 32)

// 404 映射 方法
var noRouteMap = map[string]gin.HandlerFunc{}

// 路由引擎
var engine *gin.Engine

func init() {
	// 实例化 noRouteTreeMap
	for index, _ := range noRouteTreeMap {
		noRouteTreeMap[index] = noRouteTree{}
	}

	// 是否正式环境
	if ok := app.IsProd(); ok {
		engine = gin.New()
	} else {
		engine = gin.Default()
	}

	// 全局实例Session
	engine.Use(session.Sessions())
	// 支持跨域
	if config.DefaultRouter.IsCors {
		conf := cors.DefaultConfig()
		conf.AllowOrigins = config.DefaultRouter.CorsConf.AllowOrigins
		conf.AllowHeaders = config.DefaultRouter.CorsConf.AllowHeaders
		conf.AllowMethods = config.DefaultRouter.CorsConf.AllowMethods
		engine.Use(cors.New(conf))
	}
	// 收集日志
	engine.Use(middleware.RecoveryLog())
	// 加载顺序 1 - 3 顺序不能变
	// 1 定界符
	engine.Delims(config.Template.DelimitersL, config.Template.DelimitersR)
	// 2 注册模版函数
	engine.SetFuncMap(config.Template.FuncMap)
	// 3 设置模版目录
	engine.LoadHTMLGlob(config.Template.LoadHTMLGlob)
	// 设置静态资源
	engine.Static(config.Template.StaticRouter, config.Template.Static)
	// 限定表单占用内存
	//engine.MaxMultipartMemory = 32 << 20
}

func New() *gin.Engine {
	// 配置 404 模块
	engine.NoRoute(noRoute)

	return engine
}

func Group(name string) *gin.RouterGroup {
	return engine.Group(name)
}

// 存入 NoRouteMap
func NoRoute(basePath string, handle gin.HandlerFunc) {
	// 过滤前后的/
	basePath = strings.Trim(basePath, "/")
	// 分割uri
	s := strings.Split(basePath, "/")

	i := 0
	var prev string
	for _, value := range s {
		// 去除零值
		if value == "" {
			continue
		}

		// 判断处已经否存在这个记录
		if _, ok := noRouteTreeMap[i][value]; !ok {
			noRouteTreeMap[i][value] = noRouteTreeNode{}
		}

		if i > 0 {
			noRouteTreeMap[i-1][prev][value] = noRouteTreeMap[i][value]
		}

		prev = value
		i++
	}

	noRouteMap[basePath] = handle
}

// 404处理
func noRoute(ctx *gin.Context) {
	// 分割uri
	s := strings.Split(ctx.Request.RequestURI, "/")
	i := 0
	source := make([]string, 0, 20)
	var next = noRouteTreeNode{}
	var ok bool
	for _, value := range s {
		// 去除零值
		if value == "" {
			continue
		}

		if len(next) <= 0 {
			next, ok = noRouteTreeMap[i][value]
			if !ok {
				break
			}

			source = append(source, value)
			continue
		}

		v, ok := next[value]
		if !ok {
			break
		}

		source = append(source, value)
		i++

		if next, ok = v.(noRouteTreeNode); !ok {
			break
		}

	}

	path := ""
	var f func(ctx *gin.Context)
	i = 0
	// /a/b/c 	=> func
	// /a		=> func
	// >>> /a/b ---> /a/b
	// 降级处理
	// >>> /a/b ----> /a
	for {
		if len(source) < i {
			return
		}
		path = strings.Join(source[:len(source)-i], "/")
		f, ok = noRouteMap[path]
		if ok {
			// 自定义方法
			f(ctx)
			return
		}
		i++
	}
}
