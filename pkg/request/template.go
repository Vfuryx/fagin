package request

import (
	"fagin/config"
	"fagin/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateRequestTemplate(path, name string) {
	filePath := config.App().AppPath() + "/requests/" + path + ".go"
	sl := strings.Split(filePath, "/")
	packageName := sl[len(sl)-2]
	name = utils.Camel(name)
	structName := strings.ToLower(string(name[0])) + name[1:]
	dirPath := strings.Join(sl[:len(sl)-1], "/")

	// os.Stat获取文件信息
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

	const temp = `package request

import (
	"%[3]s/pkg/request"

	"github.com/gin-gonic/gin"
)

type %[2]s struct {

}

var _ request.Request = %[2]s{}

func (%[2]s) Message() map[string]string {
	return map[string]string{
	}
}

func (%[2]s) Attributes() map[string]string {
	return map[string]string{
	}
}

func (r %[2]s) Validate(ctx *gin.Context) (any, map[string]string) {
	return request.Validate[%[2]s](r, ctx)
}
`

	var content = fmt.Sprintf(temp, packageName, structName, config.App().Name(), name)

	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("request create run successfully")
}
