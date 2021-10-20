package routes

import (
	"fagin/app/controllers/web"
	"fagin/app/middleware"
	"fagin/pkg/router/no_router"
	"github.com/gin-gonic/gin"
	"net/http"
)

var webRoute = func(Web *gin.RouterGroup) {
	// 404
	no_router.NoRoute(Web.BasePath(), func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "404 NotFoundErr")
	})

	//limiter := middleware.LimitMiddleware(2, func(ctx *gin.Context) {
	//	ctx.String(http.StatusBadGateway, "服务器繁忙")
	//})
	//
	//Web.GET("/", limiter, func(ctx *gin.Context) {
	//	time.Sleep(1 * time.Minute)
	//	ctx.JSON(http.StatusOK, gin.H{
	//		"123": 123,
	//	})
	//})

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
