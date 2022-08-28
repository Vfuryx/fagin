package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type VideoInfo struct {
		ID          uint `gorm:"primarykey"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   gorm.DeletedAt `gorm:"index"`
		AuthorID    uint           `gorm:"index;not null;default:0;comment:作者;column:author_id;"`
		Title       string         `gorm:"type:varchar(60);not null;default:'';comment:标题;column:title;"`
		Path        string         `gorm:"type:varchar(255);not null;default:'';comment:路径;column:path;"`
		Duration    string         `gorm:"type:varchar(32);not null;default:'';comment:播放时长;column:duration;"`
		Description string         `gorm:"type:varchar(255);not null;default:'';comment:视频描述;column:description;"`
		Status      uint8          `gorm:"not null;default:0;comment:状态 0:关闭 1:开启;column:status;"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20200818101116_create_videos_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='视频资源表'"
			}
			return tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&VideoInfo{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&VideoInfo{})
		},
	})
}
