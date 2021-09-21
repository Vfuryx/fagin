package admin_request

import (
	"fagin/pkg/request"
	"time"
)

//type myTime time.Time
//var _ json.Unmarshaler = &myTime{}
//// 无法使用 JSON 中设置的time_format解析时间
//// https://github.com/gin-gonic/gin/issues/1193
//func (mt *myTime) UnmarshalJSON(bs []byte) error {
//	var s string
//	err := json.Unmarshal(bs, &s)
//	if err != nil {
//		return err
//	}
//	t, err := time.ParseInLocation("2006-01-02 15:04:05", s, time.UTC)
//	if err != nil {
//		return err
//	}
//	*mt = myTime(t)
//	return nil
//}

type createArticle struct {
	Title              string    `form:"title" json:"title" binding:"required,max=128"`
	Content            string    `form:"content" json:"content" binding:"required"`
	Summary            string    `form:"summary" json:"summary" binding:"required"`
	AuthorID           uint      `form:"author_id" json:"author_id" binding:"required,min=1"`
	CategoryID         uint      `form:"category_id" json:"category_id" binding:"required,min=1"`
	Status             *uint8    `form:"status" json:"status" binding:"required,oneof=0 1"`
	PostAt             time.Time `form:"post_at" json:"post_at" binding:"required" time_format:"2006-01-02 15:04:05"`
	Tags               []uint    `form:"tags" json:"tags" binding:"required,dive,required"`
	request.Validation `binding:"-"`
}

func NewCreateArticle() *createArticle {
	r := new(createArticle)
	r.SetRequest(r)
	return r
}

func (*createArticle) Message() map[string]string {
	return map[string]string{}
}

func (*createArticle) Attributes() map[string]string {
	return map[string]string{
		"Title":      "标题",
		"Content":    "内容",
		"Summary":    "摘要",
		"AuthorID":   "作者ID",
		"CategoryID": "分类ID",
		"Status":     "状态",
		"PostAt":     "发布时间",
		"Tags":       "标签组",
	}
}
