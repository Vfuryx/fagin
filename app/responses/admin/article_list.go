package responses

import (
	"fagin/app"
	"fagin/app/models/article"
	"fagin/pkg/response"
)

type articleList []article.Article

func NewArticleList(models ...article.Article) *response.Collect[articleList] {
	return new(response.Collect[articleList]).SetCollect(models)
}

func (res articleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		tags := make([]map[string]interface{}, 0, response.DefCap)

		for _, t := range res[i].Tags {
			m := map[string]interface{}{
				"id":   t.ID,
				"name": t.Name,
			}
			tags = append(tags, m)
		}

		m := map[string]interface{}{
			"id":          res[i].ID,
			"title":       res[i].Title,
			"author_id":   res[i].AuthorID,
			"category_id": res[i].CategoryID,
			"post_at":     app.TimeToStr(res[i].PostAt),
			"status":      res[i].Status,
			"author":      res[i].Author.Name,
			"category":    res[i].Category.Name,
			"tags":        tags,
		}
		sm = append(sm, m)
	}

	return sm
}
