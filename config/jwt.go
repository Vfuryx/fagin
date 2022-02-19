package config

import (
	"fagin/pkg/conf"
	"sync/atomic"
)

var jwtConfig atomic.Value

func JWT() *JWTConfig {
	if c, ok := jwtConfig.Load().(*JWTConfig); ok {
		return c
	}

	return &JWTConfig{}
}

func jwtConfigInit() {
	c := &JWTConfig{
		secret: conf.GetString("jwt.secret", ""),
	}

	jwtConfig.Store(c)
}

type JWTConfig struct {
	secret string
}

func (jwt *JWTConfig) Secret() string {
	return jwt.secret
}
