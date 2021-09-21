package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewAdminNavsCache 网站配置缓存管理
func NewAdminNavsCache(f cache.GetterFunc) *cache.SCache {
	var c = new(cache.SCache)
	c.SetConfPrefix("admin:navs:uid:%s")
	c.SetConfLifeTime(60 * 60 * 24 * time.Second)
	c.SetFunc(f)

	return c
}
