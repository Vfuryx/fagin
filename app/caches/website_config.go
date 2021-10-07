package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewWebsiteConfig 网站配置缓存管理
func NewWebsiteConfig(f cache.GetterFunc) *cache.SCache {
	var c = new(cache.SCache)
	c.SetConfPrefix("website:config:info")
	c.SetConfLifeTime(60 * time.Second)
	c.SetFunc(f)

	return c
}
