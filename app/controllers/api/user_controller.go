package api

import (
	"fagin/app"
	"fagin/app/constants/status"
	"fagin/app/errno"
	"fagin/app/models/user"
	"fagin/app/requests/api"
	"fagin/app/responses"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

type userController struct{}

var UserController userController

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (userController) Register(ctx *gin.Context) {
	var v api_request.AddUserRequest
	msg, ok := v.Validate(ctx)
	if !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, msg)
		return
	}

	u := user.User{
		Username: v.UserName,
		Password: v.Password,
		Status:   status.Active,
	}
	err := service.User.AddUser(&u).Error
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrAddUser, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, gin.H{
		"user": u,
	})
}

func (userController) UserList(ctx *gin.Context) {
	page := ctx.DefaultQuery("page", "1")
	per := ctx.DefaultQuery("per", "15")
	currentPage, err := strconv.Atoi(page)
	if err != nil {
		currentPage = 1
	}
	perPage, err := strconv.Atoi(per)
	if err != nil {
		perPage = 15
	}

	// 分页管理器
	p := db.Paginator{
		CurrentPage: currentPage,
		PageSize:    perPage,
	}
	// 查询条件
	params := map[string]interface{}{
		"status": status.Active,
	}
	// 查询字段
	columns := []string{"id", "username", "status"}
	users, err := service.User.UserList(params, columns, nil, &p)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrUserList, nil)
		return
	}
	urs := responses.UserResponse(users...).Collection()
	app.JsonResponse(ctx, errno.OK, gin.H{
		"users":     urs,
		"paginator": p,
	})
}

func (userController) ShowUser(ctx *gin.Context) {
	i := ctx.Param("id")

	id, err := strconv.ParseUint(i, 10, 64)
	if err != nil || id <= 0 {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	columns := []string{"id", "username", "status"}
	u, err := service.User.ShowUser(uint(id), columns)
	if err != nil || u.Status == status.Disable {
		app.JsonResponse(ctx, errno.Api.ErrUserNotFound, nil)
		return
	}
	ru := responses.UserResponse(*u).Item()
	app.JsonResponse(ctx, errno.OK, ru)
}

func (userController) UpdateUser(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	var r api_request.UpdateUserRequest
	msg, ok := r.Validate(ctx)
	if !ok {
		app.JsonResponse(ctx, errno.Api.ErrBind, msg)
		return
	}

	data := map[string]interface{}{}
	if r.Status != nil {
		data["status"] = *r.Status
	}
	if r.Password != nil {
		data["password"] = *r.Password
	}

	err = service.User.UpdateUser(id, data)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrUpdateUser, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
}

func (userController) DestroyUser(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	err = service.User.DestroyUser(id)
	if err != nil {
		app.JsonResponse(ctx, errno.Api.ErrDeleteUser, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
}
