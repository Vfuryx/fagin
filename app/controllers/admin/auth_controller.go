package admin

import (
	"encoding/json"
	"fagin/app/caches"
	"fagin/app/errno"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/services"
	"fagin/pkg/auths"
	"fagin/pkg/cache"
	"fagin/pkg/request"
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
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	// 验证验证码
	if !utils.VerifyCaptcha(r.ID, r.Code, true) {
		ctr.ResponseJSONErr(ctx, errno.CtxVerifyCaptchaErr, nil)
		return
	}

	userID, err := services.AdminAuth.Login(r.Name, r.Password)
	if err != nil {
		ctr.ResponseJSONErr(ctx, err, nil)
		return
	}

	token, err := services.AdminAuth.Sign(userID, r.Name, "")
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxTokenInvalidErr, err)
		return
	}

	go func() {
		// 记录登录日志
		_ = services.AdminLoginLog.LogStore(r.Name, ctx.ClientIP(), ctx.Request.UserAgent())
	}()

	ctr.ResponseJSONOK(ctx, gin.H{
		"token": token,
	})
}

// Logout 后台退出
func (ctr *authController) Logout(ctx *gin.Context) {
	uid := auths.GetAdminID(ctx)
	_ = services.AdminAuth.Logout(uid)

	ctr.ResponseJSONOK(ctx, nil)
}

// Info 获取管理员信息
func (ctr *authController) Info(ctx *gin.Context) {
	uid := auths.GetAdminID(ctx)

	adminUser, err := services.AdminUserService.
		UserInfoByID(uid, []string{"id", "username", "email", "avatar", "status", "home_path"})
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.ReqErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{
		"userId":   adminUser.ID,
		"username": adminUser.Username,
		"realName": adminUser.Username,
		"email":    adminUser.Email,
		"avatar":   adminUser.Avatar,
		"homePath": adminUser.HomePath,
	})
}

// UpdateAdminUser 更新用户信息
func (ctr *authController) UpdateAdminUser(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	r := adminRequest.NewUpdateAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	data := map[string]interface{}{
		"email": r.Email,
	}

	err = services.AdminUserService.UpdateAdminUser(id, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// GetCaptcha 获取验证码
func (ctr *authController) GetCaptcha(ctx *gin.Context) {
	id, b64s, err := utils.NewCaptcha()

	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.CtxGetCaptchaErr, nil)
	}

	ctr.ResponseJSONOK(ctx, gin.H{
		"data": b64s,
		"id":   id,
	})
}

// Routes 动态获取路由
func (ctr *authController) Routes(ctx *gin.Context) {
	// 获取userID
	userID := auths.GetAdminID(ctx)

	ca := caches.NewAdminRoutesCache(func() ([]byte, error) {
		navs, err := services.AdminUserService.GetRoutes(userID)
		if err != nil {
			return nil, err
		}
		data := response.NewAdminNavList(navs...).Collection()
		return json.Marshal(data)
	})

	data, err := ca.Get(userID)
	if err != nil && err != cache.ErrNotOpen {
		ctr.ResponseJSONErrLog(ctx, errno.ReqErr, err)
		return
	}

	var res []map[string]interface{}

	err = json.Unmarshal(data, &res)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.ReqErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, res)
}

// PermissionCode 获取用户权限
func (ctr *authController) PermissionCode(ctx *gin.Context) {
	// 获取userID
	userID := auths.GetAdminID(ctx)

	var navs, err = services.AdminUserService.GetPermissionCode(userID)
	if err != nil && err != cache.ErrNotOpen {
		ctr.ResponseJSONErrLog(ctx, errno.ReqErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, navs)
}
