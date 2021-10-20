package config

import (
	"fagin/pkg/conf"
)

// CDNConfig 配置
type CDNConfig struct {
	URL string // 链接
}

var cdnConfig = new(CDNConfig)

// CDN CDN配置
func CDN() CDNConfig {
	return *cdnConfig
}

func (cdn *CDNConfig) init() {
	cdn.URL = conf.GetString("cdn.host", "")
}
