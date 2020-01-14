package config

import "fagin/pkg/conf"

// cdn 配置
type cdn struct {
	URL string // 链接
}

var CDN cdn

func init() {
	CDN.URL = conf.Env("CDN_URL", "")
}
