package admin_permission

import (
	group "fagin/app/models/admin_permission_group"
	"gorm.io/gorm"
	"time"
)

type AdminPermission struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt             `gorm:"index"`
	Type      uint8                      `gorm:"type:tinyint(4) unsigned;not null;default:0;comment:'菜单类型 0:总后台 1:商家后台 2:集团后台';column:type;"`
	GID       uint                       `gorm:"type:int(11) unsigned;not null;default:0;comment:'权限分组ID';column:gid;"`
	Group     group.AdminPermissionGroup `gorm:"foreignKey:gid;references:id"`
	Name      string                     `gorm:"type:varchar(35);not null;default:'';comment:'权限名称';column:name"`
	Path      string                     `gorm:"type:varchar(255);not null;default:'';comment:'路径';column:path"`
	Method    string                     `gorm:"type:varchar(16);not null;default:'';comment:'请求方法';column:method"`
	Sort      uint                       `gorm:"type:int(10) unsigned;not null;default:100;comment:'排序，数字越大越靠前';column:sort"`
	Status    uint8                      `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态 0:隐藏 1显示';column:status"`
	Content   string                     `gorm:"type:varchar(255);not null;default:'';comment:'接口描述';column:content"`
}
