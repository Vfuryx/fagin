package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewHomeNavbar 首页navbar
func NewHomeNavbar(f cache.GetterFunc) *cache.Cache {
	var c = new(cache.Cache)
	c.SetKeyFormat("home:navbar:cate")
	c.SetLifeTime(10 * time.Second)
	c.SetFunc(f)

	return c
}
