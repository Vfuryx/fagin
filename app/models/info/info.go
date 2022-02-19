package info

import (
	"time"

	"gorm.io/gorm"
)

// Info 用户详情模型
type Info struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserID    uint           `gorm:"index;not null;default:0;comment:'用户id'"`
	Email     string         `gorm:"type:varchar(100);not null;default:'';comment:'邮箱地址';"`
}

// Serializer 输出格式
type Serializer struct {
	ID        uint           `json:"id"`
	UserID    uint           `json:"user_id"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at"`
}

// Serializer 序列化接口
func (i *Info) Serializer() Serializer {
	return Serializer{
		ID:        i.ID,
		UserID:    i.UserID,
		Email:     i.Email,
		CreatedAt: i.CreatedAt,
		UpdatedAt: i.UpdatedAt,
		DeletedAt: i.DeletedAt,
	}
}
