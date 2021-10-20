package admin

import (
	"fagin/app/errno"
	"fagin/app/models/video_info"
	admin_request "fagin/app/requests/admin"
	adminResponses "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/config"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"fagin/pkg/response"
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
)

type videoController struct {
	BaseController
}

var VideoController videoController

// VideoList 视频列表
func (ctr *videoController) VideoList(ctx *gin.Context) {
	paginator := db.NewPaginatorWithCtx(ctx, 1, 15)

	params := make(map[string]interface{})
	columns := []string{"id", "title", "status", "path", "description", "duration", "created_at"}

	videos, err := service.VideoInfo.VideoList(params, columns, nil, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := adminResponses.VideoList(videos...).Collection()
	response.JsonOK(ctx, gin.H{
		"videos":    data,
		"paginator": paginator,
	})
	return
}

// CreateVideo 创建视频
func (ctr *videoController) CreateVideo(ctx *gin.Context) {
	var r = admin_request.NewCreateVideo()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
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
		ctr.ResponseJsonErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// UpdateVideo 更新视频
func (ctr *videoController) UpdateVideo(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = admin_request.NewUpdateVideo()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
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
		ctr.ResponseJsonErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
	return
}

// DeleteVideo 删除视频
func (ctr *videoController) DeleteVideo(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.VideoInfo.DeleteVideo(id)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
}

type VideoIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

// DeleteVideos 批量删除视频
func (ctr *videoController) DeleteVideos(ctx *gin.Context) {
	var ids VideoIDs
	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	err = service.VideoInfo.DeleteVideos(ids.IDs)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, nil)
}

// UploadVideo 上传视频
func (ctr *videoController) UploadVideo(ctx *gin.Context) {
	var r = admin_request.NewUploadVideo()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, data)
		return
	}

	upload := service.NewUploadService(config.App().PublicPath)
	path, err := upload.UploadFile("/upload/video/", r.File)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.ReqUploadFileErr, err)
		return
	}

	ctr.ResponseJsonOK(ctx, gin.H{"path": "/public/" + path})
}

// PlayVideo 播放视频
func (ctr *videoController) PlayVideo(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, nil)
		return
	}

	params := map[string]interface{}{"id": id}
	columns := []string{"id", "path"}
	var v video_info.VideoInfo
	err = service.VideoInfo.Query(params, columns, nil).Find(&v)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxOpenFileErr, err)
		return
	}

	path := config.App().StoragePath + v.Path
	file, err := os.Open(path)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxOpenFileErr, err)
		return
	}
	http.ServeContent(ctx.Writer, ctx.Request, "", time.Now(), file)
	_ = file.Close()
	return
}
