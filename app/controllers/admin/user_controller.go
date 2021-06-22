package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user"
	"fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type userController struct {
	BaseController
}

var UserController userController

func (uc *userController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	var r = admin_request.NewAdminUserList()
	if data, ok := r.Validate(ctx); !ok {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
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

	users, err := service.AdminUserService.Index(params, columns, nil, &paginator)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
		return
	}

	data := response.AdminUserList(users...).Collection()

	uc.ResponseJsonOK(ctx, gin.H{
		"users":     data,
		"paginator": paginator,
	})
	return
}

func (uc *userController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	columns := []string{"*"}
	user, err := service.AdminUserService.Show(id, columns, nil)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.ShowErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"nick_name": user.NickName,
		"email":     user.Email,
		"phone":     user.Phone,
		"avatar":    user.Avatar,
		"sex":       user.Sex,
		"remark":    user.Remark,
		"status":    user.Status,
		//"role_id":   user.RoleID,
		"create_at": user.CreatedAt,
	})
	return
}

func (uc *userController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.AdminUserService.Delete(id)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, nil)
	return
}

type userIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

func (uc *userController) Dels(ctx *gin.Context) {
	var ids userIDs
	err := ctx.ShouldBind(&ids)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.AdminUserService.Deletes(ids.IDs)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, nil)
	return
}

func (uc *userController) Store(ctx *gin.Context) {
	var r = admin_request.NewCreateUser()
	if data, ok := r.Validate(ctx); !ok {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	var role admin_role.AdminRole
	err := admin_role.Dao().Query(gin.H{"id": r.RoleID}, []string{"*"}, nil).First(&role)
	if err != nil || role.ID == 0 {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
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
		//Roles:   r.RoleID,
	}

	err = service.AdminUserService.Create(&user, nil)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, nil)
	return
}

func (uc *userController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	var r = admin_request.NewUpdateUser()
	if data, ok := r.Validate(ctx); !ok {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	var role admin_role.AdminRole
	err = admin_role.Dao().Query(gin.H{"id": r.RoleID}, []string{"*"}, nil).First(&role)
	if err != nil || role.ID == 0 {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	data := map[string]interface{}{
		"status":    *r.Status,
		"remark":    r.Remark,
		"username":  r.Username,
		"nick_name": r.NickName,
		"phone":     r.Phone,
		"email":     r.Email,
		//"password":  pw,
		"sex": *r.Sex,
		//"role_id": r.RoleID,
	}
	err = service.AdminUserService.Update(id, data)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, nil)
	return
}

type updateAdminUserStatus struct {
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

func (uc *userController) UpdateStatus(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	var r updateAdminUserStatus
	err = ctx.ShouldBind(&r)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}
	s := 0
	if *r.Status > 0 {
		s = 1
	}

	err = service.AdminUserService.UpdateStatus(id, s)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, nil)
	return
}

func (uc *userController) Reset(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	var r = admin_request.NewResetAdminUser()
	if data, ok := r.Validate(ctx); !ok {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	data := map[string]interface{}{
		"password": pw,
	}
	err = service.AdminUserService.Update(id, data)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, nil)
	return
}
