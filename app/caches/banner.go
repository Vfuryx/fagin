package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewBanner 轮播图缓存管理
func NewBanner(f cache.GetterFunc) *cache.SCache {
	var b = new(cache.SCache)
	b.SetConfPrefix("banner:%s")
	b.SetConfLifeTime(60 * time.Second)
	b.SetFunc(f)

	return b
}
