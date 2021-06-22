package config

import "fagin/pkg/conf"

type session struct {
	Secret string // 盐
	Key    string // cookie 的 session 键
}

var Session session

func init() {
	Session.Secret = conf.GetString("session.secret", "secret")
	Session.Key = conf.GetString("session.key", "SESSION_ID")
}
