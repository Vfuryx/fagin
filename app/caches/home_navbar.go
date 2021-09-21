package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewHomeNavbar 首页navbar
func NewHomeNavbar(f cache.GetterFunc) *cache.SCache {
	var c = new(cache.SCache)
	c.SetConfPrefix("home:navbar:cate")
	c.SetConfLifeTime(10 * time.Second)
	c.SetFunc(f)

	return c
}
