package utils

import (
	"crypto/rand"
	"math/big"
	"strconv"
	"strings"
	"unicode"
)

const Base10 = 10
const BitSize64 = 64

// Camel 下划线写法转为驼峰写法
func Camel(name string) string {
	name = strings.ReplaceAll(name, "_", " ")
	name = strings.Title(name)

	return strings.ReplaceAll(name, " ", "")
}

// Underscore 驼峰式写法转为下划线写法
// 将大写转成 _a 的形式
func Underscore(str string) string {
	var buf strings.Builder

	for index, value := range str {
		// value >= 'A' && value <= 'Z'
		if unicode.IsUpper(value) {
			if index != 0 {
				buf.WriteByte('_')
			}

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
	rnd, _ := rand.Int(rand.Reader, big.NewInt(64)) //nolint:gomnd //nt在[0, max) 中返回一个统一的随机值。如果 max <= 0，它会发生混乱。

	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rnd.Int64()%int64(len(letterBytes))]
	}

	return string(b)
}

func StrToInt64(str string) (int64, error) {
	return strconv.ParseInt(str, Base10, BitSize64)
}
func Uint64ToStr(num uint64) string {
	return strconv.FormatUint(num, Base10)
}
