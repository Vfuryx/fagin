package api

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type authController struct{}

var AuthController authController

type LoginRequest struct {
	Email    string `form:"email" json:"email" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// Login
// @Summary 登录接口
// @Description 用户登录
// @Tags auth
// @Accept  mpfd
// @Produce  json
// @Param email formData string true "用户邮箱地址" minLength(6) maxLength(16)
// @Param password formData string true "用户密码" minLength(6) maxLength(32)
// @Success 200 {object} app.Response "正确 {"code":0,"message":"OK","data":{"token":"XXXXX"}} <br/> 错误 {"code": 20102,"message": "找不到用户"} <br/> "{code": 20105, "message": "密码不正确"}"
// @Router /api/v1/login [post]
func (authController) Login(ctx *gin.Context) {
	var r LoginRequest

	if err := ctx.ShouldBind(&r); err != nil {
		app.JsonResponse(ctx, errno.ErrBind, nil)
		return
	}

	user, err := service.Auth.Login(r.Email, r.Password)
	if err != nil {
		app.JsonResponse(ctx, err, nil)
		return
	}

	token, err := service.Token.Sign(user, "")
	if err != nil {
		app.JsonResponse(ctx, errno.ErrToken, nil)
		return
	}

	app.JsonResponse(ctx, nil, gin.H{
		"token": token,
	})
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
