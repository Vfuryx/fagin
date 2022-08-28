package api

import (
	"fagin/app/enums"
	"fagin/app/errno"
	"fagin/app/models/user"
	apiRequest "fagin/app/requests/api"
	apiResponses "fagin/app/responses/api"
	"fagin/app/services"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"fagin/utils"

	"github.com/gin-gonic/gin"
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
	r, msg := request.Validation[apiRequest.AddUserRequest](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	u := user.User{
		Username: r.UserName,
		Password: r.Password,
		Status:   enums.StatusActive.Get(),
	}

	if err := services.User.AddUser(&u).Error; err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{
		"user": u,
	})
}

func (ctr *userController) UserList(ctx *gin.Context) {
	// 分页管理器
	p := db.NewPaginatorWithCtx(ctx, 1, DefaultLimit)

	// 查询条件
	params := map[string]interface{}{
		"status": enums.StatusActive.Get(),
	}
	// 查询字段
	columns := []string{"id", "username", "status"}

	users, err := services.User.UserList(params, columns, nil, p)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	urs := apiResponses.NewUserResponse(users...).Collection()

	ctr.ResponseJSONOK(ctx, gin.H{
		"users":     urs,
		"paginator": p,
	})
}

func (ctr *userController) ShowUser(ctx *gin.Context) {
	i := ctx.Param("id")

	id, err := utils.StrToInt64(i)
	if err != nil || id <= 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id", "username", "status"}

	u, err := services.User.ShowUser(uint(id), columns)
	if err != nil || u.Status == enums.StatusDisable.Get() {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ru := apiResponses.NewUserResponse(*u).Item()
	ctr.ResponseJSONOK(ctx, ru)
}

func (ctr *userController) UpdateUser(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	r, msg := request.Validation[apiRequest.UpdateUserRequest](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	data := map[string]interface{}{}
	if r.Status != nil {
		data["status"] = *r.Status
	}

	if r.Password != nil {
		data["password"] = *r.Password
	}

	err = services.User.UpdateUser(id, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *userController) DestroyUser(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.User.DestroyUser(id)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}
