package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewAdminOperationLog 后台操作缓存管理
func NewAdminOperationLog(f cache.GetterFunc) *cache.Cache {
	var c = new(cache.Cache)

	c.SetKeyFormat("admin:menu:log:%s:%s")
	c.SetLifeTime(30 * time.Second)
	c.SetFunc(f)

	return c
}
