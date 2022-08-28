package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type TagToArticle struct {
		ID uint `gorm:"primarykey"`

		TagID     uint `gorm:"index;not null;default:0;comment:标签ID;column:tag_id"`
		ArticleID uint `gorm:"index;not null;default:0;comment:文章ID;column:article_id"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20201021230200_create_tag_to_article_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='标签文章表'"
			}
			return tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&TagToArticle{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&TagToArticle{})
		},
	})
}
