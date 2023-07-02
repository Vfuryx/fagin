package config

import (
	"fadmin/pkg/conf"
	"reflect"
	"strings"
	"sync/atomic"
	"time"
)

var appConfig atomic.Value

// App App
func App() *AppConfig {
	if c, ok := appConfig.Load().(*AppConfig); ok {
		return c
	}

	return &AppConfig{}
}

func appConfigInit() {
	c := &AppConfig{
		// 初始化项目路径
		rootPath:       conf.RootPath,
		appPath:        conf.RootPath + "/app",
		publicPath:     conf.RootPath + "/public",
		staticPath:     conf.RootPath + "/static",
		storagePath:    conf.RootPath + "/storage",
		migrationsPath: conf.RootPath + "/database/migrations",
		resourcePath:   conf.RootPath + "/resources",

		env:    conf.GetString("app.env", "local"), // 环境：prod、dev、local
		port:   conf.GetString("app.port", "8080"), // 端口
		key:    conf.GetString("app.key", ""),
		debug:  conf.GetBool("app.debug"),
		url:    conf.GetString("app.url", ""),
		locale: conf.GetString("app.local", "zh"),

		// 默认时区 location, _ := time.LoadLocation("Asia/Shanghai")
		// _ = os.Setenv("TZ", location.String())
		timezone: time.FixedZone("CST", 8*3600), //nolint:gomnd // 默认时区
	}

	// 设置名称
	c.name = c.GetName()

	appConfig.Store(c)
}

// AppConfig APP 配置
type AppConfig struct {
	rootPath       string
	appPath        string
	publicPath     string
	staticPath     string
	storagePath    string
	migrationsPath string
	resourcePath   string

	name     string
	env      string
	port     string
	key      string
	debug    bool
	url      string
	locale   string
	timezone *time.Location
}

func (app *AppConfig) RootPath() string {
	return app.rootPath
}

func (app *AppConfig) AppPath() string {
	return app.appPath
}

func (app *AppConfig) PublicPath() string {
	return app.publicPath
}

func (app *AppConfig) StaticPath() string {
	return app.staticPath
}

func (app *AppConfig) StoragePath() string {
	return app.storagePath
}

func (app *AppConfig) MigrationsPath() string {
	return app.migrationsPath
}

func (app *AppConfig) ResourcePath() string {
	return app.resourcePath
}

func (app *AppConfig) Name() string {
	return app.name
}

func (app *AppConfig) Env() string {
	return app.env
}

func (app *AppConfig) Port() string {
	return app.port
}

func (app *AppConfig) Key() string {
	return app.key
}

func (app *AppConfig) Debug() bool {
	return app.debug
}

func (app *AppConfig) URL() string {
	return app.url
}

func (app *AppConfig) Locale() string {
	return app.locale
}

func (app *AppConfig) Timezone() *time.Location {
	return app.timezone
}

// GetName 获取名称
func (app *AppConfig) GetName() string {
	return strings.Split(reflect.TypeOf(*app).PkgPath(), "/")[0]
}
