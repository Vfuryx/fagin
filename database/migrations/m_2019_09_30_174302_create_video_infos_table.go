package migrations

import (
	"fagin/pkg/migrate/migration"
	"gorm.io/gorm"
	"gopkg.in/gormigrate.v1"
	"time"
)

func init() {
	type VideoInfo struct {
		ID          uint64 `gorm:"primary_key"`
		CreatedAt   time.Time
		UpdatedAt   time.Time
		DeletedAt   *time.Time `sql:"index"`
		AuthorId    uint       `gorm:"index;not null;default:0;comment:'作者';column:author_id;"`
		Title       string     `gorm:"type:varchar(60);not null;default:'';comment:'标题';column:title;"`
		Path        string     `gorm:"type:varchar(255);not null;default:'';comment:'路径';column:path;"`
		Duration    string     `gorm:"type:varchar(32);not null;default:'';comment:'播放时长';column:duration;"`
		Description string     `gorm:"type:varchar(255);not null;default:'';comment:'视频描述';column:description;"`
		Status      uint8      `gorm:"not null;default:0;comment:'状态 0:关闭 1:开启';column:status;"`
	}
	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
		ID: "m_2019_09_30_174302_create_video_infos_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&VideoInfo{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTableIfExists(&VideoInfo{}).Error
		},
	})
}
