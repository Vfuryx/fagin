package routes

import (
	"fagin/app/controllers/web"
	"fagin/pkg/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Web = router.Group("/")

func init() {
	// 404
	router.NoRoute(Web.BasePath(), func(ctx *gin.Context) {
		ctx.String(http.StatusNotFound, "404 NotFound")
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

	// 首页
	Web.GET("/", web.VideoController.VideoList)
	// banner
	Web.GET("/banner", web.IndexController.Index)
	// 视频播放
	Web.GET("/video/av:id", web.VideoController.VideoShow)
	
}
