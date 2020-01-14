package config

import "fagin/pkg/conf"

type session struct {
	Secret string // 盐
	Key    string // cookie 的 session 键
}

var Session session

func init() {
	Session.Secret = conf.Env("SESSION_SECRET", "secret")
	Session.Key = conf.Env("SESSION_KEY", "SESSION_ID")
}
