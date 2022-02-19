package config

import (
	"sync/atomic"
	"time"
)

const defTime = 12 * time.Hour

var defaultRouterConfig atomic.Value

func DefaultRouter() *DefaultRouterConfig {
	if c, ok := defaultRouterConfig.Load().(*DefaultRouterConfig); ok {
		return c
	}

	return &DefaultRouterConfig{}
}

func defaultRouterConfigInit() {
	c := &DefaultRouterConfig{
		&routerConfig{
			host:   "",
			isCors: true, // 默认配置设置了跨域 子配置才能通过
			corsConf: corsConfig{ // 最大限度通过，子配置限制
				AllowOrigins: []string{"*"},
				AllowMethods: []string{"*"},
				AllowHeaders: []string{"*"},
				MaxAge:       defTime,
			},
		},
	}

	defaultRouterConfig.Store(c)
}

var adminRouterConfig atomic.Value

func AdminRouter() *AdminRouterConfig {
	if c, ok := adminRouterConfig.Load().(*AdminRouterConfig); ok {
		return c
	}

	return &AdminRouterConfig{}
}

func adminRouterConfigInit() {
	c := &AdminRouterConfig{
		&routerConfig{
			host:   "",
			isCors: true,
			corsConf: corsConfig{
				AllowOrigins: []string{"*"},
				AllowMethods: []string{"*"},
				AllowHeaders: []string{"*"},
				MaxAge:       defTime,
			},
		},
	}

	adminRouterConfig.Store(c)
}

type corsConfig struct {
	AllowOrigins []string
	AllowMethods []string
	AllowHeaders []string
	MaxAge       time.Duration
}

// RouterConfig 路由配置
type routerConfig struct {
	host     string     // 域名（备用）
	isCors   bool       // 是否跨域
	corsConf corsConfig // 跨域配置
}

func (r *routerConfig) Host() string {
	return r.host
}

func (r *routerConfig) IsCors() bool {
	return r.isCors
}

func (r *routerConfig) CorsConf() corsConfig {
	return r.corsConf
}

// DefaultRouterConfig 默认配置
type DefaultRouterConfig struct {
	*routerConfig
}

// AdminRouterConfig 后台路由配置
type AdminRouterConfig struct {
	*routerConfig
}
