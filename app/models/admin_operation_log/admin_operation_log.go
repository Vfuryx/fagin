package admin_operation_log

import (
	"time"

	"gorm.io/gorm"
)

type AdminOperationLog struct {
	ID        uint      `gorm:"primarykey"`
	CreatedAt time.Time `gorm:"index"`
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`

	AdminID     uint   `gorm:"not null;default:0;comment:管理员ID;column:admin_id;"`
	User        string `gorm:"type:varchar(32);not null;default:'';comment:用户;column:user;"`
	Method      string `gorm:"type:varchar(8);not null;default:'';comment:方法;column:method;"`
	Path        string `gorm:"type:varchar(255);not null;index;default:'';comment:路径;column:path;"`
	IP          string `gorm:"type:varchar(16);not null;default:'';comment:IP;column:ip;"`
	Location    string `gorm:"type:varchar(128);not null;default:'';comment:访问位置;column:location;"`
	Module      string `gorm:"type:varchar(32);not null;index;default:'';comment:操作模块;column:module;"`
	Operation   string `gorm:"type:varchar(32);not null;default:'';comment:操作类型;column:operation;"`
	Input       string `gorm:"type:text;comment:输入;column:input;"`
	LatencyTime string `gorm:"type:varchar(128);not null;default:'';comment:耗时;column:latency_time;"`
	UserAgent   string `gorm:"type:varchar(255);not null;default:'';comment:ua;column:user_agent;"`
	Status      uint8  `gorm:"not null;default:1;comment:状态 0:异常 1:正常;column:status;"`
}

// Operations
const (
	OperationLogin  = "登录"
	OperationLogout = "退出"
	OperationADD    = "新增"
	OperationUpdate = "修改"
	OperationDelete = "删除"
)
