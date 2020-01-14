package user

import (
	"fagin/app/models/info"
	"time"
)

// 用户模型
type User struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Username  string     `gorm:"type:varchar(64);not null;default:'';comment:'用户名';column:username;"`
	Email     string     `gorm:"type:varchar(64);not null;default:'';comment:'邮箱';column:email;"`
	Password  string     `gorm:"type:varchar(127);not null;default:'';comment:'密码';column:password;"`
	Avatar    string     `gorm:"type:varchar(255);not null;default:'';comment:'头像';column:avatar;"`
	Status    uint8      `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
	Phone     string     `gorm:"type:varchar(20);not null;default:'';comment:'手机';column:phone;"`
	Info      info.Info
}

// 设定表名
//func (User) TableName() string {
//	return "users"
//}
