package config

import (
	"fagin/pkg/conf"
)

type CacheConfig struct {
	Open      bool   // 是否开启缓存
	Prefix    string // 前缀
	DefDriver string // 默认
	Stores    map[string]map[string]string
}

var cacheConfig = new(CacheConfig)

func Cache() CacheConfig {
	return *cacheConfig
}

func (cache *CacheConfig) init() {
	cache.Open = conf.GetBool("cache.open")
	cache.Prefix = conf.GetString("cache.prefix", "app:")
	cache.DefDriver = conf.GetString("cache.driver", "big")
	cache.Stores = map[string]map[string]string{
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
