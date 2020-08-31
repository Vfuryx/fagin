package config

import (
	"fagin/pkg/conf"
	"fagin/pkg/path"
	"reflect"
	"strings"
)

type app struct {
	RootPath       string
	AppPath        string
	PublicPath     string
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
}

var App app

func init() {
	// 初始化项目路径
	App.RootPath = path.RootPath
	App.AppPath = path.RootPath + "/app"
	App.PublicPath = path.RootPath + "/public"
	App.StoragePath = path.RootPath + "/storage"
	App.MigrationsPath = path.RootPath + "/database/migrations"
	App.ResourcePath = path.RootPath + "/resources"

	App.Name = App.GetName()
	App.Env = conf.Env("APP_ENV", "local")  // 环境：prod、dev、local
	App.Port = conf.Env("APP_PORT", "8080") // 端口
	App.Key = conf.Env("APP_KEY", "")
	App.Url = conf.Env("APP_URL", "")
	App.Debug = conf.EnvBool("APP_DEBUG")
	App.Locale = conf.Env("APP_LOCALE", "zh")
}

func (app app) GetName() string {
	return strings.Split(reflect.TypeOf(app).PkgPath(), "/")[0]
}
