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
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
	"strconv"
)

type indexController struct {
	BaseController
}

var IndexController indexController

// Home
// @Summary 首页
// @Tags web
// @Produce  html
// @Success 200
// @Router / [get]
func (ic *indexController) Home(ctx *gin.Context) {
	str := ctx.GetString("web_cate")
	var cs []category.Category
	err := json.Unmarshal([]byte(str), &cs)
	if err != nil {
		ic.ResponseJsonErr(ctx, err, "")
		return
	}
	// 获取标签
	str = ctx.GetString("web_tags")
	var tags []tag.Tag
	err = json.Unmarshal([]byte(str), &tags)
	if err != nil {
		ic.ResponseJsonErr(ctx, errno.Serve.ShowErr, err.Error())
		return
	}
	// 获取文章
	paginator := db.NewPaginator(ctx, 1, 10)
	homeArticle := caches.NewHomeArticles(func(key string) ([]byte, error) {
		params := gin.H{
			"status": 1,
			"orderBy":   "post_at desc, id asc",
		}
		columns := []string{"id", "title", "author_id", "category_id", "post_at", "status"}
		with := gin.H{
			"Author":   nil,
			"Category": nil,
			"Tags": func(db *gorm.DB) *gorm.DB {
				return db.Where("status = ?", 1)
			},
		}
		data, err := service.Article.Index(params, columns, with, &paginator)
		if err != nil {
			return nil, err
		}
		return json.Marshal(data)
	})
	str, err = homeArticle.Get(strconv.Itoa(paginator.CurrentPage))
	if err != nil {
		ic.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err.Error(), nil)
		return
	}
	var articles []article.Article
	err = json.Unmarshal([]byte(str), &articles)
	if err != nil {
		ic.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err.Error(), nil)
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
func (ic *indexController) Category(ctx *gin.Context) {
	str := ctx.GetString("web_cate")
	var cs []category.Category
	err := json.Unmarshal([]byte(str), &cs)
	if err != nil {
		ic.ResponseJsonErr(ctx, errno.Serve.ListErr, err.Error())
		return
	}

	// 获取标签
	str = ctx.GetString("web_tags")
	var tags []tag.Tag
	err = json.Unmarshal([]byte(str), &tags)
	if err != nil {
		ic.ResponseJsonErr(ctx, errno.Serve.ListErr, err.Error())
		return
	}

	var r = webRequest.NewCategoryList()
	if data, ok := r.Validate(ctx); !ok {
		ic.ResponseJsonErr(ctx, errno.Serve.ListErr, data)
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
	paginator := db.NewPaginator(ctx, 1, 10)
	articles, err := service.Article.ByCate(r.Cate, columns, with, &paginator)
	if err != nil {
		ic.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
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
func (ic *indexController) Tag(ctx *gin.Context) {
	str := ctx.GetString("web_cate")
	var cs []category.Category
	err := json.Unmarshal([]byte(str), &cs)
	if err != nil {
		ic.ResponseJsonErr(ctx, errno.Serve.ShowErr, nil)
		return
	}

	// 获取标签
	str = ctx.GetString("web_tags")
	var tags []tag.Tag
	err = json.Unmarshal([]byte(str), &tags)
	if err != nil {
		ic.ResponseJsonErr(ctx, errno.Serve.ShowErr, nil)
		return
	}

	var r = webRequest.NewTagList()
	if data, ok := r.Validate(ctx); !ok {
		ic.ResponseJsonErr(ctx, errno.Serve.ShowErr, data)
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
	paginator := db.NewPaginator(ctx, 1, 10)
	articles, err := service.Article.ByTag(r.Tag, columns, with, &paginator)
	if err != nil {
		ic.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
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
func (ic *indexController) Article(ctx *gin.Context) {
	str := ctx.GetString("web_cate")
	var cs []category.Category
	err := json.Unmarshal([]byte(str), &cs)
	if err != nil {
		ic.ResponseJsonErr(ctx, errno.Serve.ListErr, err)
		return
	}

	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ic.ResponseJsonErr(ctx, errno.Serve.BindErr, err)
		return
	}

	webArticle := caches.NewWebArticle(func(key string) ([]byte, error) {
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
	str, err = webArticle.Get(strconv.FormatUint(uint64(id), 10))
	if err != nil {
		ic.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
		return
	}
	var art article.Article
	err = json.Unmarshal([]byte(str), &art)
	if err != nil {
		ic.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
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
