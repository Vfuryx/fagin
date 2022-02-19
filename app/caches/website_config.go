package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewWebsiteConfig 网站配置缓存管理
func NewWebsiteConfig(f cache.GetterFunc) *cache.Cache {
	var c = new(cache.Cache)

	c.SetKeyFormat("website:config:info")
	c.SetLifeTime(60 * time.Second)
	c.SetFunc(f)

	return c
}
