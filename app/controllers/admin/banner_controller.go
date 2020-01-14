package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/banner"
	"fagin/app/requests/admin"
	"fagin/app/responses/admin_response"
	"fagin/app/service"
	"fagin/config"
	"fagin/pkg/db"
	"fagin/pkg/log"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"path"
	"time"
)

type bannerController struct{}

var BannerController bannerController

func (bannerController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	params := gin.H{
		"sort": "sort desc, id asc",
	}
	columns := []string{"id", "title", "banner", "path", "sort", "status"}

	banners, err := service.Banner.Index(params, columns, nil, &paginator)
	if err != nil {
		app.JsonResponse(ctx, errno.ErrBannerList, nil)
		return
	}

	data := admin_response.BannerList(banners...).Collection()

	app.JsonResponse(ctx, errno.OK, gin.H{
		"banners":   data,
		"paginator": paginator,
	})
	return
}

func (bannerController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.ErrBind, nil)
		return
	}

	columns := []string{"id", "title", "banner", "path", "sort", "status"}
	b, err := service.Banner.Show(id, columns)
	if err != nil {
		app.JsonResponse(ctx, errno.ErrBanner, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, gin.H{
		"id":     b.ID,
		"title":  b.Title,
		"banner": b.Banner,
		"path":   b.Path,
		"sort":   b.Sort,
		"status": b.Status,
	})
	return
}

func (bannerController) Store(ctx *gin.Context) {
	var r admin_request.CreateBanner
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.ErrBind, data)
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
		app.JsonResponse(ctx, errno.ErrAddBanner, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

func (bannerController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.ErrBind, nil)
		return
	}

	var r admin_request.UpdateBanner
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.ErrBind, data)
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
		app.JsonResponse(ctx, errno.ErrUpdateBanner, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

// 上传视频
func (bannerController) Upload(ctx *gin.Context) {
	var r admin_request.UploadBanner
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponseWithStatus(ctx, http.StatusBadRequest, errno.ErrBind, data)
		return
	}

	filePath := "/web/banner/" + time.Now().Format("2006123") + app.RandString(20) + path.Ext(r.File.Filename)

	// 是否创建目录
	dir := path.Dir(config.App.PublicPath + filePath)
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	err = ctx.SaveUploadedFile(r.File, config.App.PublicPath+filePath)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrUploadFile, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, gin.H{"path": filePath})

	return
}
