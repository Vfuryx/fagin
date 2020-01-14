package migrations

import (
	"fagin/app"
	"fagin/pkg/migrate/migration"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"time"
)

func init() {
	// 自定义表
	type AdminUser struct {
		ID        uint `gorm:"primary_key"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt *time.Time `sql:"index"`
		Username  string     `gorm:"type:varchar(64);not null;default:'';comment:'用户名';column:username"`
		Email     string     `gorm:"type:varchar(64);not null;default:'';comment:'邮箱';column:email;"`
		Password  string     `gorm:"type:varchar(127);not null;default:'';comment:'密码';column:password;"`
		Avatar    string     `gorm:"type:varchar(255);not null;default:'';comment:'头像';column:avatar;"`
		Status    uint8      `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
	}
	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
		ID: "m_2020_01_02_223802_create_admin_users_table",
		Migrate: func(tx *gorm.DB) error {
			pwd, _ := app.Encrypt("12345678")
			err := tx.AutoMigrate(&AdminUser{}).Error
			if err != nil {
				return err
			}
			return tx.Create(&AdminUser{
				Username: "admin",
				Password: pwd,
				Avatar:   "http://qiniu.furyx.cn/photo.jpg",
				Status:   1,
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTableIfExists(&AdminUser{}).Error
		},
	})
}
