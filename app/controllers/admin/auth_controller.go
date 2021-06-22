package admin

import (
	"encoding/json"
	"fagin/app/caches"
	"fagin/app/errno"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/app/service/admin_auth"
	"fagin/app/utils"
	"fagin/pkg/auths"
	"fagin/pkg/cache"
	"github.com/gin-gonic/gin"
	"strconv"
)

type authController struct {
	BaseController
}

var AuthController authController

// Login 后台登录
func (ac *authController) Login(ctx *gin.Context) {
	var r = adminRequest.NewLoginRequest()
	if data, ok := r.Validate(ctx); !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	// 验证验证码
	if !utils.VerifyCaptcha(r.ID, r.Code, true) {
		ac.ResponseJsonErr(ctx, errno.Serve.ErrVerifyCaptcha, nil)
		return
	}

	userID, err := admin_auth.AdminAuth.Login(r.Name, r.Password)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, err, err, nil)
		return
	}

	token, err := admin_auth.AdminAuth.Sign(userID, r.Name, "")
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.ErrToken, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, gin.H{
		"token": token,
	})
}

// Logout 后台退出
func (ac *authController) Logout(ctx *gin.Context) {
	uid := auths.GetAdminID(ctx)
	_ = admin_auth.AdminAuth.Logout(uint(uid))
	ac.ResponseJsonOK(ctx, nil)
}

// Info 获取管理员信息
func (ac *authController) Info(ctx *gin.Context) {
	uid := auths.GetAdminID(ctx)
	adminUser, err := service.AdminUserService.
		UserInfoById(uint(uid), []string{"id", "username", "email", "avatar", "status"})
	if err != nil {
		ac.ResponseJsonErrLog(ctx, err, err, nil)
		return
	}
	ac.ResponseJsonOK(ctx, gin.H{
		"name":   adminUser.Username,
		"email":  adminUser.Email,
		"avatar": adminUser.Avatar,
		"roles":  []string{"admin"},
	})
}

// UpdateAdminUser 更新用户信息
func (ac *authController) UpdateAdminUser(ctx *gin.Context) {
	//id, err := request.ShouldBindUriUintID(ctx)
	//if err != nil {
	//	go app.Log().Errorln(err, string(debug.Stack()))
	//	app.ResponseJson(ctx, errno.Serve.BindErr, nil)
	//	return
	//}
	//r := new(admin_request.UpdateAdminUser)
	//if data, ok := r.Validate(ctx); !ok {
	//	app.ResponseJson(ctx, errno.Serve.BindErr, data)
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
	//		app.ResponseJson(ctx, err, nil)
	//		return
	//	}
	//	password, err := app.Encrypt(r.NewPassword)
	//	if err != nil {
	//		go app.Log().Errorln(err, string(debug.Stack()))
	//		app.ResponseJson(ctx, err, nil)
	//		return
	//	}
	//	data["password"] = password
	//}
	//
	//err = service.AdminUserService.UpdateAdminUser(id, data)
	//if err != nil {
	//	go app.Log().Errorln(err, string(debug.Stack()))
	//	app.ResponseJson(ctx, errno.Serve.ErrUpdateUser, nil)
	//	return
	//}
	//
	//app.ResponseJson(ctx, errno.OK, nil)
}

// GetCaptcha 获取验证码
func (ac *authController) GetCaptcha(ctx *gin.Context) {
	id, b64s, err := utils.NewCaptcha()

	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.ErrGetCaptcha, nil)
	}

	ac.ResponseJsonOK(ctx, gin.H{
		"data": b64s,
		"id":   id,
	})
}

// Navs 动态获取路由
func (ac *authController) Navs(ctx *gin.Context) {
	// 获取userID
	userID := auths.GetAdminID(ctx)

	ca := caches.NewAdminNavsCache(func(key string) ([]byte, error) {
		navs, err := service.AdminUserService.GetNavs(uint(userID))
		if err != nil {
			return nil, err
		}
		data := response.AdminNavList(navs...).Collection()
		return json.Marshal(data)
	})

	data, err := ca.Get(strconv.FormatUint(userID, 10))
	if err != nil && err != cache.NotOpenErr {
		ac.ResponseJsonErrLog(ctx, errno.Serve.BindErr, err, nil)
		return
	}
	var res []map[string]interface{}
	err = json.Unmarshal([]byte(data), &res)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.BindErr, err, nil)
		return
	}
	ac.ResponseJsonOK(ctx, res)
}
