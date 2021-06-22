package caches

import (
	"fagin/pkg/cache"
	"time"
)

// 网站配置缓存管理
type homeArticles struct {
	cache.SCache
}

func NewHomeArticles(f cache.GetterFunc) *homeArticles {
	var c = new(homeArticles)
	c.Prefix = "home::article::"
	c.LifeTime = 60 * time.Minute
	c.Content = f
	c.SetFunc(c)
	return c
}

// 获取键名称
func (c *homeArticles) Key(value string) string {
	return c.Prefix + value
}

// 默认存在时间
func (c *homeArticles) Lift() time.Duration {
	return c.LifeTime
}
