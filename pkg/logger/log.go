package logger

import (
	"errors"
	"fagin/config"
	"io"
	"os"
	"sync"
)

// Logger 日志接口
type Logger interface {
	Writer() io.Writer

	Debug(args ...interface{})
	Info(args ...interface{})
	Warn(args ...interface{})
	Error(args ...interface{})

	Debugf(format string, args ...interface{})
	Infof(format string, args ...interface{})
	Warnf(format string, args ...interface{})
	Errorf(format string, args ...interface{})
}

// DefaultMode 默认
const DefaultMode = "default"

// AdminMode 后台
const AdminMode = "admin"

// APIMode API
const APIMode = "api"

var ErrChannelNotFound = errors.New("channel not found")
var ErrDriverNotFound = errors.New("driver not found")

var driverMap = map[string]func(name string) Logger{}

// DefaultLog 默认日志
var defaultLog Logger

// Init 初始化
func Init() {
	channel, ok := config.Log().Channels()[config.Log().Default()]

	// 创建文件夹 可以多层
	err := os.MkdirAll(channel.Path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	if !ok {
		panic(ErrChannelNotFound)
	}

	driver, ok := driverMap[channel.Driver]
	if !ok {
		panic(ErrDriverNotFound)
	}

	defaultLog = driver(config.Log().Default())
}

// Log 默认日志
func Log() Logger {
	return defaultLog
}

// 并发获取 channel
var channelMap sync.Map

// Channel 通道
func Channel(name string) Logger {
	if l, ok := channelMap.Load(name); ok {
		return l.(Logger)
	}

	channel := config.Log().Channels()[name]

	var p, _ = channelMap.LoadOrStore(name, driverMap[channel.Driver](name))

	return p.(Logger)
}
