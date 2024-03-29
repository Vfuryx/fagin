package conf

import (
	"fadmin/pkg/paths"
	"fmt"
	"os"
	"path"
	"path/filepath"

	"github.com/spf13/viper"
)

var RootPath string

func init() {
	configLoad()
}

// 读取 config
func configLoad() {
	// app目录
	appPath, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		panic(err)
	}
	// 工作目录
	workPath, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	var pathFiles = []string{
		workPath,
		path.Clean(workPath + "/.."),
		path.Clean(workPath + "/../.."),
		appPath,
	}

	var (
		filename   = ".config"
		configType = "yml"
		l          = len(pathFiles) - 1
		i          int
	)

	for i, workPath = range pathFiles {
		appConfigPath := filepath.Join(workPath, filename+"."+configType)
		if paths.FileExists(appConfigPath) {
			break
		}

		if i == l {
			fmt.Printf("ERROR：配置文件 .config.yml 读取失败\n %v", err)
		}
	}

	// 设置项目根目录
	RootPath = workPath

	// viper.AutomaticEnv()
	viper.SetConfigName(filename)
	viper.AddConfigPath(workPath)
	viper.SetConfigType(configType)

	err = viper.ReadInConfig()
	if err != nil {
		panic(err)
	}
}

// GetString 获取 string 类型并返回
func GetString(name, def string) string {
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

// GetInt64 获取 int64 类型并返回
func GetInt64(name string, def int64) int64 {
	v := viper.GetInt64(name)
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

func GetStringMap(key string, def map[string]interface{}) map[string]interface{} {
	v := viper.GetStringMap(key)
	if len(v) == 0 {
		return def
	}

	return v
}

func GetStringSlice(key string, def []string) []string {
	v := viper.GetStringSlice(key)
	if len(v) == 0 {
		return def
	}

	return v
}

// Get 获取配置
func Get(key string, def interface{}) interface{} {
	v := viper.Get(key)
	if v == nil {
		return def
	}

	return v
}

// UnmarshalKey 解组
func UnmarshalKey(key string, v interface{}) error {
	return viper.UnmarshalKey(key, v)
}
