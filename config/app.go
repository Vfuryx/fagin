package config

import (
	"fagin/pkg/conf"
	"reflect"
	"strings"
	"time"
)

type app struct {
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
	Url      string
	Locale   string
	Timezone *time.Location
}

var App app

func init() {
	// 初始化项目路径
	App.RootPath = conf.RootPath
	App.AppPath = App.RootPath + "/app"
	App.PublicPath = App.RootPath + "/public"
	App.StaticPath = App.RootPath + "/static"
	App.StoragePath = App.RootPath + "/storage"
	App.MigrationsPath = App.RootPath + "/database/migrations"
	App.ResourcePath = App.RootPath + "/resources"

	App.Name = App.GetName()
	App.Env = conf.GetString("app.env", "local")  // 环境：prod、dev、local
	App.Port = conf.GetString("app.port", "8080") // 端口
	App.Key = conf.GetString("app.key", "")
	App.Url = conf.GetString("app.url", "")
	App.Debug = conf.GetBool("app.debug")
	App.Locale = conf.GetString("app.local", "zh")

	// 默认时区
	App.Timezone = time.FixedZone("CST", 8*3600)
}

func (app app) GetName() string {
	return strings.Split(reflect.TypeOf(app).PkgPath(), "/")[0]
}
