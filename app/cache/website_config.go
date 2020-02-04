package cache

import (
	"encoding/json"
	"fagin/app/errno"
	"fagin/app/models/website_config"
	"fagin/pkg/cache"
	"time"
)

type websiteConfig struct {
	Prefix string
}

func (websiteConfig) cache() (cache.ICache, error) {
	if cache.Redis == nil {
		return nil, errno.Api.ErrCacheIsClose
	}
	return cache.Redis, nil
}

// 实例
var WebsiteConfig = NewWebsiteConfig()

func NewWebsiteConfig() *websiteConfig {
	return &websiteConfig{Prefix: "website:config:"}
}

func (b *websiteConfig) infoKey() string {
	return cache.Redis.Prefix + b.Prefix + "info"
}

// 是否存在
func (b *websiteConfig) HasInfo() (int64, error) {
	c, err := b.cache()
	if err != nil {
		return 0, err
	}
	return c.Exists(b.infoKey())
}

// 获取
func (b *websiteConfig) GetInfo() (*website_config.WebsiteConfig, error) {
	c, err := b.cache()
	if err != nil {
		return nil, err
	}
	h, err := b.HasInfo()
	if err != nil {
		return nil, err
	}
	if h <= 0 {
		return nil, errno.Api.ErrCacheIsNil
	}
	v, err := c.Get(b.infoKey())
	if err != nil {
		return nil, err
	}
	var data website_config.WebsiteConfig
	err = json.Unmarshal([]byte(v), &data)
	return &data, err
}

// 设置
func (b *websiteConfig) SetInfo(value *website_config.WebsiteConfig, expiration time.Duration) (string, error) {
	c, err := b.cache()
	if err != nil {
		return "", err
	}

	data, err := json.Marshal(value)
	if err != nil {
		return "", err
	}

	return c.Set(b.infoKey(), data, expiration)
}

// 删除
func (b *websiteConfig) DelInfo() (int64, error) {
	c, err := b.cache()
	if err != nil {
		return 0, err
	}
	return c.Remove(b.infoKey())
}
