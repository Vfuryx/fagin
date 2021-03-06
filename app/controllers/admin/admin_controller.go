package admin

import (
	"fagin/app"
	"fagin/app/caches"
	"fagin/app/errno"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user"
	"fagin/app/requests/admin"
	adminResponse "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/app/service/admin_auth"
	"fagin/pkg/cache"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type adminsController struct {
	BaseController
}

var AdminsController adminsController

// Index
// @Summary 管理员列表
// @Description 管理员列表接口
// @Tags 总后台接口
// @Accept  json
// @Produce  json
// @Param username query string false "用户名称" maxlength(64)
// @Param phone query string false "手机号码" minlength(11) maxlength(11)
// @Param status query int false "状态" Enums(0, 1)
// @Success 200 {object} response.Response "正确 {"code":0,"message":"OK","data":{"token":"XXXXX"}} <br/> 错误 {"code": 20102,"message": "找不到用户"} <br/> "{code": 20105, "message": "密码不正确"}"
// @Router /admin/api/v1/admins/ [post]
func (ac *adminsController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	var r = admin_request.NewAdminUserList()
	if data, ok := r.Validate(ctx); !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	params := gin.H{}
	if r.Username != "" {
		params["like_username"] = "%" + r.Username + "%"
	}
	if r.Phone != "" {
		params["like_phone"] = "%" + r.Phone + "%"
	}
	if r.Status != nil {
		params["status"] = *r.Status
	}
	columns := []string{"id", "username", "status", "phone", "nick_name"}
	with := gin.H{
		"Roles": func(db *gorm.DB) *gorm.DB {
			return db.Select([]string{"id", "name"}).Where("type = ?", 0)
		},
	}

	users, err := service.AdminUserService.Index(params, columns, with, &paginator)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
		return
	}

	data := adminResponse.AdminUserList(users...).Collection()

	ac.ResponseJsonOK(ctx, gin.H{
		"users":     data,
		"paginator": paginator,
	})
	return
}

func (ac *adminsController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, err)
		return
	}

	columns := []string{"*"}
	with := gin.H{"Roles": func(db *gorm.DB) *gorm.DB {
		return db.Where("type = ?", 0)
	}}
	user, err := service.AdminUserService.Show(id, columns, with)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.ShowErr, err, nil)
		return
	}

	// 转换为角色ID组
	var roles []uint
	for _, role := range user.Roles {
		roles = append(roles, role.ID)
	}

	ac.ResponseJsonOK(ctx, gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"nick_name": user.NickName,
		"email":     user.Email,
		"phone":     user.Phone,
		"avatar":    user.Avatar,
		"sex":       user.Sex,
		"remark":    user.Remark,
		"status":    user.Status,
		"roles":     roles,
		"create_at": user.CreatedAt,
	})
	return
}

func (ac *adminsController) Delete(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, err)
		return
	}

	err = service.AdminUserService.Delete(id)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

// Store 创建管理员
func (ac *adminsController) Store(ctx *gin.Context) {
	var r = admin_request.NewCreateAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	ok := service.AdminUserService.UsernameExist(r.Username, 0)
	if !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.AdminUsernameExistErr, nil)
		return
	}

	params := gin.H{"in_id": r.Roles, "type": 0}
	// 获取权限组
	var roles []admin_role.AdminRole
	err := admin_role.Dao().Query(params, []string{"*"}, nil).Find(&roles)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	user := admin_user.AdminUser{
		Status:   *r.Status,
		Username: r.Username,
		NickName: r.NickName,
		Phone:    r.Phone,
		Email:    r.Email,
		Password: pw,
		Remark:   r.Remark,
		Sex:      *r.Sex,
		Roles:    roles,
	}

	err = service.AdminUserService.Create(&user, roles)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

// Update 编辑管理员
func (ac *adminsController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, err)
		return
	}

	var r = admin_request.NewUpdateAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	ok := service.AdminUserService.UsernameExist(r.Username, id)
	if !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.AdminUsernameExistErr, nil)
		return
	}

	data := map[string]interface{}{
		"status":    *r.Status,
		"remark":    r.Remark,
		"username":  r.Username,
		"nick_name": r.NickName,
		"phone":     r.Phone,
		"email":     r.Email,
		"sex":       *r.Sex,
		"role_ids":  r.Roles,
	}
	err = service.AdminUserService.Update(id, data)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	ca := caches.NewAdminNavsCache(nil)
	_, err = ca.Remove(strconv.FormatUint(uint64(id), 10))
	if err != nil && err != cache.NotOpenErr {
		ac.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

type updateAdminStatus struct {
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

func (ac *adminsController) UpdateStatus(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, err)
		return
	}

	var r updateAdminStatus
	err = ctx.ShouldBind(&r)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, err)
		return
	}
	s := 0
	if *r.Status > 0 {
		s = 1
	}

	err = service.AdminUserService.UpdateStatus(id, s)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

// ResetPassword 重置密码
func (ac *adminsController) ResetPassword(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, err)
		return
	}

	var r = admin_request.NewResetAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	data := map[string]interface{}{
		"password": pw,
	}
	err = service.AdminUserService.Update(id, data)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

// Logout 强制用户下线
func (ac *adminsController) Logout(ctx *gin.Context) {
	uid, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, err)
		return
	}
	_ = admin_auth.AdminAuth.Logout(uid)
	ac.ResponseJsonOK(ctx, nil)
}

// UsernameExist 用户名是否已存在
func (ac *adminsController) UsernameExist(ctx *gin.Context) {
	var r = admin_request.NewAdminUsernameExistRequest()
	if data, ok := r.Validate(ctx); !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	ok := service.AdminUserService.UsernameExist(r.Username, r.ID)
	if !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.AdminUsernameExistErr, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
}
