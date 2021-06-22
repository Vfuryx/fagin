package caches

import (
	"fagin/pkg/cache"
	"time"
)

// 网站配置缓存管理
type webArticle struct {
	cache.SCache
}

func NewWebArticle(f cache.GetterFunc) *webArticle {
	var c = new(webArticle)
	c.Prefix = "web::article::id::"
	c.LifeTime = 24 * time.Hour
	c.Content = f
	c.SetFunc(c)
	return c
}

// 获取键名称
func (c *webArticle) Key(value string) string {
	return c.Prefix + value
}

// 默认存在时间
func (c *webArticle) Lift() time.Duration {
	return c.LifeTime
}
