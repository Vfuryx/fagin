package admin_user

import (
	"time"
)

type AdminUser struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Username  string     `gorm:"type:varchar(64);not null;default:'';comment:'用户名';comment:'username';"`
	Email     string     `gorm:"type:varchar(64);not null;default:'';comment:'邮箱';column:email;"`
	Password  string     `gorm:"type:varchar(127);not null;default:'';comment:'密码';column:password;"`
	Avatar    string     `gorm:"type:varchar(255);not null;default:'';comment:'头像';column:avatar;"`
	Status    uint8      `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
}
