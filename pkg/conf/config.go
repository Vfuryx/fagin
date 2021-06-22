package conf

import (
	"fagin/pkg/paths"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func init() {
	configLoad()
	watchConfig()
}

// 读取 config
func configLoad() {
	//viper.AutomaticEnv()

	viper.SetConfigName(".config")
	viper.AddConfigPath(paths.RootPath)
	viper.SetConfigType("yml")

	err := viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// 监控 config
func watchConfig() {
	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {

	})
}

// GetString 获取 string 类型并返回
func GetString(name string, def string) string {
	v := viper.GetString(name)
	if v == "" {
		return def
	}
	return v
}

// GetBool 获取 bool 类型并返回 默认 false
func GetBool(name string) bool {
	return viper.GetBool(name)
}

// GetInt 获取 int 类型并返回
func GetInt(name string, def int) int {
	v := viper.GetInt(name)
	if v == 0 {
		return def
	}
	return v
}

// GetFloat64 获取 float64 类型并返回
func GetFloat64(name string, def float64) float64 {
	v := viper.GetFloat64(name)
	if v == 0 {
		return def
	}
	return v
}

func GetStringMapStringSlice(key string, def map[string][]string) map[string][]string {
	v := viper.GetStringMapStringSlice(key)
	if len(v) == 0 {
		return def
	}
	return v
}
