package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	Init()
}

// Init 初始化配置
func Init() {
	appConfigInit()
	databaseConfigInit()
	templateConfigInit()
	logConfigInit()
	sessionConfigInit()
	amqpConfigInit()
	cacheConfigInit()
	cndConfigInit()

	defaultRouterConfigInit()
	adminRouterConfigInit()
	jwtConfigInit()

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
