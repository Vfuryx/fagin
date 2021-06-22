package caches

import (
	"fagin/pkg/cache"
	"time"
)

// 网站配置缓存管理
type homeTags struct {
	cache.SCache
}

func NewHomeTags(f cache.GetterFunc) *homeTags {
	var c = new(homeTags)
	c.Prefix = "home::tags::"
	c.LifeTime = 24 * time.Hour
	c.Content = f
	c.SetFunc(c)
	return c
}

// 获取键名称
func (c *homeTags) Key(value string) string {
	return c.Prefix + value
}

// 默认存在时间
func (c *homeTags) Lift() time.Duration {
	return c.LifeTime
}
