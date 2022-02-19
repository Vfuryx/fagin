package admin

import (
	"fagin/app/errno"
	"fagin/app/models/video_info"
	adminRequest "fagin/app/requests/admin"
	adminResponses "fagin/app/responses/admin"
	"fagin/app/services"
	"fagin/config"
	"fagin/pkg/db"
	"fagin/pkg/request"
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
	paginator := db.NewPaginatorWithCtx(ctx, 1, DefaultLimit)

	params := make(map[string]interface{})
	columns := []string{"id", "title", "status", "path", "description", "duration", "created_at"}

	videos, err := services.VideoInfo.VideoList(params, columns, nil, paginator)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxListErr, err)
		return
	}

	data := adminResponses.NewVideoList(videos...).Collection()

	ctr.ResponseJSONOK(ctx, gin.H{
		"videos":    data,
		"paginator": paginator,
	})
}

// CreateVideo 创建视频
func (ctr *videoController) CreateVideo(ctx *gin.Context) {
	var r = adminRequest.NewCreateVideo()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	v := video_info.VideoInfo{
		Title:       r.Title,
		Path:        r.Path,
		Description: r.Description,
		Status:      r.Status,
	}

	err := services.VideoInfo.CreateVideo(&v)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxStoreErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// UpdateVideo 更新视频
func (ctr *videoController) UpdateVideo(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	var r = adminRequest.NewUpdateVideo()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	data := map[string]interface{}{
		"title":       r.Title,
		"Path":        r.Path,
		"description": r.Description,
		"status":      *r.Status,
	}

	err = services.VideoInfo.UpdateVideo(id, data)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxUpdateErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// DeleteVideo 删除视频
func (ctr *videoController) DeleteVideo(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.VideoInfo.DeleteVideo(id)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

type VideoIDs struct {
	IDs []uint `form:"ids" json:"ids" binding:"required"`
}

// DeleteVideos 批量删除视频
func (ctr *videoController) DeleteVideos(ctx *gin.Context) {
	var ids VideoIDs

	err := ctx.ShouldBind(&ids)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	err = services.VideoInfo.DeleteVideos(ids.IDs)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxDeleteErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, nil)
}

// UploadVideo 上传视频
func (ctr *videoController) UploadVideo(ctx *gin.Context) {
	var r = adminRequest.NewUploadVideo()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, data)
		return
	}

	upload := services.NewUploadService(config.App().PublicPath())

	path, err := upload.UploadFile("/upload/video/", r.File)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.ReqUploadFileErr, err)
		return
	}

	ctr.ResponseJSONOK(ctx, gin.H{"path": "/public/" + path})
}

// PlayVideo 播放视频
func (ctr *videoController) PlayVideo(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	params := map[string]interface{}{"id": id}
	columns := []string{"id", "path"}

	var v video_info.VideoInfo

	err = services.VideoInfo.Query(params, columns, nil).Find(&v)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxOpenFileErr, err)
		return
	}

	path := config.App().StoragePath() + v.Path

	file, err := os.Open(path)
	if err != nil {
		ctr.ResponseJSONErrLog(ctx, errno.CtxOpenFileErr, err)
		return
	}

	http.ServeContent(ctx.Writer, ctx.Request, "", time.Now(), file)
	_ = file.Close()
}
