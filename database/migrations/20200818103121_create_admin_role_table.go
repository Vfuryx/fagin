package migrations

import (
	"fadmin/config"
	"fadmin/pkg/migrate"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

func init() {
	type AdminRole struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt uint
		UpdatedAt uint
		DeletedAt soft_delete.DeletedAt `gorm:"index"`

		Type   uint8  `gorm:"uniqueIndex:udx_type_key,priority:1;not null;default:0;comment:菜单类型 0:总后台 1:商家后台 2:集团后台;column:type;"`
		Name   string `gorm:"type:varchar(128);not null;default:'';comment:角色名称;column:name;"`
		Key    string `gorm:"type:varchar(45);uniqueIndex:udx_type_key,priority:2;not null;default:'';comment:角色关键字;column:key;"`
		Remark string `gorm:"type:varchar(255);not null;default:'';comment:角色备注;column:remark;"`
		Sort   uint   `gorm:"not null;default:0;comment:排序，数字越大越靠前;column:sort"`
		Status uint8  `gorm:"not null;default:1;comment:状态: 0=>关闭 1=>开启;column:status;"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20200818103121_create_admin_role_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='后台角色表'"
			}

			err := tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&AdminRole{})
			if err != nil {
				return err
			}

			return tx.Create(&AdminRole{
				Type:   0,
				Name:   "admin",
				Key:    "admin",
				Remark: "admin",
				Sort:   100,
				Status: 1,
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&AdminRole{})
		},
	})
}
