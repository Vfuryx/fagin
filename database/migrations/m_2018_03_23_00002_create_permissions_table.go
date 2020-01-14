package migrations

import (
	"fagin/pkg/migrate/migration"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

////////////// 数据迁移start ////////////////////////////////////////////////////////////////////////////////////////////

/**
 * 创建权限表
 */
func init() {
	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
		ID: "m_2018_03_23_00001_create_permission_table",
		Migrate: func(tx *gorm.DB) error {
			type Permissions struct {
				gorm.Model
				Name    string `gorm:"type:varchar(64);not null;default:'';comment:'名称';column:name;"`
				Path    string `gorm:"type:varchar(64);not null;default:'';comment:'路径';column:path;"`
				Action  string `gorm:"type:varchar(64);not null;default:'';comment:'方法';column:action;"`
				Comment string `gorm:"type:varchar(64);not null;default:'';comment:'注释';column:comment;"`
				Sort    uint64 `gorm:"type:int(4) unsigned;not null;default:0;comment:'排序 值越大越靠前';column:sort;"`
				Status  uint8  `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
			}
			return tx.AutoMigrate(&Permissions{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTableIfExists("permissions").Error
		},
	})
}


////////////// 数据迁移end //////////////////////////////////////////////////////////////////////////////////////////////
