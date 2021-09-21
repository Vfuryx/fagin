package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewAdminCasbin 后台 Casbin
func NewAdminCasbin(f cache.GetterFunc) *cache.SCache {
	var c = new(cache.SCache)
	c.SetConfPrefix("admin:casbin:roles:uid:%s")
	c.SetConfLifeTime(time.Minute)
	c.SetFunc(f)

	return c
}
