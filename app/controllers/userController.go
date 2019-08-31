package controllers

import (
	"fagin/app"
	"fagin/app/errno"
	"github.com/gin-gonic/gin"
)

type userController struct{}

var UserController = userController{}

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (this userController) Register(ctx *gin.Context) {

	var r Login

	if err := ctx.ShouldBind(&r); err != nil{
		app.JosnResponse(ctx, errno.ErrBind,nil)
		return
	}

	app.JosnResponse(ctx, nil, r)

}
