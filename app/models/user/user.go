package user

import (
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"
	"time"
)

type User struct {
	gorm.Model
	Name 		string		`gorm:"type: varchar(64); not null; column: name"`
	Avatar 		string		`gorm:"type: varchar(255); not null; default: ''; column: avatar"`
	Age 		uint8		`gorm:"type: tinyint unsigned; not null; default: 0; column: age"`
}


// 设定表名
//func (User) TableName() string {
//	return "users"
//}

// UserSerializer 接口输出格式
type UserSerializer struct {
	ID        	uint 		`json:"id"`
	Name 		string		`json:"name"`
	Avatar 		string		`json:"avatar"`
	Age 		uint8		`json:"age"`
	CreatedAt 	time.Time	`json:"created_at"`
	UpdatedAt 	time.Time	`json:"update_at"`
	DeletedAt 	*time.Time 	`json:"delete_at"`
}

// 序列化接口
func (user *User) Serializer() UserSerializer {
	return UserSerializer{
		ID: 		user.ID,
		Name: 		user.Name,
		Avatar: 	user.Avatar,
		Age: 		user.Age,
		CreatedAt: 	user.CreatedAt,
		UpdatedAt: 	user.UpdatedAt,
		DeletedAt: 	user.DeletedAt,
	}
}