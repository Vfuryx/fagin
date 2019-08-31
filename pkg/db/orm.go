package db

import (
	"fagin/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var ORM *gorm.DB

func init()  {
	var err error

	link := config.DB.GetConnectLink();

	orm, err := gorm.Open("mysql", link)

	if err != nil {
		panic(fmt.Errorf("mysql connect exception %v \n", err))
	}

	if orm.Error != nil {
		panic(fmt.Errorf("database exception %v \n", err))
	}

	ORM = orm
}
