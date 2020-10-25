package cache

import (
	"encoding/json"
	"fagin/app/models/website_config"
	"fagin/pkg/cache"
	"time"
)

// 网站配置缓存管理
type websiteConfig struct {
	cache.SCache
	Content func() (*website_config.WebsiteConfig, error)
}


// 实例
var WebsiteConfig = NewWebsiteConfig()

func NewWebsiteConfig() *websiteConfig {
	var c = new(websiteConfig)
	c.Prefix = "website:config:"
	c.LifeTime = 60 * time.Second
	c.SetFunc(c)
	return c
}

// 获取键名称
func (c *websiteConfig) Key(value string) string {
	return c.Prefix + value
}

// 默认存在时间
func (c *websiteConfig) Lift() time.Duration {
	return c.LifeTime
}

//获取数据
func (c *websiteConfig) GetContent(id string) (string, error) {
	str, err := c.Content()
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
