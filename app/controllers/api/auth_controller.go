package api

import (
	"fagin/app/errno"
	"fagin/app/services"
	"net/http"

	"github.com/gin-gonic/gin"
)

type authController struct {
	BaseController
}

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
// @Success 200 {object} response.Response "正确 {"code":0,"message":"OK","data":{"token":"XXXXX"}} <br/> 错误 {"code": 20102,"message": "找不到用户"} <br/> "{code": 20105, "message": "密码不正确"}"
// @Router /api/v1/login [post]
func (ctr *authController) Login(ctx *gin.Context) {
	var r LoginRequest

	if err := ctx.ShouldBind(&r); err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	user, err := services.Auth.Login(r.Email, r.Password)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, err, err)
		return
	}

	token, err := services.Token.Sign(user, "")
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxTokenInvalidErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{
		"token": token,
	})
}

func (ctr *authController) Logout(ctx *gin.Context) {
	ctx.String(http.StatusOK, "logout")
}

func (ctr *authController) CreateUser(ctx *gin.Context) {
}
