package admin

import (
	"fagin/app/errno"
	"fagin/app/models/category"
	admin_request "fagin/app/requests/admin"
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

func (cc *categoryController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	params := gin.H{
		"orderBy": "sort desc, id asc",
	}
	columns := []string{"id", "name", "sort", "status"}

	categories, err := service.Category.Index(params, columns, nil, &paginator)
	if err != nil {
		cc.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
		return
	}

	data := response.CategoryList(categories...).Collection()

	cc.ResponseJsonOK(ctx, gin.H{
		"categories": data,
		"paginator":  paginator,
	})
	return
}

func (cc *categoryController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		cc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	columns := []string{"id", "name", "sort", "status"}
	b, err := service.Category.Show(id, columns)
	if err != nil {
		cc.ResponseJsonErrLog(ctx, errno.Serve.ShowErr, err, nil)
		return
	}

	cc.ResponseJsonOK(ctx, gin.H{
		"id":     b.ID,
		"name":   b.Name,
		"sort":   b.Sort,
		"status": b.Status,
	})
	return
}

func (cc *categoryController) Store(ctx *gin.Context) {
	var r = admin_request.NewCreateCategory()
	if data, ok := r.Validate(ctx); !ok {
		cc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	c := category.Category{
		Name:   r.Name,
		Sort:   r.Sort,
		Status: *r.Status,
	}

	err := service.Category.Create(&c)
	if err != nil {
		cc.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	cc.ResponseJsonOK(ctx, nil)
	return
}

func (cc *categoryController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		cc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	var r  = admin_request.NewCreateCategory()
	if data, ok := r.Validate(ctx); !ok {
		cc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}
	data := map[string]interface{}{
		"Name":   r.Name,
		"Sort":   r.Sort,
		"Status": *r.Status,
	}
	err = service.Category.Update(id, data)
	if err != nil {
		cc.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	cc.ResponseJsonOK(ctx, nil)
	return
}

func (cc *categoryController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		cc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.Category.Delete(id)
	if err != nil {
		cc.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	cc.ResponseJsonOK(ctx, nil)
	return
}

type CategoryIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

func (cc *categoryController) Deletes(ctx *gin.Context) {
	var ids CategoryIDs
	err := ctx.ShouldBind(&ids)
	if err != nil {
		cc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.Category.Deletes(ids.IDs)
	if err != nil {
		cc.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	cc.ResponseJsonOK(ctx, nil)
	return
}

func (cc *categoryController) All(ctx *gin.Context) {
	params := gin.H{
		"status": 1,
		"orderBy":   "sort desc, id asc",
	}
	columns := []string{"id", "name", "sort", "status"}

	categories, err := service.Category.All(params, columns, nil)
	if err != nil {
		cc.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
		return
	}

	data := response.CategoryList(categories...).Collection()

	cc.ResponseJsonOK(ctx, gin.H{
		"categories": data,
	})
	return
}
