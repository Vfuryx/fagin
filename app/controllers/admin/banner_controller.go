package admin

import (
	"fagin/app/errno"
	"fagin/app/models/banner"
	"fagin/app/requests/admin"
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

func (bc *bannerController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	params := gin.H{
		"orderBy": "sort desc, id asc",
	}
	columns := []string{"id", "title", "banner", "path", "sort", "status"}

	banners, err := service.Banner.Index(params, columns, nil, &paginator)
	if err != nil {
		bc.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
		return
	}

	data := response.BannerList(banners...).Collection()

	bc.ResponseJsonOK(ctx, gin.H{
		"banners":   data,
		"paginator": paginator,
	})
	return
}

func (bc *bannerController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		bc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	columns := []string{"id", "title", "banner", "path", "sort", "status"}
	b, err := service.Banner.Show(id, columns)
	if err != nil {
		bc.ResponseJsonErrLog(ctx, errno.Serve.ShowErr, err, nil)
		return
	}

	bc.ResponseJsonOK(ctx, gin.H{
		"id":     b.ID,
		"title":  b.Title,
		"banner": b.Banner,
		"path":   b.Path,
		"sort":   b.Sort,
		"status": b.Status,
	})
	return
}

func (bc *bannerController) Store(ctx *gin.Context) {
	var r = admin_request.NewCreateBanner()
	if data, ok := r.Validate(ctx); !ok {
		bc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
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
		bc.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	bc.ResponseJsonOK(ctx, nil)
	return
}

func (bc *bannerController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		bc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	var r = admin_request.NewUpdateBanner()
	if data, ok := r.Validate(ctx); !ok {
		bc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
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
		bc.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	bc.ResponseJsonOK(ctx, nil)
	return
}

// Upload 上传视频
func (bc *bannerController) Upload(ctx *gin.Context) {
	var r = admin_request.NewUploadBanner()
	if data, ok := r.Validate(ctx); !ok {
		bc.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	upload := service.NewUploadService(config.App.PublicPath)
	path, err := upload.UploadFile("/web/banner/", r.File)
	if err != nil {
		bc.ResponseJsonErrLog(ctx, errno.Serve.UploadFileErr, err, nil)
		return
	}

	bc.ResponseJsonOK(ctx, gin.H{"path": "/public/" + path})
	return
}

func (bc *bannerController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		bc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.Banner.Delete(id)
	if err != nil {
		bc.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	bc.ResponseJsonOK(ctx, nil)
	return
}

type BannerIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

// DeleteBanners 批量删除轮播图
func (bc *bannerController) DeleteBanners(ctx *gin.Context) {
	var ids BannerIDs
	err := ctx.ShouldBind(&ids)
	if err != nil {
		bc.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.Banner.DeleteBanners(ids.IDs)
	if err != nil {
		bc.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	bc.ResponseJsonOK(ctx, nil)
}
