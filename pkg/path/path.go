package path

import (
	"errors"
	"path"
	"path/filepath"
	"runtime"
)

var RootPath string

func init() {
	var err error
	RootPath, err = getRootPath()
	if err != nil {
		panic(err)
	}
}

// 获取项目根目录
func getRootPath() (string, error) {
	// 获取当前文件路径
	// 0 表示调用runtime.Caller()所在的位置，1表示runtime.Caller()所在函数的调用位置
	_, filePath , _, ok := runtime.Caller(0)
	if !ok {
		return "", errors.New("get root path err")
	}
	return filepath.Abs(path.Dir(filePath)+"/../..")
}