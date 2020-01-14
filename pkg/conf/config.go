package conf

import (
	"fagin/pkg/path"
	"github.com/fsnotify/fsnotify"
	"github.com/joho/godotenv"
	"github.com/spf13/viper"
)

func init() {
	envLoad()
	configLoad()
	watchConfig()
}

// 读取 env
func envLoad() {
	// 加载 .env 文件的配置
	err := godotenv.Load(path.RootPath + "/.env")
	if err != nil {
		panic(err)
	}

}

// 读取 config
func configLoad() {
	viper.AutomaticEnv()

	viper.SetConfigName("config")
	viper.AddConfigPath(path.RootPath)
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

// 获取 env 配置，转换为 string 类型并返回
func Env(name string, def string) string {
	v := viper.GetString(name)
	if v == "" {
		return def
	}
	return v
}

// 获取 env 配置，转换为 bool 类型并返回
// def false
func EnvBool(name string) bool {
	return viper.GetBool(name)
}

// 获取 env 配置，转换为 int 类型并返回
func EnvInt(name string, def int) int {
	v := viper.GetInt(name)
	if v == 0 {
		return def
	}
	return v
}

// 获取 env 配置，转换为 float64 类型并返回
func EnvFloat64(name string, def float64) float64 {
	v := viper.GetFloat64(name)
	if v == 0 {
		return def
	}
	return v
}

func GetStringMapStringSlice(key string, def map[string][]string) map[string][]string {
	v := viper.GetStringMapStringSlice("casbin.model")
	if len(v) == 0 {
		return def
	}
	return v
}
