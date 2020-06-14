package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/service"
	"fagin/app/service/admin_auth"
	"fagin/app/utils"
	"fagin/pkg/log"
	"fmt"
	"github.com/gin-gonic/gin"
)

type authController struct{}

var AuthController authController

type LoginRequest struct {
	Name     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
	ID       string `form:"id" json:"id" binding:"required"`
	Code     string `form:"code" json:"code" binding:"required"`
}

// 后台登录
func (authController) Login(ctx *gin.Context) {
	var r LoginRequest

	if err := ctx.ShouldBind(&r); err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	// 验证验证码
	if !utils.VerifyCaptcha(r.ID, r.Code, true) {
		app.JsonResponse(ctx, errno.Api.ErrVerifyCaptcha, nil)
		return
	}

	userID, err := admin_auth.AdminAuth.Login(r.Name, r.Password)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, err, nil)
		return
	}
	token, err := admin_auth.AdminAuth.Sign(userID, r.Name, "")
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrToken, nil)
		return
	}

	app.JsonResponse(ctx, nil, gin.H{
		"token": token,
	})
}

// 后台退出
func (authController) Logout(ctx *gin.Context) {
	app.JsonResponse(ctx, errno.OK, nil)
}

// 获取管理员信息
func (authController) Info(ctx *gin.Context) {
	name := ctx.GetString("user_name")
	adminUser, err := service.AdminUserService.UserInfo(name, []string{"id", "username", "email", "avatar", "status"})
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, err, nil)
		return
	}
	app.JsonResponse(ctx, errno.OK, gin.H{
		"name":   adminUser.Username,
		"email":  adminUser.Email,
		"avatar": adminUser.Avatar,
		"roles":  []string{"admin"},
	})
}

//更新用户信息
func (authController) UpdateAdminUser(ctx *gin.Context) {
	//id, err := request.ShouldBindUriUintID(ctx)
	//if err != nil {
	//	log.Log.Errorln(err)
	//	app.JsonResponse(ctx, errno.Api.ErrBind, nil)
	//	return
	//}
	//r := new(admin_request.UpdateAdminUser)
	//if data, ok := r.Validate(ctx); !ok {
	//	app.JsonResponse(ctx, errno.Api.ErrBind, data)
	//	return
	//}
	//
	//data := map[string]interface{}{
	//	"email": r.Email,
	//}
	//
	//if r.OldPassword != "" {
	//	err := service.AdminUserService.CheckPassword(id, r.OldPassword)
	//	if err != nil {
	//		log.Log.Errorln(err)
	//		app.JsonResponse(ctx, err, nil)
	//		return
	//	}
	//	password, err := app.Encrypt(r.NewPassword)
	//	if err != nil {
	//		log.Log.Errorln(err)
	//		app.JsonResponse(ctx, err, nil)
	//		return
	//	}
	//	data["password"] = password
	//}
	//
	//err = service.AdminUserService.UpdateAdminUser(id, data)
	//if err != nil {
	//	log.Log.Errorln(err)
	//	app.JsonResponse(ctx, errno.Api.ErrUpdateUser, nil)
	//	return
	//}
	//
	//app.JsonResponse(ctx, errno.OK, nil)
}

// 获取验证码
func (authController) GetCaptcha(ctx *gin.Context) {
	id, b64s, err := utils.NewCaptcha()

	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrGetCaptcha, nil)
	}

	app.JsonResponse(ctx, errno.OK, gin.H{
		"data": b64s,
		"id":   id,
	})
}

func (authController) MenuRole(ctx *gin.Context) {
	fmt.Println(ctx.GetInt64("admin_user_id"))
	app.JsonResponse(ctx, errno.OK, gin.H{
		"menu": "123",
	})

}
