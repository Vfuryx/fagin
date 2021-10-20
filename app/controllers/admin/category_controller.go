package admin

import (
	"fagin/app/errno"
	"fagin/app/models/category"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type categoryController struct {
	BaseController
}

var CategoryController categoryController

func (ctr *categoryController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	params := gin.H{
		"orderBy": "sort desc, id asc",
	}
	columns := []string{"id", "name", "sort", "status"}

	categories, err := service.Category.Index(params, columns, nil, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.CategoryList(categories...).Collection()

	ctr.ResponseJsonOK(ctx, gin.H{
		"categories": data,
		"paginator":  paginator,
	})
	return
}

func (ctr *categoryController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id", "name", "sort", "status"}
	b, err := service.Category.Show(id, columns)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, gin.H{
		"id":     b.ID,
		"name":   b.Name,
		"sort":   b.Sort,
		"status": b.Status,
	})
	return
}

func (ctr *categoryController) Store(ctx *gin.Context) {
	var r = adminRequest.NewCreateCategory()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	c := category.Category{
		Name:   r.Name,
		Sort:   r.Sort,
		Status: *r.Status,
	}

	err := service.Category.Create(&c)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

func (ctr *categoryController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = adminRequest.NewCreateCategory()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}
	data := map[string]interface{}{
		"Name":   r.Name,
		"Sort":   r.Sort,
		"Status": *r.Status,
	}
	err = service.Category.Update(id, data)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

func (ctr *categoryController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.Category.Delete(id)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

type CategoryIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

func (ctr *categoryController) Deletes(ctx *gin.Context) {
	var ids CategoryIDs
	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.Category.Deletes(ids.IDs)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

func (ctr *categoryController) All(ctx *gin.Context) {
	params := gin.H{
		"status":  1,
		"orderBy": "sort desc, id asc",
	}
	columns := []string{"id", "name", "sort", "status"}

	categories, err := service.Category.All(params, columns, nil)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.CategoryList(categories...).Collection()

	ctr.ResponseJsonOK(ctx, gin.H{
		"categories": data,
	})
	return
}
