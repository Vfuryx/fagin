package service

import (
	"fagin/app/models/article"
	"fagin/app/models/category"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type articleService struct{}

var Article articleService

func (articleService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) ([]article.Article, error) {
	var art []article.Article
	err := article.Dao().Query(params, columns, with).Paginator(&art, p)
	return art, err
}

func (articleService) Show(id uint, columns []string) (*article.Article, error) {
	b := article.New()
	with := gin.H{
		"Author":   nil,
		"Category": nil,
		"Tags": func(db *gorm.DB) *gorm.DB {
			return db.Where("status = ?", 1)
		},
	}
	err := b.Dao().Query(gin.H{"id": id}, columns, with).First(&b)
	return b, err
}

func (articleService) Create(m *article.Article) error {
	return article.Dao().Create(m)
}

func (articleService) Update(id uint, data gin.H) error {
	return article.Dao().Update(id, data)
}

func (articleService) Delete(id uint) error {
	return article.Dao().Destroy(id)
}

func (articleService) Deletes(ids []uint) error {
	return article.Dao().Deletes(ids)
}

func (articleService) All(params gin.H, columns []string, with gin.H) ([]article.Article, error) {
	var ms []article.Article
	err := article.Dao().Query(params, columns, with).Find(&ms)
	return ms, err
}

func (articleService) ByCate(cateName string, columns []string, with gin.H, p *db.Paginator) (ms []article.Article, err error) {
	params := gin.H{
		"name":   cateName,
		"status": 1,
	}
	// 获取分类
	cate := category.New()
	err = cate.Dao().Query(params, []string{"id"}, nil).First(&cate)
	if err != nil {
		return nil, err
	}

	// 获取文章
	params = gin.H{
		"category_id": cate.ID,
		"status":      1,
		"orderBy":        "post_at desc, id asc",
	}
	err = article.Dao().Query(params, columns, with).Find(&ms)
	return ms, err
}

func (articleService) ByTag(cateName string, columns []string, with gin.H, p *db.Paginator) (ms []article.Article, err error) {
	var as []article.Article
	err = db.ORM().Table("tags as t").
		Select("a.id").
		Joins("left join tag_to_article as ta on ta.tag_id = t.id").
		Joins("left join articles as a on a.id = ta.article_id").
		Where("t.name = ?", cateName).
		Where("t.status", 1).
		Where("a.deleted_at is Null").
		Find(&as).Error
	if err != nil || len(as) < 1{
		return nil, gorm.ErrRecordNotFound
	}

	ids := make([]uint, 0, 15)
	for _, v := range as {
		ids = append(ids, v.ID)
	}

	// 获取文章
	params := gin.H{
		"id_in":  ids,
		"status": 1,
		"orderBy":   "post_at desc, id asc",
	}
	err = article.Dao().Query(params, columns, with).Paginator(&ms, p)
	return
}
