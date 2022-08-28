package category

import (
	"time"

	"gorm.io/gorm"
)

type Category struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	Name   string `gorm:"type:varchar(32);not null;default:'';comment:名称;column:name"`
	Sort   uint   `gorm:"not null;default:100;comment:排序，数字越大越靠前;column:sort"`
	Status uint8  `gorm:"not null;default:1;comment:状态 0:隐藏 1显示;column:status"`
}
