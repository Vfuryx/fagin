package migrations

import (
	"fagin/config"
	"fagin/pkg/migrate"
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

func init() {
	type AdminUser struct {
		ID        uint `gorm:"primarykey"`
		CreatedAt time.Time
		UpdatedAt time.Time
		DeletedAt gorm.DeletedAt `gorm:"index"`

		Status      uint8  `gorm:"not null;default:1;comment:状态: 0=>关闭 1=>开启;column:status;"`
		Username    string `gorm:"type:varchar(64);not null;index;default:'';comment:用户名;column:username"`
		NickName    string `gorm:"type:varchar(64);not null;default:'';comment:昵称;column:nick_name"`
		Phone       string `gorm:"type:varchar(64);not null;default:'';comment:电话;column:phone"`
		Email       string `gorm:"type:varchar(64);not null;default:'';comment:邮箱;column:email;"`
		Password    string `gorm:"type:varchar(127);not null;default:'';comment:密码;column:password;"`
		Avatar      string `gorm:"type:varchar(255);not null;default:'';comment:头像;column:avatar;"`
		Remark      string `gorm:"type:varchar(255);not null;default:'';comment:备注;column:remark;"`
		HomePath    string `gorm:"type:varchar(255);not null;default:'';comment:默认首页路径;column:home_path;"`
		Sex         uint8  `gorm:"not null;default:0;comment:性别: 0:未知 1:男 2:女;column:sex;"`
		LoginAt     uint64 `gorm:"not null;default:0;comment:总后台登录时间;column:login_at;"`
		LastLoginAt uint64 `gorm:"not null;default:0;comment:总后台上次登录时间;column:last_login_at;"`
		CheckInAt   uint64 `gorm:"not null;default:0;comment:总后台签发时间;column:check_in_at;"`
	}

	migrate.Add(&gormigrate.Migration{
		ID: "20200818103202_create_admin_user_table.go",
		Migrate: func(tx *gorm.DB) error {
			options := ""
			if config.DBDriverMysql == config.Database().Driver() {
				options = "COMMENT='后台用户表'"
			}
			err := tx.Set("gorm:table_options", options).
				Migrator().CreateTable(&AdminUser{})
			if err != nil {
				return err
			}

			return tx.Create(&AdminUser{
				Status:      1,
				Username:    "admin",
				NickName:    "admin",
				Phone:       "13800138000",
				Email:       "admin@admin.com",
				Password:    "$2a$10$v1Ji2oMpTD1sserUCuEfmeEUKQi4FKMonWah1fa4.xJhkxvPRvuN2",
				Avatar:      "",
				Remark:      "",
				HomePath:    "",
				Sex:         0,
				LoginAt:     1633228500,
				LastLoginAt: 1633228500,
				CheckInAt:   1633228500,
			}).Error
		},
		Rollback: func(tx *gorm.DB) error {
			return tx.Migrator().DropTable(&AdminUser{})
		},
	})
}
