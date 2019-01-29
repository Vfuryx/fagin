package controllers

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type IndexController struct {

}

func (IndexController) Index(ctx *gin.Context) {
	ctx.String(http.StatusOK, "index")
}



