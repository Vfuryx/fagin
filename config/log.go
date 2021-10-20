package config

import "fagin/pkg/conf"

// Channel 通道
type Channel struct {
	Driver string // zap 或 logrus
	Path   string
}

// LogConfig 日志配置
type LogConfig struct {
	Default  string
	Channels map[string]Channel
}

// logConfig 日志配置
var logConfig = new(LogConfig)

// Log 日志配置
func Log() LogConfig {
	return *logConfig
}

func (log *LogConfig) init() {
	log.Default = conf.GetString("log.channel", "default")
	log.Channels = map[string]Channel{
		"default": {
			Driver: "zap",
			Path:   App().StoragePath + `/logs/server`,
		},
		"admin": {
			Driver: "zap",
			Path:   App().StoragePath + `/logs/admin`,
		},
		"api": {
			Driver: "zap",
			Path:   App().StoragePath + `/logs/api`,
		},
	}
}
