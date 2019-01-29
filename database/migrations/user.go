package migrations

import "time"

type User struct {
	ID        	uint 		`gorm:"primary_key"`
	Name 		string		`gorm:"type: varchar(64); not null; column: name"`
	Avatar 		string		`gorm:"type: varchar(255); not null; default: ''; column: avatar"`
	Age 		uint8		`gorm:"type: tinyint unsigned; not null; default: 0; column: age"`
	CreatedAt 	time.Time
	UpdatedAt 	time.Time
	DeletedAt 	*time.Time 	`sql:"index"`
}

////////////// 数据迁移start ////////////////////////////////////////////////////////////////////////////////////////////


/**
 * 创建users表
 */
func create_user_2018_01_23_00001()  {
	ORM.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&User{})
}

func delete_user_col_phone_2018_01_23_00002()  {
	ORM.Model(&User{}).DropColumn("phone")
}

////////////// 数据迁移end //////////////////////////////////////////////////////////////////////////////////////////////
