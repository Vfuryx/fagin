package admin_department

import (
	"time"
)

type AdminDepartment struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	ParentID uint   `gorm:"type:int(10) unsigned;not null;default:0;column:parent_id;"` // 父id
	Name     string `gorm:"type:varchar(128);not null;default:'';column:name;"`         // 部门名称
	Remark   string `gorm:"type:varchar(255);not null;default:'';column:remark;"`       // 部门备注
	Sort     uint   `gorm:"type:int(10) unsigned;not null;default:0;column:sort"`       // 排序，数字越大越靠前
	Status   uint8  `gorm:"type:tinyint(4) unsigned;not null;default:1;column:status;"` // 状态: 0=>关闭 1=>开启
}
