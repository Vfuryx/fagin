package website_config

import (
	"time"

	"gorm.io/gorm"
)

type WebsiteConfig struct {
	ID                   uint `gorm:"primarykey"`
	CreatedAt            time.Time
	UpdatedAt            time.Time
	DeletedAt            gorm.DeletedAt `gorm:"index"`
	WebName              string         `gorm:"type:varchar(32);not null;default:'';comment:网站名称;column:web_name"`
	WebEnName            string         `gorm:"type:varchar(32);not null;default:'';comment:网站英文名称;column:web_en_name"`
	WebTitle             string         `gorm:"type:varchar(32);not null;default:'';comment:网站标题;column:web_title"`
	Keywords             string         `gorm:"type:varchar(127);not null;default:'';comment:关键词;column:keywords"`
	Description          string         `gorm:"type:varchar(255);not null;default:'';comment:描述;column:description"`
	CompanyName          string         `gorm:"type:varchar(32);not null;default:'';comment:公司名称;column:company_name"`
	ContactNumber        string         `gorm:"type:varchar(16);not null;default:'';comment:联系电话;column:contact_number"`
	CompanyAddress       string         `gorm:"type:varchar(32);not null;default:'';comment:公司地址;column:company_address"`
	Email                string         `gorm:"type:varchar(32);not null;default:'';comment:邮箱地址;column:email"`
	ICP                  string         `gorm:"type:varchar(32);not null;default:'';comment:ICP备案;column:icp"`
	PublicSecurityRecord string         `gorm:"type:varchar(32);not null;default:'';comment:公安部备案;column:public_security_record"`
	WebLogo              string         `gorm:"type:varchar(255);not null;default:'';comment:网站logo;column:web_logo"`
	QRCode               string         `gorm:"type:varchar(255);not null;default:'';comment:二维码;column:qr_code"`
}
