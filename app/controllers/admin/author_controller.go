package admin

import (
	"fagin/app/errno"
	"fagin/app/models/author"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/services"
	"fagin/pkg/db"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type authorController struct {
	BaseController
}

var AuthorController authorController

func (ctr *authorController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, DefaultLimit)

	params := gin.H{
		"orderBy": "id asc",
	}
	columns := []string{"id", "name", "sort", "status"}

	authors, err := services.Author.Index(params, columns, nil, paginator)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.NewAuthorList(authors...).Collection()

	ctr.ResponseJSONOK(ctx, gin.H{
		"authors":   data,
		"paginator": paginator,
	})
}

func (ctr *authorController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id", "name", "status"}

	b, err := services.Author.Show(id, columns)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{
		"id":     b.ID,
		"name":   b.Name,
		"status": b.Status,
	})
}

func (ctr *authorController) Store(ctx *gin.Context) {
	r, msg := request.Validation[adminRequest.CreateAuthor](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	c := author.Author{
		Name:   r.Name,
		Status: *r.Status,
	}

	err := services.Author.Create(&c)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *authorController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	r, msg := request.Validation[adminRequest.CreateAuthor](ctx)
	if len(msg) > 0 {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, msg)
		return
	}

	err = services.Author.Update(id, map[string]interface{}{
		"Name":   r.Name,
		"Status": *r.Status,
	})
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *authorController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.Author.Delete(id)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

type AuthorIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

func (ctr *authorController) Deletes(ctx *gin.Context) {
	var ids AuthorIDs

	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.Author.Deletes(ids.IDs)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

func (ctr *authorController) All(ctx *gin.Context) {
	params := gin.H{
		"status":  1,
		"orderBy": "id asc",
	}
	columns := []string{"id", "name", "status"}

	authors, err := services.Author.All(params, columns, nil)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.NewAuthorList(authors...).Collection()

	ctr.ResponseJSONOK(ctx, gin.H{
		"authors": data,
	})
}
