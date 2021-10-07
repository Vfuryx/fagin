package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewAdminOperationLog 后台操作缓存管理
func NewAdminOperationLog(f cache.GetterFunc) *cache.SCache {
	var c = new(cache.SCache)
	c.SetConfPrefix("admin:menu:log:%s:%s")
	c.SetConfLifeTime(30 * time.Second)
	c.SetFunc(f)

	return c
}
