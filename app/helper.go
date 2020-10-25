package app

import (
	"encoding/json"
	"fagin/app/errno"
	"fagin/config"
	"fagin/pkg/log"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strings"
	"time"
	"unicode"
)
// 默认时间格式
var TimeFormat = config.App.TimeFormat

// 是否正式环境
func IsProd() bool {
	return config.App.Env == "prod"
}

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
}

func JsonResponse(ctx *gin.Context, err error, data interface{}) {
	code, msg := errno.DecodeErr(err)
	ctx.JSON(http.StatusOK, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

func JsonResponseWithStatus(ctx *gin.Context, statusCode int, err error, data interface{}) {
	code, msg := errno.DecodeErr(err)
	ctx.JSON(statusCode, Response{
		Code:    code,
		Message: msg,
		Data:    data,
	})
}

// 驼峰式写法转为下划线写法
//将大写转成 _a 的形式
func Underscore(str string) string {
	var buf strings.Builder
	for index, value := range str {
		// value >= 'A' && value <= 'Z'
		if unicode.IsUpper(value) {
			if index != 0 {
				buf.WriteByte('_')
			}
			// byte(value + 'a' - 'A')
			buf.WriteRune(unicode.ToLower(value))
			continue
		}
		buf.WriteRune(value)
	}
	return buf.String()
}

// 下划线写法转为驼峰写法
func Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// 首字母大写
func UFirst(str string) string {
	for _, v := range str {
		return string(unicode.ToUpper(v)) + str[+1:]
	}
	return ""
}

// 首字母小写
func LFirst(str string) string {
	for _, v := range str {
		return string(unicode.ToLower(v)) + str[+1:]
	}
	return ""
}

// 每分钟回源一次，一分钟更新一次
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

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

// 加密密码
// Encrypt encrypts the plain text with bcrypt.
func Encrypt(source string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(source), bcrypt.DefaultCost)
	return string(hashedBytes), err
}

// 匹配密码
// Compare compares the encrypted text with the plain text if it's the same.
func Compare(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// 布尔转数字
func btoi(b bool) int {
	if b {
		return 1
	}
	return 0
}

// 根基IP获取地址信息
func GetLocation(ip string) string {
	if ip == "127.0.0.1" || ip == "localhost" {
		return "内部IP"
	}
	resp, err := http.Get("https://restapi.amap.com/v3/ip?ip=" + ip + "&key=3fabc36c20379fbb9300c79b19d5d05e")
	if err != nil {
		panic(err)

	}
	defer resp.Body.Close()
	s, err := ioutil.ReadAll(resp.Body)

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

// 日志模块
func Log(opt ...string) *logrus.Logger {
	if len(opt) <= 0 {
		return log.Log
	}
	return log.New(opt[0])
}