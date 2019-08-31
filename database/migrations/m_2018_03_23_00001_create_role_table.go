package migrations

import (
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

////////////// 数据迁移start ////////////////////////////////////////////////////////////////////////////////////////////

/**
 * 创建角色表
 */
func init() {
	migrations = append(migrations, &gormigrate.Migration{
		ID: "m_2018_03_23_00001_create_role_table",
		Migrate: func(tx *gorm.DB) error {
			type Role struct {
				gorm.Model
				Name   string `gorm:"type:varchar(64);not null;default:'';comment:'名称';column:name;"`
				Sort   uint64 `gorm:"type:int(4) unsigned;not null;default:0;comment:'排序 值越大越靠前';column:sort;"`
				Status uint8  `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
			}
			return tx.AutoMigrate(&Role{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTableIfExists("roles").Error
		},
	})
}

////////////// 数据迁移end //////////////////////////////////////////////////////////////////////////////////////////////
