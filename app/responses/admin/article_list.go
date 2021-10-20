package admin_responses

import (
	"fagin/app"
	"fagin/app/models/article"
	"fagin/pkg/response"
)

type articleList struct {
	Ms []article.Article
	response.Collect
}

var _ response.Response = &articleList{}

func ArticleList(models ...article.Article) *articleList {
	res := articleList{Ms: models}
	res.SetCollect(&res)
	return &res
}

func (res *articleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
		tags := make([]map[string]interface{}, 0, 20)
		for _, t := range model.Tags {
			m := map[string]interface{}{
				"id":   t.ID,
				"name": t.Name,
			}
			tags = append(tags, m)
		}
		m := map[string]interface{}{
			"id":          model.ID,
			"title":       model.Title,
			"author_id":   model.AuthorID,
			"category_id": model.CategoryID,
			"post_at":     app.TimeToStr(model.PostAt),
			"status":      model.Status,
			"author":      model.Author.Name,
			"category":    model.Category.Name,
			"tags":        tags,
		}
		sm = append(sm, m)
	}
	return sm
}
