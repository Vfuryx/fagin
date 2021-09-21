package routes

import (
	"fagin/app"
	"fagin/app/errno"
	_ "fagin/docs/swag"
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
		app.ResponseJsonWithStatus(ctx, http.StatusNotFound, errno.NotFound, nil, gin.H{"m": "swagger 404"})
	})

	// 访问 swagger 首页必须是 /swagger/index.html
	Swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
