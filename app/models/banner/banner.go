package banner

import (
	"time"
)

type Banner struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Title     string     `gorm:"type:varchar(32);not null;default:'';comment:'标题';column:title"`
	Banner    string     `gorm:"type:varchar(255);not null;default:'';comment:'轮播图';column:banner"`
	Path      string     `gorm:"type:varchar(255);not null;default:'';comment:'路径';column:path"`
	Sort      uint       `gorm:"not null;default:100;comment:'排序，数字越大越靠前';column:sort"`
	Status    uint8      `gorm:"not null;default:1;comment:'状态 0:隐藏 1显示';column:status"`
}
