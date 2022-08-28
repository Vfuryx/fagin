package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"

	gormadapter "github.com/casbin/gorm-adapter/v3"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	migrate.Add(&gormigrate.Migration{
		ID: "20200818103300_create_casbin_rule_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='casbin_rule表'"
			}
			err := tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&gormadapter.CasbinRule{})
			if err != nil {
				return err
			}

			return tx.Create(&[]gormadapter.CasbinRule{
				{
					ID:    1,
					Ptype: "g",
					V0:    "1",
					V1:    "admin",
					V2:    "",
					V3:    "",
					V4:    "",
					V5:    "",
					V6:    "",
					V7:    "",
				},
				{
					ID:    2,
					Ptype: "p",
					V0:    "admin",
					V1:    "/*",
					V2:    "|",
					V3:    "",
					V4:    "",
					V5:    "",
					V6:    "",
					V7:    "",
				},
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&gormadapter.CasbinRule{})
		},
	})
}
