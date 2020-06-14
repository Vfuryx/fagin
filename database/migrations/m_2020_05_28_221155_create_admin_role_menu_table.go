package migrations

import (
	"fagin/pkg/migrate/migration"
	"github.com/jinzhu/gorm"
	"gopkg.in/gormigrate.v1"
)

func init() {
	// 角色菜单关联表
	type AdminRoleMenu struct {
		RoleID uint `gorm:"type:int(10);not null;default:0;comment:'角色id';column:role_id;"`
		MenuID uint `gorm:"type:int(10);not null;default:0;comment:'菜单id';column:menu_id;"`
	}
	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
		ID: "m_2020_05_28_221155_create_admin_role_menu_table",
		Migrate: func(tx *gorm.DB) error {
			return tx.AutoMigrate(&AdminRoleMenu{}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTableIfExists(&AdminRoleMenu{}).Error
		},
	})
}
