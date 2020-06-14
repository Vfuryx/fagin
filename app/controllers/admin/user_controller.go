package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/admin_role"
	"fagin/app/models/admin_user"
	"fagin/app/requests/admin"
	"fagin/app/responses/admin_response"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/log"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
)

type userController struct{}

var UserController userController

func (userController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	var r admin_request.AdminUserList
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
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
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrUserList, nil)
		return
	}

	data := admin_response.AdminUserList(users...).Collection()

	app.JsonResponse(ctx, errno.OK, gin.H{
		"users":     data,
		"paginator": paginator,
	})
	return
}

func (userController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	columns := []string{"*"}
	user, err := service.AdminUserService.Show(id, columns)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrUserShow, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, gin.H{
		"id":        user.ID,
		"username":  user.Username,
		"nick_name": user.NickName,
		"email":     user.Email,
		"phone":     user.Phone,
		"avatar":    user.Avatar,
		"sex":       user.Sex,
		"remark":    user.Remark,
		"status":    user.Status,
		"role_id":   user.RoleID,
		"create_at": user.CreatedAt,
	})
	return
}

func (userController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	err = service.AdminUserService.Delete(id)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrDeleteUser, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

type userIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

func (userController) Dels(ctx *gin.Context) {
	var ids userIDs
	err := ctx.ShouldBind(&ids)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	err = service.AdminUserService.Deletes(ids.IDs)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrDeleteUser, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

func (userController) Store(ctx *gin.Context) {
	var r admin_request.CreateAdminUser
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}

	var role admin_role.AdminRole
	err := admin_role.Dao().Query(gin.H{"id": r.RoleID}, []string{"*"}, nil).First(&role)
	if err != nil || role.ID == 0 {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
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
		RoleID:   r.RoleID,
	}

	err = service.AdminUserService.Create(&user)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrAddUser, err)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

func (userController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	var r admin_request.UpdateAdminUser
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}

	var role admin_role.AdminRole
	err = admin_role.Dao().Query(gin.H{"id": r.RoleID}, []string{"*"}, nil).First(&role)
	if err != nil || role.ID == 0 {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	//pw, err := app.Encrypt(r.Password)
	//if err != nil {
	//	app.JsonResponse(ctx, errno.Api.ErrBind, nil)
	//	return
	//}

	data := map[string]interface{}{
		"status":    *r.Status,
		"remark":    r.Remark,
		"username":  r.Username,
		"nick_name": r.NickName,
		"phone":     r.Phone,
		"email":     r.Email,
		//"password":  pw,
		"sex":     *r.Sex,
		"role_id": r.RoleID,
	}
	err = service.AdminUserService.Update(id, data)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrUpdateUser, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

type updateAdminUserStatus struct {
	Status *uint8 `form:"status" json:"status" binding:"required,oneof=0 1"`
}

func (userController) UpdateStatus(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	var r updateAdminUserStatus
	err = ctx.ShouldBind(&r)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}
	s := 0
	if *r.Status > 0 {
		s = 1
	}

	err = service.AdminUserService.UpdateStatus(id, s)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrUpdateUser, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

func (userController) Reset(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	var r admin_request.ResetAdminUser
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, data)
		return
	}

	pw, err := app.Encrypt(r.Password)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	data := map[string]interface{}{
		"password": pw,
	}
	err = service.AdminUserService.Update(id, data)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrUpdateUser, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}
