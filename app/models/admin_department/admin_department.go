package admin_department

import (
	"time"
)

type AdminDepartment struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time

	ParentID uint   `gorm:"not null;default:0;comment:父id;column:parent_id;"`
	Name     string `gorm:"type:varchar(128);not null;default:'';comment:部门名称;column:name;"`
	Remark   string `gorm:"type:varchar(255);not null;default:'';comment:部门备注;column:remark;"`
	Sort     uint   `gorm:"not null;default:0;comment:排序，数字越大越靠前;column:sort"`
	Status   uint8  `gorm:"not null;default:1;comment:状态: 0=>关闭 1=>开启;column:status;"`
}
