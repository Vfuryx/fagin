package db

import (
	"errors"
	"fadmin/config"
	"fadmin/pkg/errorw"
	flog "fadmin/pkg/logger"
	"io"
	"log"
	"os"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var (
	orm                 *gorm.DB
	ErrDBDriverNotFound = errors.New("数据库不支持")
)

// Init 初始化
func Init() {
	var (
		err  error
		conf = config.Database()
	)

	var writer io.Writer
	if conf.Debug() {
		writer = io.MultiWriter(os.Stdout, flog.Channel("sql").Writer())
	} else {
		writer = flog.Channel("sql").Writer()
	}

	Logger := logger.New(
		log.New(writer, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second,     // 慢 SQL 阈值
			LogLevel:                  conf.LogLevel(), // Log level
			Colorful:                  false,           // 是否彩色打印
			IgnoreRecordNotFoundError: false,           // 是否忽略记录器的 ErrRecordNotFound 错误
		},
	)

	database := config.Database()
	link := database.GetConnectLink()

	switch database.Driver() {
	case config.DBDriverMysql:
		orm, err = gorm.Open(mysql.Open(link), &gorm.Config{
			Logger: Logger,
		})
	case config.DBDriverPostgres:
		orm, err = gorm.Open(postgres.Open(link), &gorm.Config{
			Logger: Logger,
		})
	default:
		panic(errorw.UP(ErrDBDriverNotFound))
	}

	if err != nil {
		panic(errorw.UP(err))
	}

	if orm.Error != nil {
		panic(errorw.UP(orm.Error))
	}

	if config.Database().Debug() {
		orm = orm.Debug()
	}

	sqlDB, err := orm.DB()
	if err != nil {
		panic(errorw.UP(err))
	}
	// SetMaxIdleConns 用于设置连接池中空闲连接的最大数量。
	sqlDB.SetMaxIdleConns(config.Database().MaxIdleConns())
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.Database().MaxOpenConns())
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(config.Database().ConnMaxLifetime()) * time.Minute)
}

// ORM orm
func ORM() *gorm.DB {
	if orm == nil {
		Init()
	}

	return orm
}

// Close 关闭
func Close() {
	db, err := orm.DB()
	if err != nil {
		return
	}

	_ = db.Close()
}
