package api

import (
	"fagin/app/enums"
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
		uc.ResponseJsonErr(ctx, errno.ReqErr, msg)
		return
	}

	u := user.User{
		Username: v.UserName,
		Password: v.Password,
		Status:   enums.StatusActive.Get(),
	}
	err := service.User.AddUser(&u).Error
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, gin.H{
		"user": u,
	})
}

func (uc *userController) UserList(ctx *gin.Context) {
	// 分页管理器
	p := db.NewPaginatorWithCtx(ctx, 1, 15)

	// 查询条件
	params := map[string]interface{}{
		"status": enums.StatusActive,
	}
	// 查询字段
	columns := []string{"id", "username", "status"}
	users, err := service.User.UserList(params, columns, nil, p)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
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
		uc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id", "username", "status"}
	u, err := service.User.ShowUser(uint(id), columns)
	if err != nil || u.Status == enums.StatusDisable.Get() {
		uc.ResponseJsonErrLog(ctx, errno.CtxShowErr, err, nil)
		return
	}
	ru := apiResponses.UserResponse(*u).Item()
	uc.ResponseJsonOK(ctx, ru)
}

func (uc *userController) UpdateUser(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = api_request.NewUpdateUserRequest()

	if msg, ok := r.Validate(ctx); !ok {
		uc.ResponseJsonErr(ctx, errno.ReqErr, msg)
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
		uc.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, nil)
}

func (uc *userController) DestroyUser(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		uc.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.User.DestroyUser(id)
	if err != nil {
		uc.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err, nil)
		return
	}

	uc.ResponseJsonOK(ctx, nil)
}
