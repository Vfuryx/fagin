package admin_user

import (
	r "fagin/app/models/admin_role"
	"gorm.io/gorm"
	"time"
)

// AdminUserGuard guard
const AdminUserGuard = "admin"
const AdminUserIDKey = "admin_user_id"
const AdminUserNameKey = "admin_user_name"

// AdminUser 管理员
type AdminUser struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Status      uint8          `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
	Username    string         `gorm:"type:varchar(64);not null;default:'';comment:'用户名';column:username"`
	NickName    string         `gorm:"type:varchar(64);not null;default:'';comment:'昵称';column:nick_name"`
	Phone       string         `gorm:"type:varchar(64);not null;default:'';comment:'电话';column:phone"`
	Email       string         `gorm:"type:varchar(64);not null;default:'';comment:'邮箱';column:email;"`
	Password    string         `gorm:"type:varchar(127);not null;default:'';comment:'密码';column:password;"`
	Avatar      string         `gorm:"type:varchar(255);not null;default:'';comment:'头像';column:avatar;"`
	Remark      string         `gorm:"type:varchar(255);not null;default:'';comment:'备注';column:remark;"`
	Sex         uint8          `gorm:"type:tinyint(4) unsigned;not null;default:0;comment:'性别: 0:未知 1:男 2:女';column:sex;"`
	LoginAt     uint64         `gorm:"type:bigint(19) unsigned;not null;default:0;comment:'总后台登录时间';column:login_at;"`
	LastLoginAt uint64         `gorm:"type:bigint(19) unsigned;not null;default:0;comment:'总后台上次登录时间';column:last_login_at;"`
	CheckInAt   uint64         `gorm:"type:bigint(19) unsigned;not null;default:0;comment:'总后台签发时间';column:check_in_at;"`
	Roles       []r.AdminRole  `gorm:"many2many:admin_user_roles;joinForeignKey:admin_id;JoinReferences:role_id;"`
}
