package api

import (
	"fagin/app/errno"
	apiRequest "fagin/app/requests/api"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type indexController struct {
	BaseController
}

var IndexController indexController

// Index
// @Summary Add new user to the database
// @Description Add a new user
// @securityDefinitions.basic BasicAuth
// @Tags user
// @Accept  json
// @Produce  json
func (ctr *indexController) Index(ctx *gin.Context) {
	_, msg := request.Validation[apiRequest.IndexRequest](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}
}
