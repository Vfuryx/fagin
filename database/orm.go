package database

import (
	"fmt"
	"github.com/Vfuryx/fagin/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

)

var ORM *gorm.DB

func init()  {
	var err error

	link := config.DBConfig{}.GetConnectLink();

	orm, err := gorm.Open("mysql", link)

	if err != nil {
		fmt.Printf("mysql connect error %v", err)
	}

	if orm.Error != nil {
		fmt.Printf("database error %v", orm.Error)
	}

	ORM = orm
}
