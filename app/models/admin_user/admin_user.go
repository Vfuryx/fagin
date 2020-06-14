package admin_user

import (
	"time"
)

type AdminUser struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Status    uint8      `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
	Username  string     `gorm:"type:varchar(64);not null;default:'';comment:'用户名';column:username"`
	NickName  string     `gorm:"type:varchar(64);not null;default:'';comment:'昵称';column:nick_name"`
	Phone     string     `gorm:"type:varchar(64);not null;default:'';comment:'电话';column:phone"`
	Email     string     `gorm:"type:varchar(64);not null;default:'';comment:'邮箱';column:email;"`
	Password  string     `gorm:"type:varchar(127);not null;default:'';comment:'密码';column:password;"`
	Avatar    string     `gorm:"type:varchar(255);not null;default:'';comment:'头像';column:avatar;"`
	Remark    string     `gorm:"type:varchar(255);not null;default:'';comment:'备注';column:remark;"`
	Sex       uint8      `gorm:"type:tinyint(4) unsigned;not null;default:0;comment:'性别: 0:未知 1:男 2:女';column:sex;"`
	RoleID    uint       `gorm:"index;not null;default:0;comment:'角色ID';column:role_id;"`
}
