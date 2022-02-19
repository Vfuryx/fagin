package api

import (
	"fagin/app/enums"
	"fagin/app/errno"
	"fagin/app/models/video_info"
	"fagin/app/services"
	"fagin/config"
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

// PlayVideo
// @Summary 播放视频接口
// @Description 播放视频
// @Tags video
// @Accept  mpfd
// @Produce  json
// @Param id path integer true "ID"
// @Success 200 {object} response.Response "成功返回视频溜 <br/> 失败返回 {"code": 10005,"message": "打开文件失败"}"
// @Router /api/video/play/{id} [get]
func (ctr *videoController) PlayVideo(ctx *gin.Context) {
	id, err := request.ShouldBindURIUintID(ctx)
	if err != nil {
		ctr.ResponseJSONErr(ctx, errno.ReqErr, nil)
		return
	}

	params := map[string]interface{}{
		"id":     id,
		"status": enums.StatusActive.Get(),
	}
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
