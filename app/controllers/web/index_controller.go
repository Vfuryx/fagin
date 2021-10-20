package web

import (
	"encoding/json"
	"fagin/app"
	"fagin/app/caches"
	"fagin/app/errno"
	"fagin/app/models/article"
	"fagin/app/models/category"
	"fagin/app/models/tag"
	webRequest "fagin/app/requests/web"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/errorw"
	"fagin/pkg/request"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type indexController struct {
	BaseController
}

// IndexController 首页控制器
var IndexController indexController

// Home
// @Summary 首页
// @Tags web
// @Produce  html
// @Success 200
// @Router / [get]
func (ctr *indexController) Home(ctx *gin.Context) {
	var err error
	str := ctx.GetString("web_cate")
	var cs []category.Category
	if str != "" {
		if err = json.Unmarshal([]byte(str), &cs); err != nil {
			ctr.ResponseJsonErr(ctx, err, "")
			return
		}
	}

	// 获取标签
	str = ctx.GetString("web_tags")
	var tags []tag.Tag
	if str != "" {
		if err = json.Unmarshal([]byte(str), &tags); err != nil {
			ctr.ResponseJsonErr(ctx, errno.CtxShowErr, err.Error())
			return
		}
	}

	// 获取文章
	paginator := db.NewPaginatorWithCtx(ctx, 1, 10)
	homeArticle := caches.NewHomeArticles(func() ([]byte, error) {
		params := gin.H{
			"status":  1,
			"orderBy": "post_at desc, id asc",
		}
		columns := []string{"id", "title", "author_id", "category_id", "post_at", "status"}
		with := gin.H{
			"Author":   nil,
			"Category": nil,
			"Tags": func(db *gorm.DB) *gorm.DB {
				return db.Where("status = ?", 1)
			},
		}
		data, err := service.Article.Index(params, columns, with, paginator)
		if err != nil {
			return nil, errorw.UP(err)
		}
		return json.Marshal(data)
	})
	bs, err := homeArticle.Get(strconv.Itoa(paginator.CurrentPage))
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}
	var articles []article.Article
	err = json.Unmarshal(bs, &articles)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	app.View(ctx, "web.site.home", gin.H{
		"categories": cs,
		"articles":   articles,
		"paginator":  paginator,
		"tags":       tags,
	})
	return
}

// Category 类型
func (ctr *indexController) Category(ctx *gin.Context) {
	str := ctx.GetString("web_cate")
	var cs []category.Category
	err := json.Unmarshal([]byte(str), &cs)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.CtxListErr, err.Error())
		return
	}

	// 获取标签
	str = ctx.GetString("web_tags")
	var tags []tag.Tag
	err = json.Unmarshal([]byte(str), &tags)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.CtxListErr, err.Error())
		return
	}

	var r = webRequest.NewCategoryList()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.CtxListErr, data)
		return
	}

	columns := []string{"id", "title", "author_id", "category_id", "post_at", "status", "summary"}
	with := gin.H{
		"Author":   nil,
		"Category": nil,
		"Tags": func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", 1)
		},
	}
	paginator := db.NewPaginatorWithCtx(ctx, 1, 10)
	articles, err := service.Article.ByCate(r.Cate, columns, with, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	app.View(ctx, "web.site.home", gin.H{
		"categories": cs,
		"articles":   articles,
		"paginator":  paginator,
		"tags":       tags,
	})
}

// Tag 标签
func (ctr *indexController) Tag(ctx *gin.Context) {
	str := ctx.GetString("web_cate")
	var cs []category.Category
	err := json.Unmarshal([]byte(str), &cs)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.CtxShowErr, nil)
		return
	}

	// 获取标签
	str = ctx.GetString("web_tags")
	var tags []tag.Tag
	err = json.Unmarshal([]byte(str), &tags)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.CtxShowErr, nil)
		return
	}

	var r = webRequest.NewTagList()
	if data, ok := r.Validate(ctx); !ok {
		ctr.ResponseJsonErr(ctx, errno.CtxShowErr, data)
		return
	}

	columns := []string{"id", "title", "author_id", "category_id", "post_at", "status", "summary"}
	with := gin.H{
		"Author":   nil,
		"Category": nil,
		"Tags": func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", 1)
		},
	}
	paginator := db.NewPaginatorWithCtx(ctx, 1, 10)
	articles, err := service.Article.ByTag(r.Tag, columns, with, paginator)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	app.View(ctx, "web.site.home", gin.H{
		"categories": cs,
		"articles":   articles,
		"paginator":  paginator,
		"tags":       tags,
	})
}

// Article 文章
func (ctr *indexController) Article(ctx *gin.Context) {
	str := ctx.GetString("web_cate")
	var cs []category.Category
	err := json.Unmarshal([]byte(str), &cs)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.CtxListErr, err)
		return
	}

	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ctr.ResponseJsonErr(ctx, errno.ReqErr, err)
		return
	}

	webArticle := caches.NewWebArticle(func() ([]byte, error) {
		params := gin.H{
			"id":     id,
			"status": 1,
		}
		columns := []string{"id", "title", "author_id", "category_id", "post_at", "status", "content"}
		with := gin.H{
			"Author":   nil,
			"Category": nil,
			"Tags": func(db *gorm.DB) *gorm.DB {
				return db.Where("status = ?", 1)
			},
		}
		art := article.New()
		err = art.Dao().Query(params, columns, with).First(&art)
		if err != nil {
			return nil, err
		}
		return json.Marshal(art)
	})
	bs, err := webArticle.Get(strconv.FormatUint(uint64(id), 10))
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}
	var art article.Article
	err = json.Unmarshal(bs, &art)
	if err != nil {
		ctr.ResponseJsonErrLog(ctx, errno.CtxListErr, err)
		return
	}

	app.View(ctx, "web.site.article", gin.H{
		"categories": cs,
		"article":    art,
	})
	return
}

// Search 搜索
func (*indexController) Search(ctx *gin.Context) {

}
