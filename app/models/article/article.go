package article

import (
	"fagin/app/models/author"
	"fagin/app/models/category"
	"fagin/app/models/tag"
	"gorm.io/gorm"
	"time"
)

type Article struct {
	ID         uint `gorm:"primarykey"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  gorm.DeletedAt    `gorm:"index"`
	Title      string            `gorm:"type:varchar(128);not null;default:'';comment:'标题';column:title"`
	Content    string            `gorm:"type:test;not null;comment:'内容';column:content"`
	Summary    string            `gorm:"type:test;not null;comment:'摘要';column:summary"`
	AuthorID   uint              `gorm:"index;not null;default:0;comment:'作者id';column:author_id"`
	CategoryID uint              `gorm:"index;not null;default:0;comment:'作者id';column:category_id"`
	Status     uint8             `gorm:"not null;default:1;comment:'状态 0:隐藏 1显示';column:status"`
	PostAt     time.Time         `sql:"index"`
	Author     author.Author     `gorm:"foreignKey:AuthorID;references:id"`
	Category   category.Category `gorm:"foreignKey:CategoryID;references:id"`
	Tags       []tag.Tag         `gorm:"many2many:tag_to_article;joinForeignKey:article_id;JoinReferences:tag_id"`
}
