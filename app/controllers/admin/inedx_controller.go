package admin

import (
	"fagin/app"
	"github.com/gin-gonic/gin"
)

type indexController struct{}

var IndexController indexController

// Index 后台首页
func (*indexController) Index(ctx *gin.Context) {
	app.View(ctx, "admin.layout", nil)
}
