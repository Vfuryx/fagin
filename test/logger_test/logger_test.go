package logger_test

import (
	"fagin/config"
	"fagin/pkg/logger"
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	// 初始化配置
	config.Init()
	logger.Init()
	m.Run()

	os.Exit(1)
}

func BenchmarkLogger(b *testing.B) {
	for i := 0; i < b.N; i++ {
		go logger.Channel(logger.AdminMode).Info(123)
	}
}
