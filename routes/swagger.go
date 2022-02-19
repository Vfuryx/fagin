package routes

import (
	"fagin/app"
	"fagin/app/errno"
	_ "fagin/docs/swag" // swag 初始化
	"fagin/pkg/router/no_router"
	"net/http"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

var swaggerRoute routeFunc = func(Swagger *gin.RouterGroup) {

	// 生产环境禁止加载
	if app.IsProd() {
		return
	}

	no_router.NoRoute(Swagger.BasePath(), func(ctx *gin.Context) {
		app.ResponseJSONWithStatus(ctx, http.StatusNotFound, errno.NotFoundErr, nil, gin.H{"m": "swagger 404"})
	})

	// 访问 swagger 首页必须是 /swagger/index.html
	Swagger.GET("/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))
}
