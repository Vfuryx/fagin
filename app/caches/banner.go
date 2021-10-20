package caches

import (
	"fagin/pkg/cache"
	"time"
)

// NewBanner 轮播图缓存管理
func NewBanner(f cache.GetterFunc) *cache.Cache {
	var b = new(cache.Cache)
	b.SetKeyFormat("banner:%s")
	b.SetLifeTime(60 * time.Second)
	b.SetFunc(f)

	return b
}
