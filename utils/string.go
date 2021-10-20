package utils

import (
	"math/rand"
	"strconv"
	"strings"
	"time"
	"unicode"
)

// Camel 下划线写法转为驼峰写法
func Camel(name string) string {
	name = strings.Replace(name, "_", " ", -1)
	name = strings.Title(name)
	return strings.Replace(name, " ", "", -1)
}

// Underscore 驼峰式写法转为下划线写法
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

// UFirst 首字母大写
func UFirst(str string) string {
	for _, v := range str {
		return string(unicode.ToUpper(v)) + str[+1:]
	}
	return ""
}

// LFirst 首字母小写
func LFirst(str string) string {
	for _, v := range str {
		return string(unicode.ToLower(v)) + str[+1:]
	}
	return ""
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// RandString 随机
func RandString(n int) string {
	rand.Seed(time.Now().UnixNano())
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func StrToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, 10, 64)
}
