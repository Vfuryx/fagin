package admin

import (
	"fagin/app"
	"fagin/app/caches"
	"fagin/app/errno"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user"
	adminRequest "fagin/app/requests/admin"
	adminResponse "fagin/app/responses/admin"
	"fagin/app/services"
	"fagin/pkg/cache"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"fagin/utils"

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
	paginator := db.NewPaginatorWithCtx(ctx, 1, DefaultLimit)

	var r = adminRequest.NewAdminUserList()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
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

	users, err := services.AdminUserService.Index(params, columns, with, paginator)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := adminResponse.NewAdminUserList(users...).Collection()

	ctr.ResponseJSONOK(ctx, gin.H{
		"items": data,
		"total": paginator.TotalCount,
	})
}

func (ctr *adminsController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, err)
		return
	}

	columns := []string{"*"}
	with := gin.H{"Roles": nil}

	user, err := services.AdminUserService.Show(id, columns, with)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	// 转换为角色ID组
	var roles = make([]uint, 0)

	for _, role := range user.Roles {
		roles = append(roles, role.ID)
	}

	ctr.ResponseJSONOK(ctx, gin.H{
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
		"home_path": user.HomePath,
		"create_at": user.CreatedAt,
	})
}

// Store 创建管理员
func (ctr *adminsController) Store(ctx *gin.Context) {
	var r = adminRequest.NewCreateAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	ok := services.AdminUserService.UsernameExists(r.Username, 0)
	if ok {
		ctr.ResponseJSONErr(ctx, errno.CtxUserExistErr, nil)
		return
	}

	params := gin.H{"in_id": r.Roles, "type": 0}
	// 获取权限组
	var roles []*admin_role.AdminRole

	err := admin_role.NewDao().Query(params, []string{"*"}, nil).Find(&roles)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
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
		HomePath: r.HomePath,
	}

	err = services.AdminUserService.Create(&user, roles)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Update 编辑管理员
func (ctr *adminsController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, err)
		return
	}

	var r = adminRequest.NewUpdateAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	ok := services.AdminUserService.UsernameExists(r.Username, id)
	if ok {
		ctr.ResponseJSONErr(ctx, errno.CtxUserExistErr, nil)
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
		"home_path": r.HomePath,
	}

	err = services.AdminUserService.Update(id, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ca := caches.NewAdminRoutesCache(nil)

	_, err = ca.Remove(utils.Uint64ToStr(uint64(id)))
	if err != nil && err != cache.ErrNotOpen {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Delete 删除
func (ctr *adminsController) Delete(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, err)
		return
	}

	err = services.AdminUserService.Delete(id)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

type updateAdminStatus struct {
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

func (ctr *adminsController) UpdateStatus(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, err)
		return
	}

	var r updateAdminStatus

	err = ctx.ShouldBind(&r)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, err)
		return
	}

	s := 0
	if *r.Status > 0 {
		s = 1
	}

	err = services.AdminUserService.UpdateStatus(id, s)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// ResetPassword 重置密码
func (ctr *adminsController) ResetPassword(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, err)
		return
	}

	var r = adminRequest.NewResetAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	data := map[string]interface{}{
		"password": pw,
	}

	err = services.AdminUserService.Update(id, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Logout 强制用户下线
func (ctr *adminsController) Logout(ctx *gin.Context) {
	uid, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, err)
		return
	}

	_ = services.AdminAuth.Logout(uid)

	ctr.ResponseJSONOK(ctx, nil)
}

// UsernameExists 用户名是否已存在
func (ctr *adminsController) UsernameExists(ctx *gin.Context) {
	var r = adminRequest.NewAdminUsernameExistRequest()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	ok := services.AdminUserService.UsernameExists(r.Username, r.ID)
	if ok {
		ctr.ResponseJSONErr(ctx, errno.CtxUserExistErr, nil)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// RolesRoute 角色路由
func (ctr *adminsController) RolesRoute(ctx *gin.Context) {
	var r = adminRequest.NewRolesRoute()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	navs, err := services.AdminUserService.GetRoutesByRoleIDs(r.Roles)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	data := adminResponse.NewAdminRoleRoute(navs...).Collection()

	ctr.ResponseJSONOK(ctx, data)
}
