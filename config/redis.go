package config

import (
	"fagin/pkg/conf"
)

type redis struct {
	Open     bool   // 是否开启缓存
	Prefix   string // 前缀
	Addr     string // 地址
	Password string // 密码
}

var Redis redis

func init() {
	Redis.Open = conf.GetBool("redis.open")
	Redis.Addr = conf.GetString("redis.address", "127.0.0.1")
	Redis.Password = conf.GetString("redis.password", "")
	Redis.Prefix = conf.GetString("redis.prefix", "app:")
}
