package utils

import (
	"os"
	"path"
	"strings"
)

// MkdirAll 创建路径
func MkdirAll(absolutePath string) {
	// 是否创建目录
	dir := path.Dir(absolutePath)
	_, err := os.Stat(dir)

	if err != nil && os.IsNotExist(err) {
		_ = os.MkdirAll(dir, os.ModePerm)
	}
}

func PathUnderscore(p string) string {
	str := strings.Trim(p, " /\\")
	sl := strings.Split(str, "/")

	for index, value := range sl {
		sl[index] = Underscore(value)
	}

	return strings.Join(sl, "/")
}
