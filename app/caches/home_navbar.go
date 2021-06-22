package caches

import (
	"fagin/pkg/cache"
	"time"
)

// 网站配置缓存管理
type homeNavbar struct {
	cache.SCache
}

func NewHomeNavbar(f cache.GetterFunc) *homeNavbar {
	var c = new(homeNavbar)
	c.Prefix = "home::navbar::"
	c.LifeTime = 60 * time.Minute
	c.Content = f
	c.SetFunc(c)
	return c
}

// 获取键名称
func (c *homeNavbar) Key(value string) string {
	return c.Prefix + value
}

// 默认存在时间
func (c *homeNavbar) Lift() time.Duration {
	return c.LifeTime
}

