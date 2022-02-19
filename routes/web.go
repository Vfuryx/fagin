package routes

import (
	"fagin/app/controllers/web"
	"fagin/app/middleware"
	"fagin/pkg/router/no_router"
	"net/http"

	"github.com/gin-gonic/gin"
)

var webRoute routeFunc = func(Web *gin.RouterGroup) {
	// 404
	no_router.NoRoute(Web.BasePath(), func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "404 NotFoundErr")
	})

	// 限流中间件
	// limiter := middleware.LimitMiddleware(2, func(ctx *gin.Context) {
	// 	ctx.String(http.StatusBadGateway, "服务器繁忙")
	// })

	w := Web.Group("/", middleware.WebNavbar.WebNavbar())
	{
		// 首页
		w.GET("/", middleware.WebTags.WebTags(), web.IndexController.Home)
		// 类型
		w.GET("/category/:cate", middleware.WebTags.WebTags(), web.IndexController.Category)
		// 标签
		w.GET("/tag/:tag", middleware.WebTags.WebTags(), web.IndexController.Tag)
		// 文章
		w.GET("/article/:id", web.IndexController.Article)
		// 搜索
		w.GET("/search", web.IndexController.Search)
	}
}
