package caches

import (
	"fagin/pkg/cache"
	"time"
)

// 网站配置缓存管理
type websiteConfig struct {
	cache.SCache
}

func NewWebsiteConfig(f cache.GetterFunc) *websiteConfig {
	var c = new(websiteConfig)
	c.Prefix = "website:config:"
	c.LifeTime = 60 * time.Second
	c.Content = f
	c.SetFunc(c)
	return c
}

// 获取键名称
func (c *websiteConfig) Key(value string) string {
	return c.Prefix + value
}

// 默认存在时间
func (c *websiteConfig) Lift() time.Duration {
	return c.LifeTime
}
