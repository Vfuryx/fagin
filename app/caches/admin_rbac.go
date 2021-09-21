package caches

import (
	"fagin/pkg/cache"
	"time"
)

func NewAdminRBAC(f cache.GetterFunc) *cache.SCache {
	var c = new(cache.SCache)
	// rbac::UID::uid::method::fullPath
	c.SetConfPrefix("admin:rbac:UID:%s:%s:%s")
	c.SetConfLifeTime(60 * time.Second)
	c.SetFunc(f)

	return c
}
