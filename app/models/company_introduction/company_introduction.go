package company_introduction

import (
	"time"

	"gorm.io/gorm"
)

// CompanyIntroduction 公司介绍表
type CompanyIntroduction struct {
	ID        uint `gorm:"primarykey"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt `gorm:"index"`
	Name      string         `gorm:"type:varchar(32);not null;default:'';comment:'公司名称';column:name"`
	Image     string         `gorm:"type:varchar(255);not null;default:'';comment:'图片';column:image"`
	Content   string         `gorm:"type:text;not null;comment:'内容';column:content"`
	Status    uint8          `gorm:"type:tinyint(4) unsigned;not null;default:1;comment:'状态: 0=>关闭 1=>开启';column:status;"`
}
