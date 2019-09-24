package user

import (
	"fagin/app/models/info"
	"github.com/jinzhu/gorm"
	"time"
)

// 用户模型
type User struct {
	gorm.Model
	Username string `gorm:"type:varchar(64);not null;default:'';comment:'用户名';column:username;"`
	Email    string `gorm:"type:varchar(64);not null;default:'';comment:'邮箱';column:email;"`
	Password string `gorm:"type:varchar(127);not null;default:'';comment:'密码';column:password;"`
	Avatar   string `gorm:"type:varchar(255);not null;default:'';comment:'头像';column:avatar;"`
	Status   uint8  `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
	Phone    string `gorm:"type:varchar(20);not null;default:'';comment:'手机';column:phone;"`
	Info	 info.Info
}

// 设定表名
//func (User) TableName() string {
//	return "users"
//}

// UserSerializer 接口输出格式
type UserSerializer struct {
	ID        uint       `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Avatar    string     `json:"avatar"`
	Status    uint8      `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `json:"delete_at"`
}

// 序列化接口
func (user *User) Serializer() UserSerializer {
	return UserSerializer{
		ID:        user.ID,
		Username:  user.Username,
		Email:     user.Email,
		Avatar:    user.Avatar,
		Status:    user.Status,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		DeletedAt: user.DeletedAt,
	}
}
