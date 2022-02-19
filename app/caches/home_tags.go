package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewHomeTags 首页tag
func NewHomeTags(f cache.GetterFunc) *cache.Cache {
	var c = new(cache.Cache)

	c.SetKeyFormat("home:tags")
	c.SetLifeTime(24 * time.Hour)
	c.SetFunc(f)

	return c
}
