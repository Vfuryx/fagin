package config

import (
	"fagin/pkg/conf"
	"fmt"
	"sync/atomic"

	"gorm.io/gorm/logger"
)

var databaseConfig atomic.Value

// Database 数据库配置
func Database() *DatabaseConfig {
	if c, ok := databaseConfig.Load().(*DatabaseConfig); ok {
		return c
	}

	return &DatabaseConfig{}
}

// LogLevelMap map
var LogLevelMap = map[string]logger.LogLevel{
	"silent": logger.Silent,
	"error":  logger.Error,
	"warn":   logger.Warn,
	"info":   logger.Info,
}

const DBDriverMysql = "mysql"
const DBDriverPostgres = "postgres"

// DriverMap map
var DriverMap = map[string]struct{}{
	DBDriverMysql:    {},
	DBDriverPostgres: {},
}

func databaseConfigInit() {
	var ok bool

	c := &DatabaseConfig{
		debug:           conf.GetBool("db.debug"),
		host:            conf.GetString("db.host", ""),
		port:            conf.GetString("db.port", ""),
		name:            conf.GetString("db.name", ""),
		user:            conf.GetString("db.user", ""),
		password:        conf.GetString("db.password", ""),
		maxIdleConns:    conf.GetInt("db.max_idle_conns", 25),    //nolint:gomnd // 设置最大空闲连接数
		maxOpenConns:    conf.GetInt("db.max_open_conns", 100),   //nolint:gomnd // 设置最大连接数
		connMaxLifetime: conf.GetInt("db.conn_max_life_time", 5), //nolint:gomnd // 设置每个链接的过期时间 （分钟）
	}

	c.driver = conf.GetString("db.driver", DBDriverMysql)
	if _, ok = DriverMap[c.driver]; !ok {
		c.driver = DBDriverMysql
	}

	if c.logLevel, ok = LogLevelMap[conf.GetString("db.level", "error")]; !ok {
		c.logLevel = logger.Error
	}

	databaseConfig.Store(c)
}

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	driver          string
	debug           bool
	logLevel        logger.LogLevel // (silent error warn info)
	host            string
	port            string
	name            string
	user            string
	password        string
	maxIdleConns    int // 设置最大空闲连接数
	maxOpenConns    int // 设置最大连接数
	connMaxLifetime int // 设置每个链接的过期时间 （分钟）
}

func (db *DatabaseConfig) Driver() string {
	return db.driver
}

func (db *DatabaseConfig) Debug() bool {
	return db.debug
}

func (db *DatabaseConfig) LogLevel() logger.LogLevel {
	return db.logLevel
}

func (db *DatabaseConfig) Host() string {
	return db.host
}

func (db *DatabaseConfig) Port() string {
	return db.port
}

func (db *DatabaseConfig) Name() string {
	return db.name
}

func (db *DatabaseConfig) User() string {
	return db.user
}

func (db *DatabaseConfig) Password() string {
	return db.password
}

func (db *DatabaseConfig) MaxIdleConns() int {
	return db.maxIdleConns
}

func (db *DatabaseConfig) MaxOpenConns() int {
	return db.maxOpenConns
}

func (db *DatabaseConfig) ConnMaxLifetime() int {
	return db.connMaxLifetime
}

// GetConnectLink 获取 conn 链接
func (db *DatabaseConfig) GetConnectLink() string {
	switch db.driver {
	case DBDriverMysql:
		const link = "%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci"
		return fmt.Sprintf(link, db.User(), db.Password(), db.Host(), db.Port(), db.Name())
	case DBDriverPostgres:
		const link = "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai"
		return fmt.Sprintf(link, db.Host(), db.User(), db.Password(), db.Name(), db.Port())
	default:
		return ""
	}
}
