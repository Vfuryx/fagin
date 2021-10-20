package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewAdminCasbin 后台 Casbin
func NewAdminCasbin(f cache.GetterFunc) *cache.Cache {
	var c = new(cache.Cache)
	c.SetKeyFormat("admin:casbin:roles:uid:%s")
	c.SetLifeTime(time.Minute)
	c.SetFunc(f)

	return c
}
