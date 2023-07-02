package config

import (
	"fadmin/pkg/conf"
	"sync/atomic"
)

// logConfig 日志配置
var logConfig atomic.Value

// Log 日志配置
func Log() *LogConfig {
	if c, ok := logConfig.Load().(*LogConfig); ok {
		return c
	}

	return &LogConfig{}
}

func logConfigInit() {
	c := &LogConfig{
		def: conf.GetString("log.channel", "default"),
		channels: map[string]Channel{
			"default": {
				Driver: "zap",
				Path:   App().StoragePath() + `/logs/server`,
			},
			"admin": {
				Driver: "zap",
				Path:   App().StoragePath() + `/logs/admin`,
			},
			"api": {
				Driver: "zap",
				Path:   App().StoragePath() + `/logs/api`,
			},
			"sql": {
				Driver: "zap",
				Path:   App().StoragePath() + `/logs/sql`,
			},
		},
	}

	logConfig.Store(c)
}

// Channel 通道
type Channel struct {
	Driver string // zap 或 logrus
	Path   string
}

// LogConfig 日志配置
type LogConfig struct {
	def      string
	channels map[string]Channel
}

func (log *LogConfig) Default() string {
	return log.def
}

func (log *LogConfig) Channels() map[string]Channel {
	return log.channels
}
