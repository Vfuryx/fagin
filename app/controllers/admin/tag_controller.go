package admin

import (
	"fagin/app/errno"
	"fagin/app/models/tag"
	admin_request "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/services"
	"fagin/pkg/db"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type tagController struct {
	BaseController
}

var TagController tagController

// Index 列表
func (ctr *tagController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, DefaultLimit)

	params := gin.H{
		"orderBy": "id asc",
	}
	columns := []string{"id", "name", "status"}
	with := gin.H{}

	result, err := services.Tag.Index(params, columns, with, paginator)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.NewTagList(result...).Collection()

	ctr.ResponseJSONOK(ctx, gin.H{
		"tags":      data,
		"paginator": paginator,
	})
}

// Show 展示
func (ctr *tagController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id", "name", "status"}

	data, err := services.Tag.Show(id, columns)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{
		"id":     data.ID,
		"name":   data.Name,
		"status": data.Status,
	})
}

// Store 创建
func (ctr *tagController) Store(ctx *gin.Context) {
	var r = admin_request.NewCreateTag()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	c := tag.Tag{
		Name:   r.Name,
		Status: *r.Status,
	}

	err := services.Tag.Create(&c)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Update 更新
func (ctr *tagController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = admin_request.NewCreateTag()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	data := map[string]interface{}{
		"name":   r.Name,
		"status": *r.Status,
	}

	err = services.Tag.Update(id, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Del 删除
func (ctr *tagController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.Tag.Delete(id)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// Deletes 批量删除
func (ctr *tagController) Deletes(ctx *gin.Context) {
	type R struct {
		IDs []uint `form:"ids" json:"ids" binding:"required"`
	}

	var ids R

	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.Tag.Deletes(ids.IDs)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// All 所有
func (ctr *tagController) All(ctx *gin.Context) {
	params := gin.H{
		"status":  1,
		"orderBy": "id asc",
	}
	columns := []string{"id", "name", "status"}
	with := gin.H{}

	result, err := services.Tag.All(params, columns, with)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.NewTagList(result...).Collection()

	ctr.ResponseJSONOK(ctx, gin.H{
		"tags": data,
	})
}
