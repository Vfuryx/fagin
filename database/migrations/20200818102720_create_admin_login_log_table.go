package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type AdminLoginLog struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`

		LoginName string `gorm:"type:varchar(32);not null;default:'';comment:登录账号;column:login_name;"`
		IP        string `gorm:"type:varchar(16);not null;default:'';comment:登录IP地址;column:ip;"`
		Location  string `gorm:"type:varchar(128);not null;default:'';comment:登录地点;column:location;"`
		Browser   string `gorm:"type:varchar(32);not null;default:'';comment:浏览器类型;column:browser;"`
		OS        string `gorm:"type:varchar(32);not null;default:'';comment:操作系统;column:os;"`
		Status    uint8  `gorm:"not null;default:1;comment:状态 0:异常 1:正常;column:status;"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20200818102720_create_admin_login_log_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='后台登录日志表'"
			}
			return tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&AdminLoginLog{})
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&AdminLoginLog{})
		},
	})
}
