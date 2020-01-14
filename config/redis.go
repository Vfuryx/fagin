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
	Redis.Open = conf.EnvBool("REDIS_IS_OPEN")
	Redis.Addr = conf.Env("REDIS_ADDRESS", "127.0.0.1")
	Redis.Password = conf.Env("REDIS_PASSWORD", "")
	Redis.Prefix = conf.Env("REDIS_PREFIX", "app:")
}
