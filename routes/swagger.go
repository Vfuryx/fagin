package routes

import (
	"fagin/app"
	"fagin/app/errno"
	_ "fagin/docs"
	"fagin/pkg/router/no_router"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"net/http"
)

var swaggerRoute = func(Swagger *gin.RouterGroup) {
	// 生产环境禁止加载
	if app.IsProd() {
		return
	}

	no_router.NoRoute(Swagger.BasePath(), func(ctx *gin.Context) {
		app.ResponseJsonWithStatus(ctx, http.StatusNotFound, errno.Serve.NotFound, nil, gin.H{"m": "swagger 404"})
	})

	// use ginSwagger middleware to
	Swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
