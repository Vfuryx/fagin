package config

import "fagin/pkg/conf"

type SessionConfig struct {
	Secret string // 盐
	Key    string // cookie 的 session 键
}

var sessionConfig = new(SessionConfig)

func Session() SessionConfig {
	return *sessionConfig
}

func (session *SessionConfig) init() {
	session.Secret = conf.GetString("session.secret", "secret")
	session.Key = conf.GetString("session.key", "SESSION_ID")
}
