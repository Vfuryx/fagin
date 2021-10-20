package config

import (
	"fagin/pkg/conf"
	"reflect"
	"strings"
	"time"
)

// AppConfig APP 配置
type AppConfig struct {
	RootPath       string
	AppPath        string
	PublicPath     string
	StaticPath     string
	StoragePath    string
	MigrationsPath string
	ResourcePath   string

	Name     string
	Env      string
	Port     string
	Key      string
	Debug    bool
	URL      string
	Locale   string
	Timezone *time.Location
}

var appConfig = new(AppConfig)

// App App
func App() AppConfig {
	return *appConfig
}

func (app *AppConfig) init() {
	// 初始化项目路径
	app.RootPath = conf.RootPath
	app.AppPath = app.RootPath + "/app"
	app.PublicPath = app.RootPath + "/public"
	app.StaticPath = app.RootPath + "/static"
	app.StoragePath = app.RootPath + "/storage"
	app.MigrationsPath = app.RootPath + "/database/migrations"
	app.ResourcePath = app.RootPath + "/resources"

	app.Name = app.GetName()
	app.Env = conf.GetString("app.env", "local")  // 环境：prod、dev、local
	app.Port = conf.GetString("app.port", "8080") // 端口
	app.Key = conf.GetString("app.key", "")
	app.URL = conf.GetString("app.url", "")
	app.Debug = conf.GetBool("app.debug")
	app.Locale = conf.GetString("app.local", "zh")

	// 默认时区
	app.Timezone = time.FixedZone("CST", 8*3600)
	//location, _ := time.LoadLocation("Asia/Shanghai")
	//_ = os.Setenv("TZ", location.String())
}

// GetName 获取名称
func (app AppConfig) GetName() string {
	return strings.Split(reflect.TypeOf(app).PkgPath(), "/")[0]
}
