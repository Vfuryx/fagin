package admin

import (
	"fagin/app/errno"
	"fagin/app/models/article"
	"fagin/app/models/author"
	"fagin/app/models/category"
	"fagin/app/models/tag"
	adminRequest "fagin/app/requests/admin"
	response "fagin/app/responses/admin"
	"fagin/app/service"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type articleController struct {
	BaseController
}

var ArticleController articleController

func (ac *articleController) Index(ctx *gin.Context) {
	paginator := db.NewPaginator(ctx, 1, 15)

	params := gin.H{
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
	articles, err := service.Article.Index(params, columns, with, &paginator)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.ListErr, err, nil)
		return
	}

	data := response.ArticleList(articles...).Collection()

	ac.ResponseJsonOK(ctx, gin.H{
		"articles":  data,
		"paginator": paginator,
	})
	return
}

func (ac *articleController) Show(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	columns := []string{"id", "title", "author_id", "category_id", "post_at", "status", "content", "summary"}
	data, err := service.Article.Show(id, columns)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.ShowErr, err, nil)
		return
	}
	tags := make([]map[string]interface{}, 0, 20)
	for _, t := range data.Tags {
		m := map[string]interface{}{
			"id":   t.ID,
			"name": t.Name,
		}
		tags = append(tags, m)
	}
	ac.ResponseJsonOK(ctx, gin.H{
		"id":          data.ID,
		"title":       data.Title,
		"content":     data.Content,
		"summary":     data.Summary,
		"author_id":   data.AuthorID,
		"category_id": data.CategoryID,
		"post_at":     data.PostAt,
		"status":      data.Status,
		"author":      data.Author.Name,
		"category":    data.Category.Name,
		"tags":        tags,
	})
	return
}

func (ac *articleController) Store(ctx *gin.Context) {
	var r = adminRequest.NewCreateArticle()
	if data, ok := r.Validate(ctx); !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	// 判断作者是否存在
	a := author.New()
	err := a.Dao().Query(gin.H{"id": r.AuthorID}, []string{"1"}, nil).First(&a)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	// 判断分类是否存在
	c := category.New()
	err = c.Dao().Query(gin.H{"id": r.CategoryID}, []string{"1"}, nil).First(&c)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	// 获取标签
	var tags []tag.Tag
	err = tag.Dao().Query(gin.H{"in_id": r.Tags, "status": 1}, []string{"id"}, nil).Find(&tags)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	b := article.Article{
		Title:      r.Title,
		Content:    r.Content,
		Summary:    r.Summary,
		AuthorID:   r.AuthorID,
		CategoryID: r.CategoryID,
		Status:     *r.Status,
		PostAt:     r.PostAt,
		Tags:       tags,
	}

	err = service.Article.Create(&b)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.StoreErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

func (ac *articleController) Update(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	var r = adminRequest.NewCreateArticle()
	if data, ok := r.Validate(ctx); !ok {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, data)
		return
	}

	// 判断作者是否存在
	a := author.New()
	err = a.Dao().Query(gin.H{"id": r.AuthorID}, []string{"1"}, nil).First(&a)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	// 判断分类是否存在
	c := category.New()
	err = c.Dao().Query(gin.H{"id": r.CategoryID}, []string{"1"}, nil).First(&c)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	// 获取标签
	var tags []tag.Tag
	err = tag.Dao().Query(gin.H{"in_id": r.Tags, "status": 1}, []string{"*"}, nil).Find(&tags)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	data := map[string]interface{}{
		"Title":      r.Title,
		"Content":    r.Content,
		"Summary":    r.Summary,
		"AuthorID":   r.AuthorID,
		"CategoryID": r.CategoryID,
		"Status":     r.Status,
		"PostAt":     r.PostAt,
		"Tags":       tags,
	}
	err = service.Article.Update(id, data)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.UpdateErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

func (ac *articleController) Del(ctx *gin.Context) {
	id, err := request.ShouldBindUriUintID(ctx)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.Article.Delete(id)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}

func (ac *articleController) Deletes(ctx *gin.Context) {
	type R struct {
		IDs []uint `form:"ids" json:"ids" binding:"required"`
	}
	var ids R
	err := ctx.ShouldBind(&ids)
	if err != nil {
		ac.ResponseJsonErr(ctx, errno.Serve.BindErr, nil)
		return
	}

	err = service.Article.Deletes(ids.IDs)
	if err != nil {
		ac.ResponseJsonErrLog(ctx, errno.Serve.DeleteErr, err, nil)
		return
	}

	ac.ResponseJsonOK(ctx, nil)
	return
}
