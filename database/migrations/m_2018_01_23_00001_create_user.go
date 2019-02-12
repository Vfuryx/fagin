package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
	"time"
)


////////////// 数据迁移start ////////////////////////////////////////////////////////////////////////////////////////////

/**
 * 创建users表
 */
var m_2018_01_23_00001_create_user  = &gormigrate.Migration{
	ID: "m_2018_01_23_00001_create_user",
	Migrate: func(tx *gorm.DB) error {
		type User struct {
			ID        	uint 		`gorm:"primary_key"`
			Name 		string		`gorm:"type: varchar(64); not null; column: name"`
			Avatar 		string		`gorm:"type: varchar(255); not null; default: ''; column: avatar"`
			Age 		uint8		`gorm:"type: tinyint unsigned; not null; default: 0; column: age"`
			CreatedAt 	time.Time
			UpdatedAt 	time.Time
			DeletedAt 	*time.Time 	`sql:"index"`
		}
		return tx.AutoMigrate(&User{}).Error
	},
	Rollback: func(tx *gorm.DB) error {
		return tx.DropTable("users").Error
	},
}

////////////// 数据迁移end //////////////////////////////////////////////////////////////////////////////////////////////
