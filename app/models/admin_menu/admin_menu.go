package admin_menu

import (
	"gorm.io/gorm"
	"time"
)

// AdminMenu 后台菜单
type AdminMenu struct {
	ID          uint `gorm:"primarykey"`
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   gorm.DeletedAt `gorm:"index"`
	Type        uint8          `gorm:"type:tinyint(4) unsigned;not null;default:0;comment:'菜单类型 0:总后台 1:商家后台 2:集团后台';column:type;"`
	ParentID    uint           `gorm:"type:int(10) unsigned;not null;default:0;comment:'父菜单id';column:parent_id;"`
	Paths       string         `gorm:"type:varchar(128);not null;default:'';comment:'菜单层级路径 0-1-2';column:paths;"`
	Title       string         `gorm:"type:varchar(32);not null;default:'';comment:'菜单标题';column:title;"`
	Name        string         `gorm:"type:varchar(32);not null;default:'';comment:'菜单名称';column:name;"`
	Component   string         `gorm:"type:varchar(32);not null;default:'';comment:'组件名称';column:component;"`
	Icon        string         `gorm:"type:varchar(128);not null;default:'';comment:'图标';column:icon;"`
	Path        string         `gorm:"type:varchar(128);not null;default:'';comment:'路由地址';column:path;"`
	IsShow      uint8          `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>隐藏 1=>展示';column:is_show;"`
	Redirect    string         `gorm:"type:varchar(128);not null;default:'';comment:'重定向';column:redirect;"`
	Target      string         `gorm:"type:varchar(128);not null;default:'';comment:'目标';column:target;"`
	IsHideChild uint8          `gorm:"type:tinyint(4) unsigned;not null;default:0;comment:'是否隐藏子菜单 0:关闭 1开启';column:is_hide_child"`
	Permission  string         `gorm:"type:varchar(32);not null;default:'';comment:'权限';column:permission;"`
	Sort        uint           `gorm:"type:int(10) unsigned;not null;default:100;comment:'排序，数字越大越靠前';column:sort"`
	Status      uint8          `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态 0:关闭 1开启';column:status"`
	Method      string         `gorm:"type:varchar(16);not null;default:'';comment:'菜单请求方法';column:method;"`
	Children    []AdminMenu    `json:"children" gorm:"-"`
}
