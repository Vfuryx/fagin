package caches

import (
	"fagin/pkg/cache"
	"time"
)

// 网站配置缓存管理
type adminNavsCache struct {
	cache.SCache
}

func NewAdminNavsCache(f cache.GetterFunc) *adminNavsCache {
	var c = new(adminNavsCache)
	c.Prefix = "admin::navs::uid::"
	c.LifeTime = 60 * 60 * 24 * time.Second
	c.Content = f
	c.SetFunc(c)
	return c
}

// Key 获取键名称
func (c *adminNavsCache) Key(value string) string {
	return c.Prefix + value
}

// Lift 默认存在时间
func (c *adminNavsCache) Lift() time.Duration {
	return c.LifeTime
}
