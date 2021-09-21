package config

import (
	"fagin/pkg/conf"
)

type cache struct {
	Open      bool   // 是否开启缓存
	Prefix    string // 前缀
	DefDriver string // 默认
	Stores    map[string]map[string]string
}

var Cache cache

func init() {
	Cache.Open = conf.GetBool("cache.open")

	Cache.Prefix = conf.GetString("cache.prefix", "app:")

	Cache.DefDriver = conf.GetString("cache.driver", "big")

	Cache.Stores = map[string]map[string]string{
		"big": {
			"eviction": conf.GetString("cache.big.eviction", "3600"),
		},
		"redis": {
			"address":  conf.GetString("cache.redis.address", "127.0.0.1"),
			"password": conf.GetString("cache.redis.password", ""),
		},
		"ristretto": {
			"num_counters": conf.GetString("cache.ristretto.num_counters", "100000"),
			"max_cost":     conf.GetString("cache.ristretto.max_cost", "10000"),
			"buffer_items": conf.GetString("cache.ristretto.buffer_items", "64"),
		},
	}
}
