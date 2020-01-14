package web

import (
	"fagin/app"
	"fagin/app/constants/status"
	"fagin/app/models/video_info"
	"fagin/app/responses"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/log"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
	"net/http"
)

type videoController struct{}

var VideoController videoController

func (videoController) VideoList(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	// 查询
	params := map[string]interface{}{
		"status": status.Active,
	}
	// 字段
	columns := []string{"id", "title", "path", "status"}
	vs, err := service.VideoInfo.VideoList(params, columns, nil, &paginator)
	if err != nil {
		log.Log.Errorln(err)
		ctx.String(http.StatusNotFound, "404")
		return
	}

	list := responses.VideoInfoList(vs...).Collection()
	app.View(ctx, "web.video.list", gin.H{
		"title":     "视频列表",
		"list":      list,
		"paginator": paginator,
	})
}

func (videoController) VideoShow(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		log.Log.Errorln(err)
		ctx.String(http.StatusNotFound, "404")
		return
	}

	params := map[string]interface{}{
		"id":     id,
		"status": status.Active,
	}
	columns := []string{"id", "title", "path", "duration", "description"}
	var v video_info.VideoInfo
	err = service.VideoInfo.Query(params, columns, nil).Find(&v)
	if err != nil {
		log.Log.Errorln(err)
		ctx.String(http.StatusNotFound, "404")
		return
	}

	data := responses.VideoInfoShow(v).Item()
	app.View(ctx, "web.video.show", gin.H{
		"title": v.Title,
		"video": data,
	})
}
