package admin_permission_group

import (
	"gorm.io/gorm"
	"time"
)

type AdminPermissionGroup struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:varchar(35);not null;default:'';comment:'组用户';column:name"`
	Type      uint8          `gorm:"type:tinyint(4) unsigned;not null;default:0;comment:'菜单类型 0:总后台 1:商家后台 2:集团后台';column:type;"`
	Sort      uint           `gorm:"type:int(10) unsigned;not null;default:100;comment:'排序，数字越大越靠前';column:sort"`
}
