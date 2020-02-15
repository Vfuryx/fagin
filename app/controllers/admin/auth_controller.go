package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/requests/admin"
	"fagin/app/service"
	"fagin/app/service/admin_auth"
	"fagin/pkg/log"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type authController struct{}

var AuthController authController

type LoginRequest struct {
	Name     string `form:"username" json:"username" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

// 后台登录
func (authController) Login(ctx *gin.Context) {
	var r LoginRequest

	if err := ctx.ShouldBind(&r); err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	err := admin_auth.AdminAuth.Login(r.Name, r.Password)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, err, nil)
		return
	}
	token, err := admin_auth.AdminAuth.Sign(r.Name, "")
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
	adminUser, err := service.AdminUser.UserInfo(name, []string{"id", "username", "email", "avatar", "status"})
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, err, nil)
		return
	}
	app.JsonResponse(ctx, errno.OK, gin.H{
		"name":   adminUser.Username,
		"email":  adminUser.Email,
		"avatar": adminUser.Avatar,
	})
}

//更新用户信息
func (authController) UpdateAdminUser(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}
	r := new(admin_request.UpdateAdminUser)
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}

	data := map[string]interface{}{
		"email": r.Email,
	}

	if r.OldPassword != "" {
		err := service.AdminUser.CheckPassword(id, r.OldPassword)
		if err != nil {
			log.Log.Errorln(err)
			app.JsonResponse(ctx, err, nil)
			return
		}
		password, err := app.Encrypt(r.NewPassword)
		if err != nil {
			log.Log.Errorln(err)
			app.JsonResponse(ctx, err, nil)
			return
		}
		data["password"] = password
	}

	err = service.AdminUser.UpdateAdminUser(id, data)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrUpdateUser, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
}
