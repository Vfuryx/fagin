package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewAdminRoutesCache 后台路由表缓存
func NewAdminRoutesCache(f cache.GetterFunc) *cache.Cache {
	var c = new(cache.Cache)
	c.SetKeyFormat("admin:routes:uid:%s")
	c.SetLifeTime(60 * 60 * 24 * time.Second)
	c.SetFunc(f)

	return c
}
