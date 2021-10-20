package admin

import (
	"fagin/app"
	"fagin/app/caches"
	"fagin/app/errno"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user"
	adminRequest "fagin/app/requests/admin"
	adminResponse "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/cache"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type adminsController struct {
	BaseController
}

// AdminsController 管理员控制器
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
func (ctr *adminsController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	var r = adminRequest.NewAdminUserList()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
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

	users, err := service.AdminUserService.Index(params, columns, with, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := adminResponse.AdminUserList(users...).Collection()

	ctr.ResponseJsonOK(ctx, gin.H{
		"items": data,
		"total": paginator.TotalCount,
	})
	return
}

func (ctr *adminsController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, err)
		return
	}

	columns := []string{"*"}
	with := gin.H{"Roles": nil}
	user, err := service.AdminUserService.Show(id, columns, with)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	// 转换为角色ID组
	var roles = make([]uint, 0, 5)
	for _, role := range user.Roles {
		roles = append(roles, role.ID)
	}

	ctr.ResponseJsonOK(ctx, gin.H{
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

// Store 创建管理员
func (ctr *adminsController) Store(ctx *gin.Context) {
	var r = adminRequest.NewCreateAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	ok := service.AdminUserService.UsernameExists(r.Username, 0)
	if ok {
		ctr.ResponseJsonErr(ctx, errno.CtxUserExistErr, nil)
		return
	}

	params := gin.H{"in_id": r.Roles, "type": 0}
	// 获取权限组
	var roles []admin_role.AdminRole
	err := admin_role.NewDao().Query(params, []string{"*"}, nil).Find(&roles)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
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
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Update 编辑管理员
func (ctr *adminsController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, err)
		return
	}

	var r = adminRequest.NewUpdateAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	ok := service.AdminUserService.UsernameExists(r.Username, id)
	if ok {
		ctr.ResponseJsonErr(ctx, errno.CtxUserExistErr, nil)
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
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ca := caches.NewAdminRoutesCache(nil)
	_, err = ca.Remove(strconv.FormatUint(uint64(id), 10))
	if err != nil && err != cache.ErrNotOpen {
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Delete 删除
func (ctr *adminsController) Delete(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, err)
		return
	}

	err = service.AdminUserService.Delete(id)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

type updateAdminStatus struct {
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

func (ctr *adminsController) UpdateStatus(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, err)
		return
	}

	var r updateAdminStatus
	err = ctx.ShouldBind(&r)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, err)
		return
	}
	s := 0
	if *r.Status > 0 {
		s = 1
	}

	err = service.AdminUserService.UpdateStatus(id, s)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// ResetPassword 重置密码
func (ctr *adminsController) ResetPassword(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, err)
		return
	}

	var r = adminRequest.NewResetAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	data := map[string]interface{}{
		"password": pw,
	}
	err = service.AdminUserService.Update(id, data)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Logout 强制用户下线
func (ctr *adminsController) Logout(ctx *gin.Context) {
	uid, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, err)
		return
	}
	_ = service.AdminAuth.Logout(uid)
	ctr.ResponseJsonOK(ctx, nil)
}

// UsernameExists 用户名是否已存在
func (ctr *adminsController) UsernameExists(ctx *gin.Context) {
	var r = adminRequest.NewAdminUsernameExistRequest()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	ok := service.AdminUserService.UsernameExists(r.Username, r.ID)
	if ok {
		ctr.ResponseJsonErr(ctx, errno.CtxUserExistErr, nil)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
}
