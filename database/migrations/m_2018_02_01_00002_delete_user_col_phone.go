package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

////////////// 数据迁移start ////////////////////////////////////////////////////////////////////////////////////////////

/**
 * 删除users表phone字段
 */
var m_2018_02_01_00002_delete_user_col_phone  = &gormigrate.Migration{
	ID: "m_2018_02_01_00002_delete_user_col_phone",
	Migrate: func(tx *gorm.DB) error {
		return tx.Table("users").DropColumn("phone").Error
	},
	Rollback: func(tx *gorm.DB) error {
		type User struct {
			Phone string
		}
		return tx.AutoMigrate(&User{}).Error
	},
}

////////////// 数据迁移end //////////////////////////////////////////////////////////////////////////////////////////////
