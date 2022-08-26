package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/services"
	"fagin/pkg/db"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type userController struct {
	BaseController
}

var UserController userController

func (ctr *userController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, DefaultLimit)

	r, msg := request.Validation[adminRequest.AdminUserList](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
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

	columns := []string{"*"}

	users, err := services.AdminUserService.Index(params, columns, nil, paginator)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.NewAdminUserList(users...).Collection()

	ctr.ResponseJSONOK(ctx, gin.H{
		"users":     data,
		"paginator": paginator,
	})
}

func (ctr *userController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"*"}

	user, err := services.AdminUserService.Show(id, columns, nil)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
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
		// "role_id":   user.RoleID,
		"create_at": user.CreatedAt,
	})
}

func (ctr *userController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.AdminUserService.Delete(id)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

type userIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

func (ctr *userController) Dels(ctx *gin.Context) {
	var ids userIDs

	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.AdminUserService.Deletes(ids.IDs)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *userController) Store(ctx *gin.Context) {
	r, msg := request.Validation[adminRequest.CreateUser](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	var role admin_role.AdminRole

	err := admin_role.NewDao().Query(gin.H{"id": r.RoleID}, []string{"*"}, nil).First(&role)
	if err != nil || role.ID == 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
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
	}

	err = services.AdminUserService.Create(&user, nil)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *userController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	r, msg := request.Validation[adminRequest.UpdateUser](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	var role admin_role.AdminRole

	err = admin_role.NewDao().Query(gin.H{"id": r.RoleID}, []string{"*"}, nil).First(&role)
	if err != nil || role.ID == 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	data := map[string]interface{}{
		"status":    *r.Status,
		"remark":    r.Remark,
		"username":  r.Username,
		"nick_name": r.NickName,
		"phone":     r.Phone,
		"email":     r.Email,
		// "password":  pw,
		"sex": *r.Sex,
		// "role_id": r.RoleID,
	}

	err = services.AdminUserService.Update(id, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

type updateAdminUserStatus struct {
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

func (ctr *userController) UpdateStatus(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	var r updateAdminUserStatus

	err = ctx.ShouldBind(&r)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
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

func (ctr *userController) Reset(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	r, msg := request.Validation[adminRequest.ResetAdminUser](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
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
