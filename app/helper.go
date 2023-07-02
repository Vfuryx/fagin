package app

import (
	"fadmin/config"
)

// IsProd 是否正式环境
func IsProd() bool {
	return config.App().Env() == "prod"
}
