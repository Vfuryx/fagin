package api

import (
	"fagin/app/errno"
	api_request "fagin/app/requests/api"

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
	var v = api_request.NewIndexRequest()

	if data, ok := v.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}
}
