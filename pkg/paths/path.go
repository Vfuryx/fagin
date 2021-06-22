package paths

import (
	"fagin/app/utils"
	"os"
	"strings"
)

var RootPath string

// 获取项目根目录
func init() {
	var err error
	RootPath, err = os.Getwd()
	if err != nil {
		panic(err)
	}
}

func GetPathAndUnderscore(name string) (string, string) {
	str := strings.Trim(name, " /\\")
	sl := strings.Split(str, "/")
	for index, value := range sl {
		sl[index] = utils.Underscore(value)
	}
	path := strings.Join(sl, "/")
	return path, sl[len(sl)-1]
}