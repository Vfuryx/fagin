package log

import (
	"fagin/config"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log *logrus.Logger

func init() {
	Log = logrus.New()

	// 禁止 logrus 的输出
	file, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	Log.SetOutput(file)

	// 设置LfHook
	lfHook, err := NewLfHook(config.App.StoragePath+`/logs/server.log.%Y-%m-%d.log`, config.App.RootPath+`/server.log`)
	if err != nil {
		panic(err)
	}
	Log.AddHook(lfHook)
}

// path 日志文件路径
// linkName 软连接路径
func NewLfHook(path, linkName string) (*lfshook.LfsHook, error) {
	logWriter, err := rotatelogs.New(
		path,
		rotatelogs.WithLinkName(linkName), // 生成软链，指向最新日志文件
		rotatelogs.WithMaxAge(7*24*time.Hour),     // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		return nil, err
	}

	// 不同等级可以配置不同的写入方式
	// 一下用的是同一种方式
	writeMap := lfshook.WriterMap{
		logrus.InfoLevel:  logWriter,
		logrus.FatalLevel: logWriter,
		logrus.ErrorLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &Formatter{
		TimestampFormat: "2006-01-02 15:04:05",
		NoColors:        true,
		TrimMessages:    false,
	})

	return lfHook, nil
}
