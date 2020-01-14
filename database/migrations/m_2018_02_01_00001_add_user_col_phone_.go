package migrations

import (
	"fagin/pkg/migrate/migration"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)


////////////// 数据迁移start ////////////////////////////////////////////////////////////////////////////////////////////

/**
 * users表新增phone字段
 */
func init() {
	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
		ID: "m_2018_02_01_00001_add_user_col_phone_",
		Migrate: func(tx *gorm.DB) error {
			type User struct {
				Phone string `gorm:"type:varchar(20);not null;default:'';comment:'手机';column:phone;"`
			}
			return tx.AutoMigrate(&User{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Table("users").DropColumn("phone").Error
			//return tx.Table("users").DropColumn("phone").Error
		},
	})
}

////////////// 数据迁移end //////////////////////////////////////////////////////////////////////////////////////////////

