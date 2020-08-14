package migrations

import (
	"fagin/pkg/migrate/migration"
	"gorm.io/gorm"
	"gopkg.in/gormigrate.v1"
	"time"
)

func init() {
	// 公司介绍表
	type CompanyIntroduction struct {
		ID        uint `gorm:"primary_key"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time `sql:"index"`
		Name      string     `gorm:"type:varchar(32);not null;default:'';comment:'公司名称';column:name"`
		Image     string     `gorm:"type:varchar(255);not null;default:'';comment:'图片';column:image"`
		Content   string     `gorm:"type:text;not null;comment:'内容';column:content"`
		Status    uint8      `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
	}
	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
		ID: "m_2020_01_10_173033_create_company_introduction_table",
		Migrate: func(tx *gorm.DB) error {
			err := tx.AutoMigrate(&CompanyIntroduction{}).Error
			if err != nil {
				return err
			}
			return tx.Create(&CompanyIntroduction{
				Name:    "公司名称",
				Image:   "",
				Status:  1,
				Content: "",
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTableIfExists(&CompanyIntroduction{}).Error
		},
	})
}
