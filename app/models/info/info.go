package info

import (
	"gorm.io/gorm"
	"time"
)

// Info 用户详情模型
type Info struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	UserId    uint           `gorm:"index;not null;default:0;comment:'用户id'"`
	Email     string         `gorm:"type:varchar(100);not null;default:'';comment:'邮箱地址';"`
}

// InfoSerializer 输出格式
type InfoSerializer struct {
	ID        uint           `json:"id"`
	UserId    uint           `json:"user_id"`
	Email     string         `json:"email"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"update_at"`
	DeletedAt gorm.DeletedAt `json:"delete_at"`
}

// Serializer 序列化接口
func (this *Info) Serializer() InfoSerializer {
	return InfoSerializer{
		ID:        this.ID,
		UserId:    this.UserId,
		Email:     this.Email,
		CreatedAt: this.CreatedAt,
		UpdatedAt: this.UpdatedAt,
		DeletedAt: this.DeletedAt,
	}
}
