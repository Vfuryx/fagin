package cache

import (
	"encoding/json"
	"fagin/app/models/banner"
	"fagin/pkg/cache"
	"time"
)

// 轮播图缓存管理
type bannerCache struct {
	cache.SCache
	Content func() ([]banner.Banner, error)
}

// 实例
var Banner = NewBanner()

func NewBanner() *bannerCache {
	var b = new(bannerCache)
	b.Prefix = "banner:"
	b.LifeTime = 60 * time.Second
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
//获取数据
func (b *bannerCache) GetContent(id string) (string, error) {
	banners, err := b.Content()
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(banners)
	if err != nil {
		return "", err
	}
	return string(data), nil
}