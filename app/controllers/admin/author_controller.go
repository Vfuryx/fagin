package admin

import (
	"fagin/app/errno"
	"fagin/app/models/author"
	admin_request "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type authorController struct {
	BaseController
}

var AuthorController authorController

func (ctr *authorController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	params := gin.H{
		"orderBy": "id asc",
	}
	columns := []string{"id", "name", "sort", "status"}

	authors, err := service.Author.Index(params, columns, nil, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.AuthorList(authors...).Collection()

	ctr.ResponseJsonOK(ctx, gin.H{
		"authors":   data,
		"paginator": paginator,
	})
	return
}

func (ctr *authorController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id", "name", "status"}
	b, err := service.Author.Show(id, columns)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, gin.H{
		"id":     b.ID,
		"name":   b.Name,
		"status": b.Status,
	})
	return
}

func (ctr *authorController) Store(ctx *gin.Context) {
	var r = admin_request.NewCreateAuthor()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	c := author.Author{
		Name:   r.Name,
		Status: *r.Status,
	}

	err := service.Author.Create(&c)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

func (ctr *authorController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = admin_request.NewCreateAuthor()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}
	data := map[string]interface{}{
		"Name":   r.Name,
		"Status": *r.Status,
	}
	err = service.Author.Update(id, data)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

func (ctr *authorController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.Author.Delete(id)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

type AuthorIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

func (ctr *authorController) Deletes(ctx *gin.Context) {
	var ids AuthorIDs
	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.Author.Deletes(ids.IDs)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

func (ctr *authorController) All(ctx *gin.Context) {
	params := gin.H{
		"status":  1,
		"orderBy": "id asc",
	}
	columns := []string{"id", "name", "status"}

	authors, err := service.Author.All(params, columns, nil)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.AuthorList(authors...).Collection()

	ctr.ResponseJsonOK(ctx, gin.H{
		"authors": data,
	})
	return
}
