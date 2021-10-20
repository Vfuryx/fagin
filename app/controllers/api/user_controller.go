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

func (ctr *userController) Register(ctx *gin.Context) {
	var v = api_request.NewAddUserRequest()

	if msg, ok := v.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, msg)
		return
	}

	u := user.User{
		Username: v.UserName,
		Password: v.Password,
		Status:   enums.StatusActive.Get(),
	}
	err := service.User.AddUser(&u).Error
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, gin.H{
		"user": u,
	})
}

func (ctr *userController) UserList(ctx *gin.Context) {
	// 分页管理器
	p := db.NewPaginatorWithCtx(ctx, 1, 15)

	// 查询条件
	params := map[string]interface{}{
		"status": enums.StatusActive.Get(),
	}
	// 查询字段
	columns := []string{"id", "username", "status"}
	users, err := service.User.UserList(params, columns, nil, p)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}
	urs := apiResponses.UserResponse(users...).Collection()
	ctr.ResponseJsonOK(ctx, gin.H{
		"users":     urs,
		"paginator": p,
	})
}

func (ctr *userController) ShowUser(ctx *gin.Context) {
	i := ctx.Param("id")

	id, err := strconv.ParseUint(i, 10, 64)
	if err != nil || id <= 0 {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id", "username", "status"}
	u, err := service.User.ShowUser(uint(id), columns)
	if err != nil || u.Status == enums.StatusDisable.Get() {
		ctr.ResponseJsonErrLog(ctx, errno.CtxShowErr, err)
		return
	}
	ru := apiResponses.UserResponse(*u).Item()
	ctr.ResponseJsonOK(ctx, ru)
}

func (ctr *userController) UpdateUser(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = api_request.NewUpdateUserRequest()

	if msg, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, msg)
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
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
}

func (ctr *userController) DestroyUser(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.User.DestroyUser(id)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
}
