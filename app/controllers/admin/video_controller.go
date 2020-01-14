package admin

import (
	"fagin/app"
	"fagin/app/errno"
	"fagin/app/models/video_info"
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

type videoController struct{}

var VideoController videoController

// 视频列表
func (videoController) VideoList(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	params := make(map[string]interface{})
	columns := []string{"id", "title", "status", "path", "description"}

	videos, err := service.VideoInfo.VideoList(params, columns, nil, &paginator)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.InternalServerError, nil)
		return
	}

	data := admin_response.VideoList(videos...).Collection()

	app.JsonResponse(ctx, errno.OK, gin.H{
		"videos":    data,
		"paginator": paginator,
	})
	return
}

// 创建视频
func (videoController) CreateVideo(ctx *gin.Context) {
	var r admin_request.CreateVideo
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.ErrBind, data)
		return
	}
	v := video_info.VideoInfo{
		Title:       r.Title,
		Path:        r.Path,
		Description: r.Description,
		Status:      r.Status,
	}
	err := service.VideoInfo.CreateVideo(&v)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrAddVideo, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

// 更新视频
func (videoController) UpdateVideo(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.ErrBind, nil)
		return
	}

	var r admin_request.UpdateVideo
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.ErrBind, data)
		return
	}

	data := map[string]interface{}{
		"title":       r.Title,
		"Path":        r.Path,
		"description": r.Description,
		"status":      *r.Status,
	}

	err = service.VideoInfo.UpdateVideo(id, data)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrUpdateVideo, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, nil)
	return
}

// 删除视频
func (videoController) DeleteVideo(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		app.JsonResponse(ctx, errno.ErrBind, nil)
		return
	}

	err = service.VideoInfo.DeleteVideo(id)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrDeleteVideo, nil)
	}

	app.JsonResponse(ctx, errno.OK, nil)
}

// 上传视频
func (videoController) UploadVideo(ctx *gin.Context) {
	var r admin_request.UploadVideo
	if data, ok := r.Validate(ctx); !ok {
		app.JsonResponse(ctx, errno.ErrBind, data)
		return
	}

	filePath := "/video/" + time.Now().Format("2006123") + app.RandString(20) + path.Ext(r.File.Filename)

	// 是否创建目录
	dir := path.Dir(config.App.PublicPath + filePath)
	_, err := os.Stat(dir)
	if err != nil && os.IsNotExist(err) {
		log.Log.Errorln(err)
		_ = os.MkdirAll(dir, os.ModePerm)
	}

	err = ctx.SaveUploadedFile(r.File, config.App.StoragePath+filePath)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrUploadFile, nil)
		return
	}

	app.JsonResponse(ctx, errno.OK, gin.H{"path": filePath})
	return
}

// 播放视频
func (videoController) PlayVideo(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrBind, nil)
		return
	}

	params := map[string]interface{}{"id": id}
	columns := []string{"id", "path"}
	var v video_info.VideoInfo
	err = service.VideoInfo.Query(params, columns, nil).Find(&v)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrOpenFile, nil)
		return
	}

	path := config.App.StoragePath + v.Path
	file, err := os.Open(path)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrOpenFile, nil)
		return
	}
	http.ServeContent(ctx.Writer, ctx.Request, "", time.Now(), file)
	_ = file.Close()
	return
}
