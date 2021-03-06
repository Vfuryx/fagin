package author

import (
	"gorm.io/gorm"
	"time"
)

type Author struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:varchar(32);not null;default:'';comment:'名称';column:name"`
	Status    uint8          `gorm:"not null;default:1;comment:'状态 0:隐藏 1显示';column:status"`
}
