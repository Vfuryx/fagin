package api

import (
	"fagin/app"
	"fagin/app/constants/status"
	"fagin/app/errno"
	"fagin/app/models/video_info"
	"fagin/app/service"
	"fagin/config"
	"fagin/pkg/log"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
	"time"
)

type videoController struct{}

var VideoController videoController

// @Summary 播放视频接口
// @Description 播放视频
// @Tags video
// @Accept  mpfd
// @Produce  json
// @Param id path integer true "ID"
// @Success 200 {object} app.Response "成功返回视频溜 <br/> 失败返回 {"code": 10005,"message": "打开文件失败"}"
// @Router /api/video/play/{id} [get]
func (videoController) PlayVideo(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrBind, nil)
		return
	}

	params := map[string]interface{}{
		"id":     id,
		"status": status.Active,
	}
	columns := []string{"id", "path"}
	var v video_info.VideoInfo
	err = service.VideoInfo.Query(params, columns, nil).Find(&v)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrOpenFile, nil)
		return
	}

	path := config.App.StoragePath + v.Path
	file, err := os.Open(path)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrOpenFile, nil)
		return
	}
	http.ServeContent(ctx.Writer, ctx.Request, "", time.Now(), file)
	_ = file.Close()
	return
}
