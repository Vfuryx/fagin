package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewHomeArticles 首页文章
func NewHomeArticles(f cache.GetterFunc) *cache.Cache {
	var c = new(cache.Cache)
	c.SetKeyFormat("home:article:page:%s")
	c.SetLifeTime(60 * time.Minute)
	c.SetFunc(f)

	return c
}
