package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewHomeTags 首页tag
func NewHomeTags(f cache.GetterFunc) *cache.SCache {
	var c = new(cache.SCache)
	c.SetConfPrefix("home:tags")
	c.SetConfLifeTime(24 * time.Hour)
	c.SetFunc(f)

	return c
}
