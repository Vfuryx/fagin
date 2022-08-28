package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type Category struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`

		Name   string `gorm:"type:varchar(32);not null;default:'';comment:名称;column:name"`
		Sort   uint   `gorm:"not null;default:100;comment:排序，数字越大越靠前;column:sort"`
		Status uint8  `gorm:"not null;default:1;comment:状态 0:隐藏 1显示;column:status"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20201021225344_create_category_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='分类表'"
			}
			return tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&Category{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&Category{})
		},
	})
}
