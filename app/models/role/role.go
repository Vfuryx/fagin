package role

import "time"

// 角色模型
type Role struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
	Name      string     `gorm:"type:varchar(64);not null;default:'';comment:'名称';column:name;"`
	Sort      uint64     `gorm:"type:int(4) unsigned;not null;default:0;comment:'排序 值越大越靠前"`
	Status    uint8      `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
}
