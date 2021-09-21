package utils

import (
	"os"
	"path"
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