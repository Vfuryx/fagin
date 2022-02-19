package routes

import (
	"github.com/gin-gonic/gin"
)

type routeFunc func(*gin.RouterGroup)

func Handle(e *gin.Engine) {
	// 监控路由 metricsRoute(e.Group("/metrics"))
	webRoute(e.Group("/"))
	apiRoute(e.Group("/api"))
	adminRoute(e.Group("/admin"))
	swaggerRoute(e.Group("/swagger"))
}
