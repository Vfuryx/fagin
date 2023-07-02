package config

import (
	"fadmin/pkg/conf"
	"sync/atomic"
)

// 读写并发安全
var cacheConfig atomic.Value

func Cache() *CacheConfig {
	if c, ok := cacheConfig.Load().(*CacheConfig); ok {
		return c
	}

	return &CacheConfig{}
}

func cacheConfigInit() {
	c := &CacheConfig{
		open:      conf.GetBool("cache.open"),
		prefix:    conf.GetString("cache.prefix", "app:"),
		defDriver: conf.GetString("cache.driver", "big"),
		stores: map[string]map[string]string{
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
		},
	}

	cacheConfig.Store(c)
}

type CacheConfig struct {
	open      bool   // 是否开启缓存
	prefix    string // 前缀
	defDriver string // 默认
	stores    map[string]map[string]string
}

func (c *CacheConfig) Open() bool {
	return c.open
}

func (c *CacheConfig) Prefix() string {
	return c.prefix
}

func (c *CacheConfig) DefDriver() string {
	return c.defDriver
}

func (c *CacheConfig) Stores() map[string]map[string]string {
	return c.stores
}
