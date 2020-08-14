package db

import (
	"fagin/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	ORM *gorm.DB
	Err error
)

func init() {
	link := config.DB.GetConnectLink()

	ORM, Err = gorm.Open(mysql.Open(link), &gorm.Config{})

	if Err != nil {
		panic(fmt.Errorf("mysql connect exception %v \n", Err))
	}

	if ORM.Error != nil {
		panic(fmt.Errorf("database exception %v \n", Err))
	}
}

func Close() error {
	db, err := ORM.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
	return err
}
