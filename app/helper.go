package app

import (
	"fadmin/app/enums"
	"fadmin/config"
	"fadmin/pkg/logger"
	"golang.org/x/crypto/bcrypt"
	"time"
)

// IsProd 是否正式环境
func IsProd() bool {
	return config.App().Env() == "prod"
}

// Encrypt 加密密码
// Encrypt encrypts the plain text with bcrypt.
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// Compare 匹配密码
// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// Log 日志模块
func Log(opt ...string) logger.Logger {
	if len(opt) == 0 {
		return logger.Log()
	}

	return logger.Channel(opt[0])
}

// Now 获取当前时间
func Now() time.Time {
	return time.Now().In(config.App().Timezone())
}

// TimeToStr 时间转字符串
func TimeToStr(t time.Time) string {
	return t.In(config.App().Timezone()).Format(enums.TimeFormatDef.Get())
}

// StrToTime 时间转字符串
func StrToTime(value string) (time.Time, error) {
	return time.ParseInLocation(enums.TimeFormatDef.Get(), value, config.App().Timezone())
}
