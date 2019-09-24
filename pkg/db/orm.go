package db

import (
	"fagin/config"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

var (
	ORM *gorm.DB
	Err error
)

func init()  {
	link := config.DB.GetConnectLink()

	ORM, Err = gorm.Open("mysql", link)

	if Err != nil {
		panic(fmt.Errorf("mysql connect exception %v \n", Err))
	}

	if ORM.Error != nil {
		panic(fmt.Errorf("database exception %v \n", Err))
	}
}

func Close() {
	Err = ORM.Close()
	if Err != nil {
		panic(Err)
	}
}