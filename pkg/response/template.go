package response

import (
	"fagin/config"
	"fagin/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateResponseTemplate(path, name string) {
	filePath := config.App().AppPath() + "/responses/" + path + ".go"
	sl := strings.Split(filePath, "/")
	dirPath := strings.Join(sl[:len(sl)-1], "/")
	packageName := sl[len(sl)-2]
	name = utils.Camel(name)
	structName := strings.ToLower(string(name[0])) + name[1:]

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

	const temp = `package %s_response

import (
	"%[4]s/app/models/M"
	"%[4]s/pkg/response"
)

type %[2]s []M.m

func New%[3]s(models ...M.m) *response.Collect[%[2]s] {
	return new(response.Collect[%[2]s]).SetCollect(models)
}

func (res %[2]s) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, response.DefCap)

	for i := range res {
		m := map[string]interface{}{
			"id": res[i].ID,
		}
		sm = append(sm, m)
	}

	return sm
}
`

	content := fmt.Sprintf(temp, packageName, structName, name, config.App().Name())

	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("response create run successfully")
}
