package caches

import (
	"fagin/pkg/cache"
	"time"
)

func NewAdminRBAC(f cache.GetterFunc) *cache.Cache {
	var c = new(cache.Cache)
	// rbac::UID::uid::method::fullPath
	c.SetKeyFormat("admin:rbac:UID:%s:%s:%s")
	c.SetLifeTime(60 * time.Second)
	c.SetFunc(f)

	return c
}
