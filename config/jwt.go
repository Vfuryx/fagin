package config

import (
	"fagin/pkg/conf"
)

type JWTConfig struct {
	Secret string
}

var jwtConfig = new(JWTConfig)

func JWT() JWTConfig {
	return *jwtConfig
}

func (jwt *JWTConfig) init() {
	jwt.Secret = conf.GetString("jwt.secret", "")
}
