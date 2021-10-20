package admin

import (
	"fagin/app/errno"
	"fagin/app/models/tag"
	admin_request "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
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
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	params := gin.H{
		"orderBy": "id asc",
	}
	columns := []string{"id", "name", "status"}
	with := gin.H{}
	result, err := service.Tag.Index(params, columns, with, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.TagList(result...).Collection()

	ctr.ResponseJsonOK(ctx, gin.H{
		"tags":      data,
		"paginator": paginator,
	})
	return
}

// Show 展示
func (ctr *tagController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id", "name", "status"}
	data, err := service.Tag.Show(id, columns)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, gin.H{
		"id":     data.ID,
		"name":   data.Name,
		"status": data.Status,
	})
	return
}

// Store 创建
func (ctr *tagController) Store(ctx *gin.Context) {
	var r = admin_request.NewCreateTag()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	c := tag.Tag{
		Name:   r.Name,
		Status: *r.Status,
	}

	err := service.Tag.Create(&c)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Update 更新
func (ctr *tagController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = admin_request.NewCreateTag()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}
	data := map[string]interface{}{
		"name":   r.Name,
		"status": *r.Status,
	}
	err = service.Tag.Update(id, data)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Del 删除
func (ctr *tagController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.Tag.Delete(id)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Deletes 批量删除
func (ctr *tagController) Deletes(ctx *gin.Context) {
	type R struct {
		IDs []uint `form:"ids" json:"ids" binding:"required"`
	}
	var ids R
	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.Tag.Deletes(ids.IDs)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// All 所有
func (ctr *tagController) All(ctx *gin.Context) {
	params := gin.H{
		"status":  1,
		"orderBy": "id asc",
	}
	columns := []string{"id", "name", "status"}
	with := gin.H{}
	result, err := service.Tag.All(params, columns, with)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.TagList(result...).Collection()

	ctr.ResponseJsonOK(ctx, gin.H{
		"tags": data,
	})
	return
}
