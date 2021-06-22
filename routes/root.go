package routes

import (
	"github.com/gin-gonic/gin"
)

func Handle(e *gin.Engine) {
	webRoute(e.Group("/"))
	adminRoute(e.Group("/admin"))
	apiRoute(e.Group("/api"))
	swaggerRoute(e.Group("/swagger"))
	//metricsRoute(e.Group("/metrics"))
}
