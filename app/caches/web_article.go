package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewWebArticle 文章
func NewWebArticle(f cache.GetterFunc) *cache.SCache {
	var c = new(cache.SCache)
	c.SetConfPrefix("web:article:id:%s")
	c.SetConfLifeTime(24 * time.Hour)
	c.SetFunc(f)

	return c
}
