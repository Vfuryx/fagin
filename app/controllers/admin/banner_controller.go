package admin

import (
	"fagin/app/errno"
	"fagin/app/models/banner"
	admin_request "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/config"
	"fagin/pkg/db"
	"fagin/pkg/request"

	"github.com/gin-gonic/gin"
)

type bannerController struct {
	BaseController
}

var BannerController bannerController

func (ctr *bannerController) Index(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	params := gin.H{
		"orderBy": "sort desc, id asc",
	}
	columns := []string{"id", "title", "banner", "path", "sort", "status"}

	banners, err := service.Banner.Index(params, columns, nil, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := response.BannerList(banners...).Collection()

	ctr.ResponseJsonOK(ctx, gin.H{
		"banners":   data,
		"paginator": paginator,
	})
	return
}

func (ctr *bannerController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	columns := []string{"id", "title", "banner", "path", "sort", "status"}
	b, err := service.Banner.Show(id, columns)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxShowErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, gin.H{
		"id":     b.ID,
		"title":  b.Title,
		"banner": b.Banner,
		"path":   b.Path,
		"sort":   b.Sort,
		"status": b.Status,
	})
	return
}

func (ctr *bannerController) Store(ctx *gin.Context) {
	var r = admin_request.NewCreateBanner()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	b := banner.Banner{
		Title:  r.Title,
		Banner: r.Banner,
		Path:   r.Path,
		Sort:   r.Sort,
		Status: *r.Status,
	}

	err := service.Banner.Create(&b)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

func (ctr *bannerController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = admin_request.NewUpdateBanner()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}
	data := map[string]interface{}{
		"Title":  r.Title,
		"Banner": r.Banner,
		"Path":   r.Path,
		"Sort":   r.Sort,
		"Status": *r.Status,
	}
	err = service.Banner.Update(id, data)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// Upload 上传视频
func (ctr *bannerController) Upload(ctx *gin.Context) {
	var r = admin_request.NewUploadBanner()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	upload := service.NewUploadService(config.App().PublicPath)
	path, err := upload.UploadFile("/web/banner/", r.File)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.ReqUploadFileErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, gin.H{"path": "/public/" + path})
	return
}

func (ctr *bannerController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.Banner.Delete(id)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

type BannerIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

// DeleteBanners 批量删除轮播图
func (ctr *bannerController) DeleteBanners(ctx *gin.Context) {
	var ids BannerIDs
	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.Banner.DeleteBanners(ids.IDs)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
}
