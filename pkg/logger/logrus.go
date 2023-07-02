package logger

import (
	"fadmin/config"
	"io"
	"os"
	"time"

	rotatelogs "github.com/lestrrat/go-file-rotatelogs"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

type logrusLog struct {
	log *logrus.Logger
}

var _ Logger = &logrusLog{}

// DriverLogrus logrus 引擎
const DriverLogrus = "logrus"

func init() {
	driverMap[DriverLogrus] = NewLogrus
}

// NewLogrus 实例化
func NewLogrus(name string) Logger {
	log := logrus.New()
	// 禁止 logrus 的输出
	file, err := os.OpenFile(os.DevNull, os.O_APPEND|os.O_WRONLY, os.ModeAppend)
	if err != nil {
		panic(err)
	}

	log.SetOutput(file)

	path := config.Log().Channels()[name].Path

	// 创建文件夹 可以多层
	err = os.MkdirAll(path, os.ModePerm)
	if err != nil {
		panic(err)
	}

	// 设置LfHook
	var lfHook *lfshook.LfsHook

	lfHook, err = NewLfHook(path + "/%Y-%m-%d.log")
	if err != nil {
		panic(err)
	}

	log.AddHook(lfHook)

	return &logrusLog{log: log}
}

// NewLfHook path 日志文件路径
// linkName 软连接路径
func NewLfHook(path string) (*lfshook.LfsHook, error) {
	logWriter, err := rotatelogs.New(
		path,
		rotatelogs.WithMaxAge(365*24*time.Hour),   // 文件最大保存时间
		rotatelogs.WithRotationTime(24*time.Hour), //nolint:gomnd // 日志切割时间间隔
	)
	if err != nil {
		return nil, err
	}

	// 不同等级可以配置不同的写入方式 一下用的是同一种方式
	writeMap := lfshook.WriterMap{
		logrus.DebugLevel: logWriter,
		logrus.InfoLevel:  logWriter,
		logrus.WarnLevel:  logWriter,
		logrus.ErrorLevel: logWriter,
		// logrus.FatalLevel: logWriter,
		// logrus.PanicLevel: logWriter,
	}

	lfHook := lfshook.NewHook(writeMap, &Formatter{
		TimestampFormat: "2006-01-02 15:04:05", // 固定格式
		NoColors:        true,
		TrimMessages:    false,
	})

	return lfHook, nil
}

func (l *logrusLog) Debug(args ...interface{}) {
	l.log.Debug(args...)
}

func (l *logrusLog) Info(args ...interface{}) {
	l.log.Info(args...)
}

func (l *logrusLog) Warn(args ...interface{}) {
	l.log.Warn(args...)
}

func (l *logrusLog) Error(args ...interface{}) {
	l.log.Error(args...)
}

func (l *logrusLog) Debugf(format string, args ...interface{}) {
	l.log.Debugf(format, args...)
}

func (l *logrusLog) Infof(format string, args ...interface{}) {
	l.log.Infof(format, args...)
}

func (l *logrusLog) Warnf(format string, args ...interface{}) {
	l.log.Warnf(format, args...)
}

func (l *logrusLog) Errorf(format string, args ...interface{}) {
	l.log.Errorf(format, args...)
}

func (l *logrusLog) Writer() io.Writer {
	return l.log.Writer()
}
