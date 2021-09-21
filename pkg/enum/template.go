package enum

import (
	"fagin/config"
	"fagin/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateEnumTemplate(path, name string) {
	filePath := config.App.AppPath + "/enums/" + path + ".go"
	sl := strings.Split(filePath, "/")
	packageName := sl[len(sl)-2]
	name = utils.Camel(name)
	structName := strings.ToLower(string(name[0])) + name[1:]
	dirPath := strings.Join(sl[:len(sl)-1], "/")

	//os.Stat获取文件信息
	if _, err := os.Stat(filePath); err == nil {
		panic("文件已存在")
	}

	// 创建文件夹 可以多层
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(filePath)
	if err != nil {
		panic(err)
	}

	const temp = `package %s

import (
	"%[3]s/pkg/enum"
)

// %[4]s 枚举类型
type %[4]s int

var _ enum.Enum = new(%[4]s)

const (
	Code  %[4]s = 0 // 枚举
)

func (code %[4]s) String() string {
	switch code {
	case Code:
		return "名称"
	}
	return ""
}

func (code %[4]s) Get() int {
	return int(code)
}

// All%[4]s 所有
func All%[4]s() map[string]int {
	return map[string]int{
		Code.String():  Code.Get(),
	}
}
`
	content := fmt.Sprintf(temp, packageName, structName, config.App.Name, name)

	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("request create run successfully")
}
