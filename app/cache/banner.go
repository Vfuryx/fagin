package cache

import (
	"encoding/json"
	"fagin/app/errno"
	"fagin/app/models/banner"
	"fagin/pkg/cache"
	"time"
)

// 轮播图缓存管理
type bannerCache struct {
	Prefix string
}

func (bannerCache) cache() (cache.ICache, error) {
	if cache.Redis == nil {
		return nil, errno.ErrCacheIsClose
	}
	return cache.Redis, nil
}

// 实例
var Banner = NewBanner()

func NewBanner() *bannerCache {
	return &bannerCache{Prefix: "banner:"}
}

func (b *bannerCache) listKey() string {
	return cache.Redis.Prefix + b.Prefix + "list"
}

// 是否存在列表
func (b *bannerCache) HasList() (int64, error) {
	c, err := b.cache()
	if err != nil {
		return 0, err
	}
	return c.Exists(b.listKey())
}

// 获取列表
func (b *bannerCache) GetList() ([]banner.Banner, error) {
	c, err := b.cache()
	if err != nil {
		return nil, err
	}

	h, err := b.HasList()
	if err != nil {
		return nil, err
	}
	if h <= 0 {
		return nil, errno.ErrCacheIsNil
	}

	v, err := c.Get(b.listKey())
	if err != nil {
		return nil, err
	}
	var data []banner.Banner
	err = json.Unmarshal([]byte(v), &data)

	return data, err
}

// 设置列表
func (b *bannerCache) SetList(value []banner.Banner, expiration time.Duration) (string, error) {
	c, err := b.cache()
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return c.Set(b.listKey(), data, expiration)
}
