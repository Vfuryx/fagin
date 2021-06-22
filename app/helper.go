package app

import (
	"encoding/json"
	"fagin/config"
	"fagin/pkg/logger"
	"fagin/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io"
	"net/http"
	"time"
)

// TimeFormat 默认时间格式
var TimeFormat = config.App.TimeFormat

// IsProd 是否正式环境
func IsProd() bool {
	return config.App.Env == "prod"
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
	return config.CDN.URL + path + "?t=" + time.Now().Format("200612154")
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

// Btoi 布尔转数字
func Btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// GetLocation 根基IP获取地址信息
func GetLocation(ip string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=3fabc36c20379fbb9300c79b19d5d05e")
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	s, err := io.ReadAll(resp.Body)
	m := make(map[string]string)
	err = json.Unmarshal(s, &m)
	if err != nil {
		return "未知位置"
	}
	if m["province"] == "" {
		return "未知位置"
	}
	return m["province"] + "-" + m["city"]
}

// Log 日志模块
func Log(opt ...string) *logrus.Logger {
	if len(opt) <= 0 {
		return logger.Log
	}
	return logger.New(opt[0])
}
