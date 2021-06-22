package response

import (
	"fagin/app/utils"
	"fagin/config"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateResponseTemplate(path, name string) {
	filePath := config.App.AppPath + "/responses/" + path + ".go"
	sl := strings.Split(filePath, "/")
	dirPath := strings.Join(sl[:len(sl)-1], "/")
	packageName := sl[len(sl)-2]
	name = utils.Camel(name)
	structName := strings.ToLower(string(name[0])) + name[1:]

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

	const temp = `package %s_response

import (
	"%[4]s/app/models/M"
	"%[4]s/pkg/response"
)

type %[2]s struct {
	ms []M.m
	response.Collect
}
var _ response.IResponse = &%[2]s{}

func %[3]s(models ...M.m) *%[2]s {
	res := %[2]s{ms:models}
	res.Collect.IResponse = &res
	return &res
}

func (res *%[2]s) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.ms {
		m := map[string]interface{}{
			"id": model.ID,
		}
		sm = append(sm, m)
	}
	return sm
}
`
	content := fmt.Sprintf(temp, packageName, structName, name, config.App.Name)

	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("response create run successfully")
}