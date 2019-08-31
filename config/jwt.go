package config

import (
	"fagin/pkg/conf"
)

type jwt struct {
	Secret string
}

var Jwt= jwt{}

func init()  {
	Jwt.Secret = conf.Env("JWT_SECRET", "")
}