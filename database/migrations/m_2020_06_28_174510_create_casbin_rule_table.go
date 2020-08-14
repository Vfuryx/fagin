package migrations

import (
	"fagin/pkg/migrate/migration"
	"gorm.io/gorm"
	"gopkg.in/gormigrate.v1"
)

// 自定义表
type CasbinRule struct {
	PType string `gorm:"type:varchar(100);default:NULL;comment:'p_type';column:p_type;"`
	V0    string `gorm:"type:varchar(100);default:NULL;comment:'v0';column:v0;"`
	V1    string `gorm:"type:varchar(100);default:NULL;comment:'v1';column:v1;"`
	V2    string `gorm:"type:varchar(100);default:NULL;comment:'v2';column:v2;"`
	V3    string `gorm:"type:varchar(100);default:NULL;comment:'v3';column:v3;"`
	V4    string `gorm:"type:varchar(100);default:NULL;comment:'v4';column:v4;"`
	V5    string `gorm:"type:varchar(100);default:NULL;comment:'v5';column:v5;"`
}
func (CasbinRule) TableName() string {
	return "casbin_rule"
}

func init() {

	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
		ID: "m_2020_06_28_174510_create_casbin_rule_table",
		Migrate: func(tx *gorm.DB) error {
			err := tx.AutoMigrate(&CasbinRule{}).Error
			if err != nil {
				return err
			}
			crs := []CasbinRule{
				{
					PType: "g",
					V0:    "1",
					V1:    "admin",
				},
				{
					PType: "p",
					V0:    "admin",
					V1:    "/*",
					V2:    "|",
				},
			}
			for _, cr := range crs {
				err := tx.Create(&cr).Error
				if err != nil {
					return err
				}
			}
			return nil
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.DropTableIfExists(&CasbinRule{}).Error
		},
	})
}
