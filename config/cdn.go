package config

import (
	"fagin/pkg/conf"
	"sync/atomic"
)

var cdnConfig atomic.Value

// CDN CDN配置
func CDN() *CDNConfig {
	if c, ok := cdnConfig.Load().(*CDNConfig); ok {
		return c
	}

	return &CDNConfig{}
}

func cndConfigInit() {
	c := &CDNConfig{
		url: conf.GetString("cdn.host", ""),
	}

	cdnConfig.Store(c)
}

// CDNConfig 配置
type CDNConfig struct {
	url string // 链接
}

func (c *CDNConfig) URL() string {
	return c.url
}
