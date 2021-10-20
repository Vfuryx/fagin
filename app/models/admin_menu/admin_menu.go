package admin_menu

import (
	"gorm.io/gorm"
	"time"
)

// AdminMenu 后台菜单
type AdminMenu struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ParentID      uint   `gorm:"type:int(10) unsigned;not null;default:0;column:parent_id;"`       // 父id
	Paths         string `gorm:"type:varchar(128);not null;default:'';column:paths;"`              // 层级路径 0-1-2
	Type          int    `gorm:"type:tinyint(4) unsigned;not null;default:0;column:type;"`         // 类型: 0:目录; 1:菜单; 2:权限;
	Name          string `gorm:"type:varchar(32);not null;default:'';column:name;"`                // 名称
	Permission    string `gorm:"type:varchar(32);not null;default:'';column:permission;"`          // 权限
	Title         string `gorm:"type:varchar(32);not null;default:'';column:title;"`               // 标题
	Component     string `gorm:"type:varchar(32);not null;default:'';column:component;"`           // 组件名称
	Path          string `gorm:"type:varchar(255);not null;default:'';column:path;"`               // 路由地址
	Method        string `gorm:"type:varchar(16);not null;default:'';column:method;"`              // 请求方法
	Redirect      string `gorm:"type:varchar(255);not null;default:'';column:redirect;"`           // 重定向
	FrameSrc      string `gorm:"type:varchar(255);not null;default:'';column:frame_src;"`          // 内嵌iframe的地址
	CurrentActive string `gorm:"type:varchar(255);not null;default:'';column:current_active;"`     // 当前激活的菜单
	Icon          string `gorm:"type:varchar(128);not null;default:'';column:icon;"`               // 图标
	IsShow        uint8  `gorm:"type:tinyint(4) unsigned;not null;default:1;column:is_show;"`      // 状态: 0=>隐藏 1=>展示
	IsHideChild   uint8  `gorm:"type:tinyint(4) unsigned;not null;default:0;column:is_hide_child"` // 是否隐藏子菜单 0:关闭 1开启
	IsNoCache     uint8  `gorm:"type:tinyint(4) unsigned;not null;default:0;column:is_no_cache"`   // 是否无缓存: 0=>否 1=>是
	Sort          uint   `gorm:"type:int(10) unsigned;not null;default:100;column:sort"`           // 排序，数字越大越靠前
	Status        uint8  `gorm:"type:tinyint(4) unsigned;not null;default:1;column:status"`        // 状态 0:关闭 1开启

	Children []AdminMenu `json:"children" gorm:"-"`
}
