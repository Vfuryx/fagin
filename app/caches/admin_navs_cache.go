package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewAdminRoutesCache 后台路由表缓存
func NewAdminRoutesCache(f cache.GetterFunc) *cache.SCache {
	var c = new(cache.SCache)
	c.SetConfPrefix("admin:routes:uid:%s")
	c.SetConfLifeTime(60 * 60 * 24 * time.Second)
	c.SetFunc(f)

	return c
}
