package db

import (
	"fagin/config"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var (
	orm *gorm.DB = nil
	Err error
)

func Init() {
	link := config.DB.GetConnectLink()

	//newLogger := logger.New(
	//	log.New(Log.New("sql").Writer(), "\r\n", log.LstdFlags), // io writer
	//	logger.Config{
	//		SlowThreshold: time.Second,   // 慢 SQL 阈值
	//		LogLevel:      logger.Silent, // Log level
	//		Colorful:      false,         // 禁用彩色打印
	//	},
	//)

	orm, Err = gorm.Open(mysql.Open(link), &gorm.Config{
		//Logger: newLogger,
	})

	if config.DB.Debug {
		orm = orm.Debug()
	}

	if Err != nil {
		panic(fmt.Errorf("mysql connect exception %v \n", Err))
	}

	if orm.Error != nil {
		panic(fmt.Errorf("database exception %v \n", Err))
	}

	sqlDB, err := orm.DB()
	if err != nil {
		panic(fmt.Errorf("database exception %v \n", Err))
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(config.DB.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.DB.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(config.DB.ConnMaxLifetime) * time.Minute)
}

func ORM() *gorm.DB {
	if orm == nil {
		Init()
	}
	return orm
}

func Close() error {
	db, err := orm.DB()
	if err != nil {
		return err
	}
	err = db.Close()
	if err != nil {
		panic(err)
	}
	return err
}
