package config

import (
	"fmt"
	"github.com/spf13/viper"
)

type DBConfig struct {
	Host 		string
	Port 		string
	DBName 		string
	User 		string
	Password 	string
}

func (DB DBConfig) GetConnectLink() string{

	DB.Host 	= viper.GetString("db.host")
	DB.Port 	= viper.GetString("db.port")
	DB.DBName 	= viper.GetString("db.dbname")
	DB.User 	= viper.GetString("db.user")
	DB.Password = viper.GetString("db.password")

	return fmt.Sprintf("%s:%s@(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",DB.User,DB.Password,DB.Host,DB.Port,DB.DBName)
}