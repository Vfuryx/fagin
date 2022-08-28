package tag_to_article

type TagToArticle struct {
	ID uint `gorm:"primarykey"`

	TagID     uint `gorm:"index;not null;default:0;comment:标签ID;column:tag_id"`
	ArticleID uint `gorm:"index;not null;default:0;comment:文章ID;column:article_id"`
}
