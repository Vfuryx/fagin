package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewWebArticle 文章
func NewWebArticle(f cache.GetterFunc) *cache.Cache {
	var c = new(cache.Cache)

	c.SetKeyFormat("web:article:id:%s")
	c.SetLifeTime(24 * time.Hour)
	c.SetFunc(f)

	return c
}
