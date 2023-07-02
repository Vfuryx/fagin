package router

import (
	"fadmin/app"
	"fadmin/config"
	"fadmin/pkg/router/no_router"
	"fadmin/routes"
	"github.com/bytedance/sonic"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/pprof"
	"github.com/gofiber/fiber/v2/middleware/recover"
	"github.com/pkg/errors"
	"html/template"
	"net/http"
	"strings"
)

// Init 初始化化
func Init() *fiber.App {
	App := fiber.New(fiber.Config{
		Prefork:           false,
		CaseSensitive:     true,
		EnablePrintRoutes: true,

		JSONEncoder: sonic.Marshal,
		JSONDecoder: sonic.Unmarshal,

		ErrorHandler: DefaultErrorHandler,

		ServerHeader: "Furyx",
		AppName:      "Furyx App v1.0.0",
	})
	//// 加载模版
	//temp := loadHTMLGlobFS(App)

	// 不是是正式环境
	if ok := app.IsProd(); !ok {
		//view.DebugPrintLoadTemplate(temp)
		//App.Use(gin.Logger())
	} else {
		//gin.SetMode(gin.ReleaseMode)
	}

	// 收集日志并恢复
	App.Use(recover.New())
	//App.Use(recover.New(), middleware.RecoveryLog())

	// 加载路由
	routes.Handle(App)

	// 设置公开静态资源 （上传公开文件）
	App.Static(config.Template().PublicRouter(), config.Template().Public())

	// 设置固定静态资源文件
	App.Use(config.Template().StaticRouter(), filesystem.New(filesystem.Config{
		Root: http.FS(config.Template().StaticEmbed()),
	}))

	// 全局实例Session
	// App.Use(session.Sessions())

	// 支持跨域
	if config.DefaultRouter().IsCors() {
		conf := cors.Config{}
		conf.AllowOrigins = strings.Join(config.DefaultRouter().CorsConf().AllowOrigins, ",")
		conf.AllowHeaders = strings.Join(config.DefaultRouter().CorsConf().AllowHeaders, ",")
		conf.AllowMethods = strings.Join(config.DefaultRouter().CorsConf().AllowMethods, ",")
		App.Use(cors.New(conf))
	}

	// 限定表单占用内存
	// App.MaxMultipartMemory = 32 << 20

	// 设置pprof
	setPprof(App)

	// 404
	no_router.NoRoute("/", func(ctx *fiber.Ctx) error {
		return ctx.JSON(fiber.Map{"m": "/ 404"})
	})

	return App
}

var DefaultErrorHandler = func(ctx *fiber.Ctx, err error) error {
	// Status code defaults to 500
	code := fiber.StatusInternalServerError

	// Retrieve the custom status code if it's a *fiber.Error
	var e *fiber.Error
	if errors.As(err, &e) {
		code = e.Code
	}

	if code == fiber.StatusNotFound {
		// 配置 404 模块
		return no_router.NoRouteHandle(ctx)
	} else {
		// Set Content-Type: text/plain; charset=utf-8
		ctx.Set(fiber.HeaderContentType, fiber.MIMETextPlainCharsetUTF8)
		// Return status code with error message
		return ctx.Status(code).SendString(err.Error())
	}
}

// 设置 Pprof
func setPprof(e *fiber.App) {
	if app.IsProd() {
		e.Use(pprof.New(pprof.Config{Prefix: "/debug"}))
	}
}

// 加载模版
func loadHTMLGlobFS(e *fiber.App) *template.Template {
	temp := template.Must(template.New("").
		Delims(config.Template().DelimitersL(), config.Template().DelimitersR()). // 1 设置定界符
		Funcs(config.Template().FuncMap()). // 2 注册模版函数
		ParseFS(config.Template().TemplatesEmbed(), config.Template().TemplatePattern())) // 3 导入模版

	//e.SetHTMLTemplate(temp)

	return temp
}
