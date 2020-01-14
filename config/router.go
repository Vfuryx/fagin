package config

import "time"

type corsConfig struct {
	AllowOrigins []string
	AllowMethods []string
	AllowHeaders []string
	MaxAge       time.Duration
}

// 路由配置
type router struct {
	Host     string     // 域名（备用）
	IsCors   bool       // 是否跨域
	CorsConf corsConfig // 跨域配置
}

// 默认配置
var DefaultRouter router

// 后台路由配置
var AdminRouter router

func init() {
	DefaultRouter.Host = ""
	DefaultRouter.IsCors = true          // 默认配置设置了跨域 子配置才能通过
	DefaultRouter.CorsConf = corsConfig{ // 最大限度通过，子配置限制
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		MaxAge:       12 * time.Hour,
	}

	AdminRouter.Host = ""
	AdminRouter.IsCors = true
	AdminRouter.CorsConf = corsConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
		AllowHeaders: []string{"*"},
		MaxAge:       12 * time.Hour,
	}
}
