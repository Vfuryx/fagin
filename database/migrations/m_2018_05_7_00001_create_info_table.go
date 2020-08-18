package migrations

//import (
//	"fagin/pkg/migrate/migration"
//	"github.com/jinzhu/gorm"
//	"gopkg.in/gormigrate.v1"
//)
//
//////////////// 数据迁移start ////////////////////////////////////////////////////////////////////////////////////////////
//type Info struct {
//	gorm.Model
//	UserId 	uint 	`gorm:"index;not null;default:0;comment:'用户id'"`
//	Email 	string	`gorm:"type:varchar(100);not null;default:'';comment:'邮箱地址';"`
//}
//
///**
// * 创建权限表
// */
//func init() {
//	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
//		ID: "m_2018_05_7_00001_create_info_table",
//		Migrate: func(tx *gorm.DB) error {
//			return tx.AutoMigrate(&Info{}).Error
//		},
//		Rollback: func(tx *gorm.DB) error {
//			return tx.DropTableIfExists(&Info{}).Error
//		},
//	})
//}
//
//
//////////////// 数据迁移end //////////////////////////////////////////////////////////////////////////////////////////////
