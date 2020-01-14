package web

import (
	"fagin/app"
	"fagin/app/cache"
	"fagin/app/errno"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/log"
	"github.com/gin-gonic/gin"
	"time"
)

type indexController struct{}

var IndexController indexController

func (indexController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 2)
	params := gin.H{
		"sort": "sort desc, id asc",
	}
	columns := []string{"id", "title", "banner", "path", "sort", "status"}

	// 获取缓存
	data, err := cache.Banner.GetList()
	if err == nil {
		app.JsonResponse(ctx, errno.OK, data)
		return
	} else {
		log.Log.Errorln(err)
	}

	// 数据库中获取数据
	banners, err := service.Banner.Index(params, columns, nil, &paginator)
	if err != nil {
		log.Log.Errorln(err)
		app.JsonResponse(ctx, errno.ErrBannerList, nil)
		return
	}
	// 设置缓存
	_, err = cache.Banner.SetList(banners, 10*time.Minute)
	if err != nil {
		log.Log.Errorln(err)
	}

	app.JsonResponse(ctx, errno.OK, banners)
}
