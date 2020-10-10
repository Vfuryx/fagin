package web

import (
	"encoding/json"
	"fagin/app"
	"fagin/app/cache"
	"fagin/app/errno"
	"fagin/app/models/banner"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/log"
	"github.com/gin-gonic/gin"
)

type indexController struct{}

var IndexController indexController

func (indexController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 2)
	// 闭包保存环境
	cache.Banner.Content = func() ([]banner.Banner, error) {
		params := gin.H{
			"sort": "sort desc, id asc",
		}
		columns := []string{"id", "title", "banner", "path", "sort", "status"}
		// 数据库中获取数据
		return service.Banner.Index(params, columns, nil, &paginator)
	}

	data, err := cache.Banner.Get("")
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.Api.ErrBannerList, nil)
		return
	}
	var res []banner.Banner
	err = json.Unmarshal([]byte(data), &res)
	app.JsonResponse(ctx, errno.OK, res)
	return
}
