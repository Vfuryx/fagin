package caches

import (
	"fagin/pkg/cache"
	"time"
)

// 轮播图缓存管理
type bannerCache struct {
	cache.SCache
}

func NewBanner(f cache.GetterFunc) *bannerCache {
	var b = new(bannerCache)
	b.Prefix = "banner:"
	b.LifeTime = 60 * time.Second
	b.Content = f
	b.SetFunc(b)
	return b
}

// 获取键名称
func (b *bannerCache) Key(value string) string {
	return b.Prefix + "list"
}
// 默认存在时间
func (b *bannerCache) Lift() time.Duration {
	return b.LifeTime
}