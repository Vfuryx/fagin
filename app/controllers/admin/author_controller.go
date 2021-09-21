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

func (ac *authorController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	params := gin.H{
		"orderBy": "id asc",
	}
	columns := []string{"id", "name", "sort", "status"}

	authors, err := service.Author.Index(params, columns, nil, paginator)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
		return
	}

	data := response.AuthorList(authors...).Collection()

	ac.ResponseJsonOK(ctx, gin.H{
		"authors":   data,
		"paginator": paginator,
	})
	return
}

func (ac *authorController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id", "name", "status"}
	b, err := service.Author.Show(id, columns)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.CtxShowErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, gin.H{
		"id":     b.ID,
		"name":   b.Name,
		"status": b.Status,
	})
	return
}

func (ac *authorController) Store(ctx *gin.Context) {
	var r = admin_request.NewCreateAuthor()
	if data, ok := r.Validate(ctx); !ok {
		ac.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	c := author.Author{
		Name:   r.Name,
		Status: *r.Status,
	}

	err := service.Author.Create(&c)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

func (ac *authorController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = admin_request.NewCreateAuthor()
	if data, ok := r.Validate(ctx); !ok {
		ac.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}
	data := map[string]interface{}{
		"Name":   r.Name,
		"Status": *r.Status,
	}
	err = service.Author.Update(id, data)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

func (ac *authorController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.Author.Delete(id)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

type AuthorIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

func (ac *authorController) Deletes(ctx *gin.Context) {
	var ids AuthorIDs
	err := ctx.ShouldBind(&ids)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.Author.Deletes(ids.IDs)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

func (ac *authorController) All(ctx *gin.Context) {
	params := gin.H{
		"status":  1,
		"orderBy": "id asc",
	}
	columns := []string{"id", "name", "status"}

	authors, err := service.Author.All(params, columns, nil)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.CtxListErr, err, nil)
		return
	}

	data := response.AuthorList(authors...).Collection()

	ac.ResponseJsonOK(ctx, gin.H{
		"authors": data,
	})
	return
}
