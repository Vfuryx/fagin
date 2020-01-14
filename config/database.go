package config

import (
	"fagin/pkg/conf"
	"fmt"
)

type database struct {
	Host     string
	Port     string
	Name     string
	User     string
	Password string
}

var DB database

func init() {
	DB.Host = conf.Env("DATABASE_HOST", "")
	DB.Port = conf.Env("DATABASE_PORT", "")
	DB.Name = conf.Env("DATABASE_NAME", "")
	DB.User = conf.Env("DATABASE_USER", "")
	DB.Password = conf.Env("DATABASE_PASSWORD", "")
}

// 获取 conn 链接
func (db *database) GetConnectLink() string {
	const link = "%s:%s@(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local&collation=utf8mb4_unicode_ci"
	return fmt.Sprintf(link, db.User, db.Password, db.Host, db.Port, db.Name)
}
