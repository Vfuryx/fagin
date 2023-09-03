package migrations

import (
	"fadmin/config"
	"fadmin/pkg/migrate"
	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
	"gorm.io/plugin/soft_delete"
)

func init() {
	type User struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt uint
		UpdatedAt uint
		DeletedAt soft_delete.DeletedAt `gorm:"index"`
		Username  string                `gorm:"type:varchar(45);not null;default:'';comment:用户名;column:username;"`
		Email     string                `gorm:"type:varchar(45);not null;default:'';comment:邮箱;column:email;"`
		Password  string                `gorm:"type:varchar(128);not null;default:'';comment:密码;column:password;"`
		Avatar    string                `gorm:"type:varchar(255);not null;default:'';comment:头像;column:avatar;"`
		Phone     string                `gorm:"type:varchar(30);not null;default:'';comment:手机;column:phone;"`
		Status    uint8                 `gorm:"not null;default:1;comment:状态: 0=>关闭 1=>开启;column:status;"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20200818094901_create_users_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='用户表'"
			}
			return tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&User{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&User{})
		},
	})
}
