package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewHomeArticles 首页文章
func NewHomeArticles(f cache.GetterFunc) *cache.SCache {
	var c = new(cache.SCache)
	c.SetConfPrefix("home:article:page:%s")
	c.SetConfLifeTime(60 * time.Minute)
	c.SetFunc(f)

	return c
}
