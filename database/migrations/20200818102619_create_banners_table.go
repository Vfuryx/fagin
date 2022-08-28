package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type Banner struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`
		Title     string         `gorm:"type:varchar(32);not null;default:'';comment:标题;column:title"`
		Banner    string         `gorm:"type:varchar(255);not null;default:'';comment:轮播图;column:banner"`
		Path      string         `gorm:"type:varchar(255);not null;default:'';comment:路径;column:path"`
		Sort      uint           `gorm:"not null;default:100;comment:排序，数字越大越靠前;column:sort"`
		Status    uint8          `gorm:"not null;default:1;comment:状态 0:隐藏 1显示;column:status"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20200818102619_create_banners_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='Banner表'"
			}
			return tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&Banner{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&Banner{})
		},
	})
}
