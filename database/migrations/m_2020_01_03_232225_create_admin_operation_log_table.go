package migrations

//import (
//	"fagin/pkg/migrate/migration"
//	"github.com/jinzhu/gorm"
//	"gopkg.in/gormigrate.v1"
//	"time"
//)
//
//func init() {
//	// 自定义表
//	type AdminOperationLog struct {
//		ID          uint64 `gorm:"primary_key"`
//		CreatedAt   time.Time
//		UpdatedAt   time.Time
//		DeletedAt   *time.Time `sql:"index"`
//		User        string     `gorm:"type:varchar(32);not null;default:'';comment:'用户';column:user;"`
//		Method      string     `gorm:"type:varchar(8);not null;default:'';comment:'方法';column:method;"`
//		Path        string     `gorm:"type:varchar(255);not null;default:'';comment:'路径';column:path;"`
//		IP          string     `gorm:"type:varchar(16);not null;default:'';comment:'IP';column:ip;"`
//		Location    string     `gorm:"type:varchar(128);not null;default:'';comment:'访问位置';column:location;"` //访问位置
//		Module      string     `gorm:"type:varchar(32);not null;default:'';comment:'操作模块';column:module;"`
//		Operation   string     `gorm:"type:varchar(32);not null;default:'';comment:'操作类型';column:operation;"`
//		Input       string     `gorm:"type:text;comment:'输入';column:input;"`
//		LatencyTime string     `gorm:"type:varchar(128);not null;default:'';comment:'耗时';column:latency_time;"` //耗时
//		UserAgent   string     `gorm:"type:varchar(255);not null;default:'';comment:'ua';column:user_agent;"`   //ua
//		Status      uint8      `gorm:"not null;default:0;comment:'状态 0:异常 1:正常';column:status;"` //操作状态
//	}
//	migration.Migrations = append(migration.Migrations, &gormigrate.Migration{
//		ID: "m_2020_01_03_232225_create_admin_operation_log_table",
//		Migrate: func(tx *gorm.DB) error {
//			return tx.AutoMigrate(&AdminOperationLog{}).Error
//		},
//		Rollback: func(tx *gorm.DB) error {
//			return tx.DropTableIfExists(&AdminOperationLog{}).Error
//		},
//	})
//}
