package controllers

import (
	User "github.com/Vfuryx/fagin/app/models/user"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

type AuthController struct {

}

func (AuthController) Login(ctx *gin.Context)  {
	ctx.String(http.StatusOK, "login")
}

func (AuthController) Logout(ctx *gin.Context)  {
	ctx.String(http.StatusOK, "logout")
}

func (AuthController) CreateUser(ctx *gin.Context) {
	user := ctx.DefaultPostForm("name", "1")
	avatar := ctx.DefaultPostForm("avatar", "1")
	age := ctx.DefaultPostForm("age", "1")
	int,_ := strconv.Atoi(age)

	userModel := User.User{
		Name: user,
		Avatar: avatar,
		Age: uint8(int),
	}

	User.Create(&userModel)

}