package controllers

import (
	"fagin/app/models/info"
	"fagin/app/models/user"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authController struct{}

var AuthController = authController{}

type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// @Summary 登录接口
// @Description 用户登录
// @Tags auth
// @Accept  json
// @Produce  json
// @Param email formData string true "用户邮箱地址"
// @Param password formData string true "用户密码"
// @Success 200 {string} json "{"code":0,"message":"OK","data":{"token":"XXXXX"}}"
// @Router /api/v1/login [post]
func (authController) Login(ctx *gin.Context) {

	u := user.User{
		Username:	"xfy",
		Info: 		info.Info{
			Email:		"12121@qq.com",
		},
	}

	db.ORM.Create(&u)

	//
	//var r LoginRequest
	//
	//if err := ctx.ShouldBind(&r); err != nil {
	//	app.JosnResponse(ctx, errno.ErrBind, nil)
	//	return
	//}
	//
	//user, err := service.Auth.Login(r.Email, r.Password)
	//
	//if err != nil {
	//	app.JosnResponse(ctx, err, nil)
	//	return
	//}
	//
	//token, err := service.Token.Sign(user, "")
	//
	//if err != nil {
	//	app.JosnResponse(ctx, errno.ErrToken, nil)
	//	return
	//}
	//
	//app.JosnResponse(ctx, nil, gin.H{
	//	"token": token,
	//})
}

func (authController) Logout(ctx *gin.Context) {
	ctx.String(http.StatusOK, "logout")
}

func (authController) CreateUser(ctx *gin.Context) {
	//user := ctx.DefaultPostForm("name", "1")
	//avatar := ctx.DefaultPostForm("avatar", "1")
	//age := ctx.DefaultPostForm("age", "1")
	//int,_ := strconv.Atoi(age)

	//userModel := User.User{
	//	Name: user,
	//	Avatar: avatar,
	//	Age: uint8(int),
	//}

	//User.Create(&userModel)
}
