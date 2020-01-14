package migrations

import (
	"fagin/pkg/migrate/migration"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)


////////////// 数据迁移start ////////////////////////////////////////////////////////////////////////////////////////////
/**
 * 创建users表
 */
func init() {
	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
		ID: "m_2018_01_23_00001_create_user",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				gorm.Model
				Username 	string		`gorm:"type:varchar(64);not null;default:'';comment:'用户名';column:username;"`
				Email 		string		`gorm:"type:varchar(64);not null;default:'';comment:'邮箱';column:email;"`
				Password	string		`gorm:"type:varchar(127);not null;default:'';comment:'密码';column:password;"`
				Avatar 		string		`gorm:"type:varchar(255);not null;default:'';comment:'头像';column:avatar;"`
				Status		uint8		`gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
			}
			return tx.AutoMigrate(&User{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTableIfExists("users").Error
		},
	})
}

////////////// 数据迁移end //////////////////////////////////////////////////////////////////////////////////////////////
