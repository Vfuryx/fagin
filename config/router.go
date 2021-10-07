package config

import "time"

type corsConfig struct {
	AllowOrigins []string
	AllowMethods []string
	AllowHeaders []string
	MaxAge       time.Duration
}

// RouterConfig 路由配置
type RouterConfig struct {
	Host     string     // 域名（备用）
	IsCors   bool       // 是否跨域
	CorsConf corsConfig // 跨域配置
}

// DefaultRouterConfig 默认配置
type DefaultRouterConfig struct {
	RouterConfig
}

var defaultRouterConfig = new(DefaultRouterConfig)

func DefaultRouter() DefaultRouterConfig {
	return *defaultRouterConfig
}

// AdminRouterConfig 后台路由配置
type AdminRouterConfig struct {
	RouterConfig
}

var adminRouterConfig = new(AdminRouterConfig)

func AdminRouter() AdminRouterConfig {
	return *adminRouterConfig
}

func (defaultRouter *DefaultRouterConfig) init() {
	defaultRouter.Host = ""
	defaultRouter.IsCors = true          // 默认配置设置了跨域 子配置才能通过
	defaultRouter.CorsConf = corsConfig{ // 最大限度通过，子配置限制
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		MaxAge:       12 * time.Hour,
	}

}

func (adminRouter *AdminRouterConfig) init() {
	adminRouter.Host = ""
	adminRouter.IsCors = true
	adminRouter.CorsConf = corsConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		MaxAge:       12 * time.Hour,
	}
}
