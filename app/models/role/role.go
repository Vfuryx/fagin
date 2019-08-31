package role

import (
	"github.com/jinzhu/gorm"
	"time"
)

// 角色模型
type Role struct {
	gorm.Model
	Name   string `gorm:"type:varchar(64);not null;default:'';comment:'名称';column:name;"`
	Sort   uint64 `gorm:"type:int(4) unsigned;not null;default:0;comment:'排序 值越大越靠前"`
	Status uint8  `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
}

// 输出格式
type RoleSerializer struct {
	ID        uint       `json:"id"`
	Name      string     `json:"name"`
	Sort      uint64     `json:"sort"`
	Status    uint8      `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"update_at"`
	DeletedAt *time.Time `json:"delete_at"`
}

// 序列化接口
func (this *Role) Serializer() RoleSerializer {
	return RoleSerializer{
		ID:        this.ID,
		Name:      this.Name,
		Sort:      this.Sort,
		Status:    this.Status,
		CreatedAt: this.CreatedAt,
		UpdatedAt: this.UpdatedAt,
		DeletedAt: this.DeletedAt,
	}
}
