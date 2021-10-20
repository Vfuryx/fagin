package config

import (
	"fagin/pkg/conf"
	"fmt"
)

// DatabaseConfig 数据库配置
type DatabaseConfig struct {
	Host            string
	Port            string
	Name            string
	User            string
	Password        string
	Debug           bool
	MaxIdleConns    int // 设置最大空闲连接数
	MaxOpenConns    int // 设置最大连接数
	ConnMaxLifetime int // 设置每个链接的过期时间 （分钟）
}

var databaseConfig = new(DatabaseConfig)

// Database 数据库配置
func Database() DatabaseConfig {
	return *databaseConfig
}

func (db *DatabaseConfig) init() {
	db.Debug = conf.GetBool("db.debug")
	db.Host = conf.GetString("db.host", "")
	db.Port = conf.GetString("db.port", "")
	db.Name = conf.GetString("db.name", "")
	db.User = conf.GetString("db.user", "")
	db.Password = conf.GetString("db.password", "")
	db.MaxIdleConns = conf.GetInt("db.max_idle_conns", 25)
	db.MaxOpenConns = conf.GetInt("db.max_open_conns", 100)
	db.ConnMaxLifetime = conf.GetInt("db.conn_max_life_time", 5)
}

// GetConnectLink 获取 conn 链接
func (db DatabaseConfig) GetConnectLink() string {
	const link = "%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci"
	return fmt.Sprintf(link, db.User, db.Password, db.Host, db.Port, db.Name)
}
