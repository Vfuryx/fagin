package logger

import (
	"fagin/app/enums"
	"fagin/config"
	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

const AdminModel = "admin"

var Log *logrus.Logger = New("")

func New(name string) *logrus.Logger {
	log := logrus.New()
	// 禁止 logrus 的输出
	file, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}
	log.SetOutput(file)

	// 设置LfHook
	var lfHook *lfshook.LfsHook
	var path = config.App.StoragePath + `/logs/server` // 默认路径
	if name != "" {                                    // 重置路径
		path = config.App.StoragePath + "/logs/" + name
	}
	// 创建文件夹 可以多层
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	lfHook, err = NewLfHook(path + "/%Y-%m-%d.log")
	if err != nil {
		panic(err)
	}
	log.AddHook(lfHook)
	return log
}

// NewLfHook path 日志文件路径
// linkName 软连接路径
func NewLfHook(path string) (*lfshook.LfsHook, error) {
	logWriter, err := rotatelogs.New(
		path,
		rotatelogs.WithMaxAge(365*24*time.Hour),   // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), // 日志切割时间间隔
	)
	if err != nil {
		return nil, err
	}

	// 不同等级可以配置不同的写入方式
	// 一下用的是同一种方式
	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		//logrus.FatalLevel: logWriter,
		//logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &Formatter{
		TimestampFormat: enums.TimeFormatDef.Get(),
		NoColors:        true,
		TrimMessages:    false,
	})

	return lfHook, nil
}
