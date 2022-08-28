package admin_menu

import (
	"time"

	"gorm.io/gorm"
)

// AdminMenu 后台菜单
type AdminMenu struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	ParentID      uint   `gorm:"not null;index;default:0;comment:父id;column:parent_id;"`
	Paths         string `gorm:"type:varchar(128);not null;default:'';comment:层级路径 0-1-2-;column:paths;"`
	Type          int    `gorm:"not null;index:idx_type_path_method,priority:1;default:0;comment:类型: 0:目录; 1:菜单; 2:权限;column:type;"`
	Name          string `gorm:"type:varchar(32);not null;default:'';comment:名称;column:name;"`
	Permission    string `gorm:"type:varchar(64);not null;default:'';comment:权限标记;column:permission;"`
	Title         string `gorm:"type:varchar(32);not null;default:'';comment:标题;column:title;"`
	Component     string `gorm:"type:varchar(255);not null;default:'';comment:组件名称;column:component;"`
	Path          string `gorm:"type:varchar(255);not null;index:idx_type_path_method,priority:2;default:'';comment:路由地址;column:path;"`
	Method        string `gorm:"type:varchar(16);not null;index:idx_type_path_method,priority:3;default:'';comment:请求方法;column:method;"`
	Redirect      string `gorm:"type:varchar(255);not null;default:'';comment:重定向;column:redirect;"`
	FrameSrc      string `gorm:"type:varchar(255);not null;default:'';comment:内嵌iframe的地址;column:frame_src;"`
	CurrentActive string `gorm:"type:varchar(255);not null;default:'';comment:当前激活的菜单;column:current_active;"`
	Icon          string `gorm:"type:varchar(128);not null;default:'';comment:图标;column:icon;"`
	IsShow        uint8  `gorm:"not null;default:1;comment:展示状态: 0=>隐藏 1=>展示;column:is_show;"`
	IsHideChild   uint8  `gorm:"not null;default:0;comment:是否隐藏子菜单: 0=>展示 1=>隐藏;column:is_hide_child"`
	IsNoCache     uint8  `gorm:"not null;default:0;comment:是否无缓存: 0=>否 1=>是;column:is_no_cache"`
	Sort          uint   `gorm:"not null;default:100;comment:排序，数字越大越靠前;column:sort"`
	Status        uint8  `gorm:"not null;default:1;comment:状态 0:关闭 1开启;column:status"`

	Children []AdminMenu `json:"children" gorm:"-"`
}
