package admin_role

import (
	m "fagin/app/models/admin_menu"
	"time"
)

type AdminRole struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	Type      uint8         `gorm:"type:tinyint(4) unsigned;not null;default:0;comment:'菜单类型 0:总后台 1:商家后台 2:集团后台';column:type;"`
	Name      string        `gorm:"type:varchar(128);not null;default:'';comment:'角色名称';column:name;"`
	Key       string        `gorm:"type:varchar(32);not null;default:'';comment:'角色关键字';column:key;"`
	Remark    string        `gorm:"type:varchar(255);not null;default:'';comment:'角色备注';column:remark;"`
	Sort      uint          `gorm:"type:int(10) unsigned;not null;default:0;comment:'排序，数字越大越靠前';column:sort"`
	Status    uint8         `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
	Menus     []m.AdminMenu `gorm:"many2many:admin_role_menus;joinForeignKey:role_id;JoinReferences:menu_id;"`
}
