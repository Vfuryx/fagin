package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type CompanyIntroduction struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`

		Name    string `gorm:"type:varchar(32);not null;default:'';comment:公司名称;column:name"`
		Image   string `gorm:"type:varchar(255);not null;default:'';comment:图片;column:image"`
		Content string `gorm:"type:text;not null;comment:内容;column:content"`
		Status  uint8  `gorm:"not null;default:1;comment:状态: 0=>关闭 1=>开启;column:status;"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20201021212532_create_company_introduction_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='公司简介'"
			}
			err := tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&CompanyIntroduction{})
			if err != nil {
				return err
			}

			return tx.Create(&CompanyIntroduction{
				Name:    "公司名称",
				Image:   "",
				Content: "",
				Status:  1,
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&CompanyIntroduction{})
		},
	})
}
