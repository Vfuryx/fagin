package config

import (
	"fagin/pkg/conf"
)

type app struct {
	Name 	string
	Env  	string
	Port 	string
	Key  	string
	Debug 	bool
	Url		string
}

var App = app{}

func init() {
	App.Name 	= conf.Env("APP_NAME", "app")
	App.Env 	= conf.Env("APP_ENV", "local")
	App.Port 	= conf.Env("APP_PORT", "8080")
	App.Key 	= conf.Env("APP_KEY", "")
	App.Debug 	= conf.EnvBool("APP_DEBUG")
	App.Url 	= conf.Env("APP_URL", "")
}


