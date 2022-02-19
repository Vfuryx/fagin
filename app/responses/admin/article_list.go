package responses

import (
	"fagin/app"
	"fagin/app/models/article"
	"fagin/pkg/response"
)

type articleList struct {
	ms []*article.Article

	response.Collect
}

func NewArticleList(models ...*article.Article) response.Response {
	res := articleList{ms: models}
	res.SetCollect(&res)

	return &res
}

func (res *articleList) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res.ms {
		tags := make([]map[string]interface{}, 0, response.DefCap)

		for _, t := range res.ms[i].Tags {
			m := map[string]interface{}{
				"id":   t.ID,
				"name": t.Name,
			}
			tags = append(tags, m)
		}

		m := map[string]interface{}{
			"id":          res.ms[i].ID,
			"title":       res.ms[i].Title,
			"author_id":   res.ms[i].AuthorID,
			"category_id": res.ms[i].CategoryID,
			"post_at":     app.TimeToStr(res.ms[i].PostAt),
			"status":      res.ms[i].Status,
			"author":      res.ms[i].Author.Name,
			"category":    res.ms[i].Category.Name,
			"tags":        tags,
		}
		sm = append(sm, m)
	}

	return sm
}
