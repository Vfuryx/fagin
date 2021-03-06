package api

import (
	"fagin/app/constants/status"
	"fagin/app/errno"
	"fagin/app/models/user"
	"fagin/app/requests/api"
	apiResponses "fagin/app/responses/api"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
	"strconv"
)

type userController struct {
	BaseController
}

var UserController userController

type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}

func (uc *userController) Register(ctx *gin.Context) {
	var v = api_request.NewAddUserRequest()

	if msg, ok := v.Validate(ctx); !ok {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, msg)
		return
	}

	u := user.User{
		Username: v.UserName,
		Password: v.Password,
		Status:   status.Active,
	}
	err := service.User.AddUser(&u).Error
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, gin.H{
		"user": u,
	})
}

func (uc *userController) UserList(ctx *gin.Context) {
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
		uc.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
		return
	}
	urs := apiResponses.UserResponse(users...).Collection()
	uc.ResponseJsonOK(ctx, gin.H{
		"users":     urs,
		"paginator": p,
	})
}

func (uc *userController) ShowUser(ctx *gin.Context) {
	i := ctx.Param("id")

	id, err := strconv.ParseUint(i, 10, 64)
	if err != nil || id <= 0 {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	columns := []string{"id", "username", "status"}
	u, err := service.User.ShowUser(uint(id), columns)
	if err != nil || u.Status == status.Disable {
		uc.ResponseJsonErrLog(ctx, errno.Serve.ShowErr, err, nil)
		return
	}
	ru := apiResponses.UserResponse(*u).Item()
	uc.ResponseJsonOK(ctx, ru)
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	var r = api_request.NewUpdateUserRequest()

	if msg, ok := r.Validate(ctx); !ok {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, msg)
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
		uc.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, nil)
}

func (uc *userController) DestroyUser(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.User.DestroyUser(id)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, nil)
}
