package routes

import (
	"github.com/gin-gonic/gin"
)

func Handle(e *gin.Engine) {
	webRoute(e.Group("/"))
	apiRoute(e.Group("/api"))
	adminRoute(e.Group("/admin"))
	swaggerRoute(e.Group("/swagger"))
	//metricsRoute(e.Group("/metrics"))
}
