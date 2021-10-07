package admin

import (
	"fagin/resources/static"
	"github.com/gin-gonic/gin"
	"net/http"
)

type indexController struct{}

var IndexController indexController

// Index 后台首页
func (*indexController) Index(ctx *gin.Context) {
	ctx.Data(http.StatusOK, "text/html;charset=UTF-8", static.AdminHTML)
}
