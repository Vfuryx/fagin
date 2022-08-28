package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type AdminDepartment struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdatedAt time.Time

		ParentID uint   `gorm:"not null;default:0;comment:父id;column:parent_id;"`
		Name     string `gorm:"type:varchar(128);not null;default:'';comment:部门名称;column:name;"`
		Remark   string `gorm:"type:varchar(255);not null;default:'';comment:部门备注;column:remark;"`
		Sort     uint   `gorm:"not null;default:0;comment:排序，数字越大越靠前;column:sort"`
		Status   uint8  `gorm:"not null;default:1;comment:状态: 0=>关闭 1=>开启;column:status;"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20200927151610_create_admin_departments_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='后台部门表'"
			}
			err := tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&AdminDepartment{})
			if err != nil {
				return err
			}
			return tx.Create(&AdminDepartment{
				ParentID: 0,
				Name:     "总部",
				Remark:   "",
				Sort:     100,
				Status:   1,
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&AdminDepartment{})
		},
	})
}
