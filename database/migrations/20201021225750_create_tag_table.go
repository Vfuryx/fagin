package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type Tag struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`

		Name   string `gorm:"type:varchar(32);not null;default:'';comment:名称;column:name"`
		Status uint8  `gorm:"not null;default:1;comment:状态 0:隐藏 1显示;column:status"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20220826095604_20201021225750_create_tag_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='标签表'"
			}
			return tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&Tag{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&Tag{})
		},
	})
}
