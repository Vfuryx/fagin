package paths

import (
	"fagin/utils"
	"os"
	"strings"
)

func GetPathAndUnderscore(name string) (string, string) {
	str := strings.Trim(name, " /\\")
	sl := strings.Split(str, "/")
	for index, value := range sl {
		sl[index] = utils.Underscore(value)
	}
	path := strings.Join(sl, "/")
	return path, sl[len(sl)-1]
}

func FileExists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}
