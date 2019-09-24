package config

import (
	"fagin/pkg/conf"
	"fmt"
)

type Database struct {
	Host 		string
	Port 		string
	Name 		string
	User 		string
	Password 	string
}

var DB = Database{}

func init() {
	DB.Host 	= conf.Env("DATABASE_HOST", "")
	DB.Port 	= conf.Env("DATABASE_PORT", "")
	DB.Name 	= conf.Env("DATABASE_NAME", "")
	DB.User 	= conf.Env("DATABASE_USER", "")
	DB.Password = conf.Env("DATABASE_PASSWORD", "")
}

// 获取 conn 链接
func (db *Database)GetConnectLink() string{
	return fmt.Sprintf(
		"%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
		db.User,db.Password,db.Host,db.Port,db.Name,
	)
}