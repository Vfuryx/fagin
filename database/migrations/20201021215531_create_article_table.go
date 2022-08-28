package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type Article struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`

		Title      string    `gorm:"type:varchar(128);not null;default:'';comment:标题;column:title"`
		Content    string    `gorm:"type:text;not null;comment:内容;column:content"`
		Summary    string    `gorm:"type:text;not null;comment:摘要;column:summary"`
		AuthorID   uint      `gorm:"index;not null;default:0;comment:作者ID;column:author_id"`
		CategoryID uint      `gorm:"index;not null;default:0;comment:分类ID;column:category_id"`
		Status     uint8     `gorm:"index:idx_status_post_at,priority:1;not null;default:1;comment:状态 0:隐藏 1显示;column:status"`
		PostAt     time.Time `gorm:"index:idx_status_post_at,priority:2;comment:发布时间;column:post_at"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20201021215531_create_article_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='文章表'"
			}
			return tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&Article{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&Article{})
		},
	})
}
