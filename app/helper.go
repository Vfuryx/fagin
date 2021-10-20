package app

import (
	"fagin/app/enums"
	"fagin/config"
	"fagin/pkg/logger"
	"fagin/pkg/response"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

// IsProd 是否正式环境
func IsProd() bool {
	return config.App().Env == "prod"
}

func ResponseJsonOK(ctx *gin.Context, data interface{}) {
	response.JsonOK(ctx, data)
}
func ResponseJsonErr(ctx *gin.Context, err error, errors interface{}) {
	response.JsonErr(ctx, err, errors)
}

func ResponseJsonWithStatus(ctx *gin.Context, statusCode int, err error, data interface{}, errors interface{}) {
	response.JsonWithStatus(ctx, statusCode, err, data, errors)
}

func ResponseJson(ctx *gin.Context, err error, data interface{}, errors interface{}, statusCode int) {
	response.Json(ctx, err, data, errors, statusCode)
}

// WebAsset 每分钟回源一次，一分钟更新一次
func WebAsset(path string) string {
	// 是否正式环境
	if ok := IsProd(); !ok {
		return path
	}
	return config.CDN().URL + path + "?t=" + time.Now().Format("200612154")
}

func View(ctx *gin.Context, path string, obj interface{}) {
	ctx.HTML(http.StatusOK, path, obj)
}

func ViewWithStatus(ctx *gin.Context, statusCode int, path string, obj interface{}) {
	ctx.HTML(statusCode, path, obj)
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
	if len(opt) <= 0 {
		return logger.Log()
	}
	return logger.Channel(opt[0])
}

// Now 获取当前时间
func Now() time.Time {
	return time.Now().In(config.App().Timezone)
}

// TimeToStr 时间转字符串
func TimeToStr(t time.Time) string {
	return t.In(config.App().Timezone).Format(enums.TimeFormatDef.Get())
}

// StrToTime 时间转字符串
func StrToTime(value string) (time.Time, error) {
	return time.ParseInLocation(enums.TimeFormatDef.Get(), value, config.App().Timezone)
}
