package config

import (
	"fmt"
	"github.com/astaxie/beego/logs"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Config struct {
	Name string
	Path string
	Type string
}

func init() {
	c := Config{
		Name: "env",
		Path: "$GOPATH/src/fagin/",
		Type: "yml",
	}

	// 初始化配置文件
	if err := c.initConfig(); err != nil{
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	// 监控配置文件并热加载程序
	c.watchConfig()

}

func (c *Config) initConfig() error {
	viper.SetConfigName(c.Name)
	viper.AddConfigPath(c.Path)
	viper.SetConfigType(c.Type)
	return viper.ReadInConfig()
}


func Get(str string) interface{} {
	return viper.Get(str)
}

func (c *Config) watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		logs.Info("Config file changed: %s", e.Name)
	})
}
