package admin

import (
	"encoding/json"
	"fagin/app/caches"
	"fagin/app/errno"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/auths"
	"fagin/pkg/cache"
	"fagin/utils"
	"github.com/gin-gonic/gin"
)

type authController struct {
	BaseController
}

// AuthController 授权控制器
var AuthController authController

// Login 后台登录
func (ctr *authController) Login(ctx *gin.Context) {
	var r = adminRequest.NewLoginRequest()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	// 验证验证码
	if !utils.VerifyCaptcha(r.ID, r.Code, true) {
		ctr.ResponseJsonErr(ctx, errno.CtxVerifyCaptchaErr, nil)
		return
	}

	userID, err := service.AdminAuth.Login(r.Name, r.Password)
	if err != nil {
		ctr.ResponseJsonErr(ctx, err, nil)
		return
	}

	token, err := service.AdminAuth.Sign(userID, r.Name, "")
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxTokenInvalidErr, err)
		return
	}

	go func() {
		// 记录登录日志
		_ = service.AdminLoginLog.LogStore(r.Name, ctx.ClientIP(), ctx.Request.UserAgent())
	}()

	ctr.ResponseJsonOK(ctx, gin.H{
		"token": token,
	})
}

// Logout 后台退出
func (ctr *authController) Logout(ctx *gin.Context) {
	uid := auths.GetAdminID(ctx)
	_ = service.AdminAuth.Logout(uint(uid))
	ctr.ResponseJsonOK(ctx, nil)
}

// Info 获取管理员信息
func (ctr *authController) Info(ctx *gin.Context) {
	uid := auths.GetAdminID(ctx)
	adminUser, err := service.AdminUserService.
		UserInfoByID(uint(uid), []string{"id", "username", "email", "avatar", "status"})
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.ReqErr, err)
		return
	}
	ctr.ResponseJsonOK(ctx, gin.H{
		"userId":   adminUser.ID,
		"username": adminUser.Username,
		"realName": adminUser.Username,
		"email":    adminUser.Email,
		"avatar":   adminUser.Avatar,
		"roles":    []string{"admin"},
	})
}

// UpdateAdminUser 更新用户信息
func (ctr *authController) UpdateAdminUser(ctx *gin.Context) {
	//id, err := request.ShouldBindUriUintID(ctx)
	//if err != nil {
	//	go app.Log().Errorln(err, string(debug.Stack()))
	//	app.ResponseJson(ctx, errno.ReqErr, nil)
	//	return
	//}
	//r := new(admin_request.UpdateAdminUser)
	//if data, ok := r.Validate(ctx); !ok {
	//	app.ResponseJson(ctx, errno.ReqErr, data)
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
	//		go app.Log().Errorln(err, string(debug.Stack()))
	//		app.ResponseJson(ctx, err)
	//		return
	//	}
	//	password, err := app.Encrypt(r.NewPassword)
	//	if err != nil {
	//		go app.Log().Errorln(err, string(debug.Stack()))
	//		app.ResponseJson(ctx, err)
	//		return
	//	}
	//	data["password"] = password
	//}
	//
	//err = service.AdminUserService.UpdateAdminUser(id, data)
	//if err != nil {
	//	go app.Log().Errorln(err, string(debug.Stack()))
	//	app.ResponseJson(ctx, errno.ErrUpdateUser, nil)
	//	return
	//}
	//
	//app.ResponseJson(ctx, errno.OK, nil)
}

// GetCaptcha 获取验证码
func (ctr *authController) GetCaptcha(ctx *gin.Context) {
	id, b64s, err := utils.NewCaptcha()

	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.CtxGetCaptchaErr, nil)
	}

	ctr.ResponseJsonOK(ctx, gin.H{
		"data": b64s,
		"id":   id,
	})
}

// Routes 动态获取路由
func (ctr *authController) Routes(ctx *gin.Context) {
	// 获取userID
	userID := auths.GetAdminID(ctx)

	ca := caches.NewAdminRoutesCache(func() ([]byte, error) {
		navs, err := service.AdminUserService.GetRoutes(userID)
		if err != nil {
			return nil, err
		}
		data := response.AdminNavList(navs...).Collection()
		return json.Marshal(data)
	})

	data, err := ca.Get(userID)
	if err != nil && err != cache.ErrNotOpen {
		ctr.ResponseJsonErrLog(ctx, errno.ReqErr, err)
		return
	}
	var res []map[string]interface{}
	err = json.Unmarshal(data, &res)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.ReqErr, err)
		return
	}
	ctr.ResponseJsonOK(ctx, res)
}

// PermissionCode 获取用户权限
func (ctr *authController) PermissionCode(ctx *gin.Context) {
	// 获取userID
	userID := auths.GetAdminID(ctx)

	navs, err := service.AdminUserService.GetPermissionCode(uint(userID))
	if err != nil && err != cache.ErrNotOpen {
		ctr.ResponseJsonErrLog(ctx, errno.ReqErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, navs)
}
