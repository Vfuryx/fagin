package admin_menu

import (
	"time"
)

type AdminMenu struct {
	ID         uint `gorm:"primary_key"`
	CreatedAt  time.Time
	UpdatedAt  time.Time
	DeletedAt  *time.Time             `sql:"index"`
	ParentId   uint                   `gorm:"index;not null;default:0;comment:'父菜单id';column:parent_id;"`
	Paths      string                 `gorm:"type:varchar(128);not null;default:'';comment:'菜单层级路径 0-1-2';column:paths;"`
	Name       string                 `gorm:"type:varchar(32);not null;default:'';comment:'菜单名称';column:name;"`
	Title      string                 `gorm:"type:varchar(32);not null;default:'';comment:'菜单标题';column:title;"`
	Icon       string                 `gorm:"type:varchar(128);not null;default:'';comment:'图标';column:icon;"`
	Type       uint8                  `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'菜单类型';column:type;"`
	Path       string                 `gorm:"type:varchar(128);not null;default:'';comment:'路由地址';column:path;"`
	Component  string                 `gorm:"type:varchar(64);not null;default:'';comment:'组件路径';column:component;"`
	Action     string                 `gorm:"type:varchar(16);not null;default:'';comment:'菜单请求方法';column:action;"`
	Permission string                 `gorm:"type:varchar(32);not null;default:'';comment:'权限';column:permission;"`
	Sort       uint                   `gorm:"not null;default:100;comment:'排序，数字越大越靠前';column:sort"`
	Visible    uint8                  `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:visible;"`
	IsLink     uint8                  `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'是否外链';column:is_link;"`
	Children   []AdminMenu            `json:"children" gorm:"-"`
}
