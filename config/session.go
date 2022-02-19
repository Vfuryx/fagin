package config

import (
	"fagin/pkg/conf"
	"sync/atomic"
)

var sessionConfig atomic.Value

func Session() *SessionConfig {
	if c, ok := sessionConfig.Load().(*SessionConfig); ok {
		return c
	}

	return &SessionConfig{}
}

func sessionConfigInit() {
	c := SessionConfig{
		secret: conf.GetString("session.secret", "secret"),
		key:    conf.GetString("session.key", "SESSION_ID"),
	}

	sessionConfig.Store(c)
}

type SessionConfig struct {
	secret string // 盐
	key    string // cookie 的 session 键
}

func (s *SessionConfig) Secret() string {
	return s.secret
}

func (s *SessionConfig) Key() string {
	return s.key
}
