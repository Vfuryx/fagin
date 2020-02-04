package routes

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/pkg/router"
	"github.com/gin-gonic/gin"
	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
	"net/http"
)

func init() {
	// 生产环境禁止加载
	if app.IsProd() {
		return
	}

	var Swagger = router.Group("swagger")

	router.NoRoute(Swagger.BasePath(), func(ctx *gin.Context) {
		app.JsonResponseWithStatus(ctx, http.StatusNotFound, errno.Api.NotFound, gin.H{"m": "swagger 404"})
	})

	// use ginSwagger middleware to
	Swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
