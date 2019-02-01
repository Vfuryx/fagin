package models

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
		panic(fmt.Errorf("mysql connect error %v \n", err))
	}

	if orm.Error != nil {
		panic(fmt.Errorf("database error %v \n", err))
	}

	ORM = orm
}
