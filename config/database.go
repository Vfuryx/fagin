package config

import (
	"fagin/pkg/conf"
	"fmt"
)

type database struct {
	Host            string
	Port            string
	Name            string
	User            string
	Password        string
	Debug           bool
	MaxIdleConns    int // 设置最大空闲连接数
	MaxOpenConns    int // 设置最大连接数
	ConnMaxLifetime int // 设置每个链接的过期时间
}

var DB database

func init() {
	DB.Debug = conf.GetBool("db.debug")
	DB.Host = conf.GetString("db.host", "")
	DB.Port = conf.GetString("db.port", "")
	DB.Name = conf.GetString("db.name", "")
	DB.User = conf.GetString("db.user", "")
	DB.Password = conf.GetString("db.password", "")
	DB.MaxIdleConns = conf.GetInt("db.max_idle_conns", 25)
	DB.MaxOpenConns = conf.GetInt("db.max_open_conns", 100)
	DB.ConnMaxLifetime = conf.GetInt("db.conn_max_life_time", 5)
}

// GetConnectLink 获取 conn 链接
func (db *database) GetConnectLink() string {
	const link = "%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci"
	return fmt.Sprintf(link, db.User, db.Password, db.Host, db.Port, db.Name)
}
