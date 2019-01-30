package database

import (
	"fmt"
	"fagin/config"
	"github.com/jinzhu/gorm"
	_ "github.com/go-sql-driver/mysql"

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
