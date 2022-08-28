package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type AdminUserRole struct {
		ID uint `gorm:"primarykey"`

		AdminID uint `gorm:"uniqueIndex:udx_admin_id_role_id,priority:1;not null;default:0;comment:管理员id;column:admin_id;"`
		RoleID  uint `gorm:"uniqueIndex:udx_admin_id_role_id,priority:2;index;not null;default:0;comment:角色id;column:role_id;"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20200819103202_create_admin_user_role_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='管理员角色关联表'"
			}
			err := tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&AdminUserRole{})
			if err != nil {
				return err
			}

			return tx.Create(&AdminUserRole{
				ID:      1,
				AdminID: 1,
				RoleID:  1,
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&AdminUserRole{})
		},
	})
}
