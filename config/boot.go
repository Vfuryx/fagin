package config

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// Init 初始化配置
func Init() {
	appConfig.init()
	logConfig.init()
	sessionConfig.init()
	amqpConfig.init()
	cacheConfig.init()
	databaseConfig.init()
	cdnConfig.init()
	templateConfig.init()
	defaultRouterConfig.init()
	adminRouterConfig.init()
	jwtConfig.init()

	// 监听
	watchConfig()
}

// 监控 config
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("Config file changed:", e.Name)
		// 重新加载配置
		Init()
	})
}
