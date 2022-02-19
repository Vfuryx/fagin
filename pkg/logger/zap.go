package logger

import (
	"fagin/config"
	"io"
	"os"
	"time"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

type zapLog struct {
	log    *zap.SugaredLogger
	writer io.Writer
}

var _ Logger = &zapLog{}

// DriverZap zap 引擎
const DriverZap = "zap"

func init() {
	driverMap[DriverZap] = NewZap
}

// NewZap 实例化
func NewZap(name string) Logger {
	path := config.Log().Channels()[name].Path

	w := getLogWriter(path + "/log.log")

	core := zapcore.NewCore(
		getEncoder(),       // 编码器配置
		w,                  // 指定日志将写到哪里去
		zapcore.DebugLevel, // 日志级别
	)

	// 构造日志
	l := zap.New(
		core,
		// zap.Development(), // 开启文件及行号
		// zap.AddStacktrace(zap.ErrorLevel),
	)

	return &zapLog{
		log:    l.Sugar(),
		writer: w,
	}
}

// 获取编码器
func getEncoder() zapcore.Encoder {
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = timeEncoder

	return zapcore.NewConsoleEncoder(encoderConfig)
}

// 打印到文件
func getLogWriter(path string) zapcore.WriteSyncer {
	lumberJackLogger := &lumberjack.Logger{
		Filename:   path, // 日志文件的位置
		MaxSize:    10,   //nolint:gomnd // 在进行切割之前，日志文件的最大大小（以MB为单位）
		MaxBackups: 365,  //nolint:gomnd // 保留旧文件的最大个数
		MaxAge:     365,  //nolint:gomnd // 保留旧文件的最大天数
		Compress:   true, // 是否压缩/归档旧文件
	}

	return zapcore.AddSync(lumberJackLogger)
}

//nolint:deadcode,unused // 打印到控制台
func stdout() zapcore.WriteSyncer {
	return zapcore.AddSync(os.Stdout)
}

// timeEncoder 自定义时间编码器
func timeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	zapcore.TimeEncoderOfLayout("2006-01-02 15:04:05")(t, enc)
}

func (z *zapLog) Debug(args ...interface{}) {
	defer func() {
		_ = z.log.Sync()
	}()
	z.log.Debug(args...)
}

func (z *zapLog) Info(args ...interface{}) {
	defer func() {
		_ = z.log.Sync()
	}()
	z.log.Info(args...)
}

func (z *zapLog) Warn(args ...interface{}) {
	defer func() {
		_ = z.log.Sync()
	}()
	z.log.Warn(args...)
}

func (z *zapLog) Error(args ...interface{}) {
	defer func() {
		_ = z.log.Sync()
	}()
	z.log.Error(args...)
}

func (z *zapLog) Debugf(format string, args ...interface{}) {
	defer func() {
		_ = z.log.Sync()
	}()
	z.log.Debugf(format, args...)
}

func (z *zapLog) Infof(format string, args ...interface{}) {
	defer func() {
		_ = z.log.Sync()
	}()
	z.log.Infof(format, args...)
}

func (z *zapLog) Warnf(format string, args ...interface{}) {
	defer func() {
		_ = z.log.Sync()
	}()
	z.log.Warnf(format, args...)
}

func (z *zapLog) Errorf(format string, args ...interface{}) {
	defer func() {
		_ = z.log.Sync()
	}()
	z.log.Errorf(format, args...)
}

func (z *zapLog) Writer() io.Writer {
	return z.writer
}
