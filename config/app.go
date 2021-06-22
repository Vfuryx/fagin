package config

import (
	"fagin/app/constants/time_format"
	"fagin/pkg/conf"
	"fagin/pkg/paths"
	"reflect"
	"strings"
)

type app struct {
	RootPath       string
	AppPath        string
	PublicPath     string
	StaticPath     string
	StoragePath    string
	MigrationsPath string
	ResourcePath   string

	Name   string
	Env    string
	Port   string
	Key    string
	Debug  bool
	Url    string
	Locale string

	TimeFormat string
	Timezone   string
}

var App app

func init() {
	// 初始化项目路径
	App.RootPath = paths.RootPath
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

	// 默认时间格式
	App.TimeFormat = time_format.Def
	// 默认时区
	App.Timezone = "PRC"
}

func (app app) GetName() string {
	return strings.Split(reflect.TypeOf(app).PkgPath(), "/")[0]
}
