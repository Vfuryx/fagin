package migrations
//
//import (
//	"fagin/pkg/migrate/migration"
//	"github.com/jinzhu/gorm"
//	"gopkg.in/gormigrate.v1"
//)
//
//func init() {
//	// 角色表
//	type AdminRole struct {
//		gorm.Model
//		Name   string `gorm:"type:varchar(128);not null;default:'';comment:'角色名称';column:name;"`
//		Key    string `gorm:"type:varchar(32);not null;default:'';comment:'角色关键字';column:key;"`
//		Remark string `gorm:"type:varchar(255);not null;default:'';comment:'角色备注';column:remark;"`
//		Sort   uint   `gorm:"not null;default:0;comment:'排序，数字越大越靠前';column:sort"`
//		Status uint8  `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
//	}
//	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
//		ID: "m_2020_05_28_214751_create_admin_role_table",
//		Migrate: func(tx *gorm.DB) error {
//			err := tx.AutoMigrate(&AdminRole{}).Error
//			if err != nil {
//				return err
//			}
//			err = tx.Create(&AdminRole{
//				Name:   "admin",
//				Key:    "admin",
//				Remark: "admin",
//				Sort:   0,
//				Status: 1,
//			}).Error
//			if err != nil {
//				return err
//			}
//			return nil
//		},
//		Rollback: func(tx *gorm.DB) error {
//			return tx.DropTableIfExists(&AdminRole{}).Error
//		},
//	})
//}
