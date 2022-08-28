package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type AdminRoleMenu struct {
		ID uint `gorm:"primarykey"`

		RoleID uint `gorm:"uniqueIndex:udx_role_id_menu_id,priority:1;not null;default:0;comment:角色id;column:role_id;"`
		MenuID uint `gorm:"uniqueIndex:udx_role_id_menu_id,priority:2;index;not null;default:0;comment:菜单id;column:menu_id;"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20200818103136_create_admin_role_menu_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='后台角色关联菜单表'"
			}
			err := tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&AdminRoleMenu{})
			if err != nil {
				return err
			}

			adminRoleToMenu := []AdminRoleMenu{
				{
					ID:     1,
					RoleID: 1,
					MenuID: 1,
				},
				{
					ID:     2,
					RoleID: 1,
					MenuID: 2,
				},
				{
					ID:     3,
					RoleID: 1,
					MenuID: 3,
				},
				{
					ID:     4,
					RoleID: 1,
					MenuID: 4,
				},
				{
					ID:     5,
					RoleID: 1,
					MenuID: 7,
				},
				{
					ID:     6,
					RoleID: 1,
					MenuID: 8,
				},
				{
					ID:     7,
					RoleID: 1,
					MenuID: 9,
				},
				{
					ID:     8,
					RoleID: 1,
					MenuID: 10,
				},
				{
					ID:     9,
					RoleID: 1,
					MenuID: 13,
				},
				{
					ID:     10,
					RoleID: 1,
					MenuID: 14,
				},
				{
					ID:     11,
					RoleID: 1,
					MenuID: 15,
				},
				{
					ID:     12,
					RoleID: 1,
					MenuID: 16,
				},
				{
					ID:     13,
					RoleID: 1,
					MenuID: 17,
				}, {
					ID:     14,
					RoleID: 1,
					MenuID: 18,
				},
				{
					ID:     15,
					RoleID: 1,
					MenuID: 19,
				},
				{
					ID:     16,
					RoleID: 1,
					MenuID: 20,
				},
				{
					ID:     17,
					RoleID: 1,
					MenuID: 21,
				},
				{
					ID:     18,
					RoleID: 1,
					MenuID: 22,
				},
				{
					ID:     19,
					RoleID: 1,
					MenuID: 23,
				},
				{
					ID:     20,
					RoleID: 1,
					MenuID: 24,
				},
				{
					ID:     21,
					RoleID: 1,
					MenuID: 25,
				},
				{
					ID:     22,
					RoleID: 1,
					MenuID: 26,
				},
				{
					ID:     23,
					RoleID: 1,
					MenuID: 27,
				},
				{
					ID:     24,
					RoleID: 1,
					MenuID: 28,
				},
				{
					ID:     25,
					RoleID: 1,
					MenuID: 29,
				},
				{
					ID:     26,
					RoleID: 1,
					MenuID: 30,
				},
				{
					ID:     27,
					RoleID: 1,
					MenuID: 31,
				},
				{
					ID:     28,
					RoleID: 1,
					MenuID: 32,
				},
				{
					ID:     29,
					RoleID: 1,
					MenuID: 33,
				},
				{
					ID:     30,
					RoleID: 1,
					MenuID: 34,
				},
				{
					ID:     31,
					RoleID: 1,
					MenuID: 35,
				},
				{
					ID:     32,
					RoleID: 1,
					MenuID: 36,
				},
				{
					ID:     33,
					RoleID: 1,
					MenuID: 37,
				},
				{
					ID:     34,
					RoleID: 1,
					MenuID: 38,
				},
				{
					ID:     35,
					RoleID: 1,
					MenuID: 39,
				},
				{
					ID:     36,
					RoleID: 1,
					MenuID: 40,
				},
				{
					ID:     37,
					RoleID: 1,
					MenuID: 41,
				},
				{
					ID:     38,
					RoleID: 1,
					MenuID: 42,
				},
				{
					ID:     39,
					RoleID: 1,
					MenuID: 43,
				},
				{
					ID:     40,
					RoleID: 1,
					MenuID: 44,
				},
				{
					ID:     41,
					RoleID: 1,
					MenuID: 45,
				},
				{
					ID:     42,
					RoleID: 1,
					MenuID: 46,
				},
				{
					ID:     43,
					RoleID: 1,
					MenuID: 47,
				},
				{
					ID:     44,
					RoleID: 1,
					MenuID: 48,
				},
				{
					ID:     45,
					RoleID: 1,
					MenuID: 49,
				},
				{
					ID:     46,
					RoleID: 1,
					MenuID: 50,
				},
			}
			err = tx.Create(&adminRoleToMenu).Error
			if err != nil {
				return err
			}

			if config.DBDriverPostgres == config.Database().Driver() {
				// 设定自增序列
				err = tx.Exec(`SELECT setval( 'admin_role_menus_id_seq', ( SELECT MAX ( ID ) FROM admin_role_menus ) )`).Error
				if err != nil {
					return err
				}
			}

			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&AdminRoleMenu{})
		},
	})
}
