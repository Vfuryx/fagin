package admin_role

import (
	"fagin/app/models/admin_menu"
	"time"
)

type AdminRole struct {
	ID        uint `gorm:"primary_key"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Name      string                  `gorm:"type:varchar(128);not null;default:'';comment:'角色名称';column:name;"`
	Key       string                  `gorm:"type:varchar(32);not null;default:'';comment:'角色关键字';column:key;"`
	Remark    string                  `gorm:"type:varchar(255);not null;default:'';comment:'角色备注';column:remark;"`
	Sort      uint                    `gorm:"not null;default:0;comment:'排序，数字越大越靠前';column:sort"`
	Status    uint8                   `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
	Menus     []admin_menu.AdminMenu  `gorm:"many2many:admin_role_menus;joinForeignKey:role_id;JoinReferences:menu_id;"`
}
