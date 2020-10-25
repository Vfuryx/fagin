package response

import (
	"fagin/app"
	"fagin/config"
	"fmt"
	"log"
	"os"
	"strings"
)

type Response interface {
	Serialize() []map[string]interface{}
	Handle() []map[string]interface{}
	Item() map[string]interface{}
	Collection() []map[string]interface{}
}

type Collect struct {
	Response
}

func (c Collect) Handle() []map[string]interface{} {
	return c.Serialize()
}

func (c Collect) Item() map[string]interface{} {
	return c.Handle()[0]
}

func (c Collect) Collection() []map[string]interface{} {
	return c.Handle()
}

func CreateResponseTemplate(path, name string) {
	filePath := config.App.AppPath + "/responses/" + path + ".go"
	sl := strings.Split(filePath, "/")
	dirPath := strings.Join(sl[:len(sl)-1], "/")
	packageName := sl[len(sl)-2]
	name = app.Camel(name)
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

	const temp = `package %s

import (
	"%[4]s/app/models/M"
	"%[4]s/pkg/response"
)

type %[2]s struct {
	Ms []M.m
	response.Collect
}
var _ response.Response = &%[2]s{}

func %[3]s(models ...M.m) *%[2]s {
	res := %[2]s{Ms:models}
	res.Collect.Response = &res
	return &res
}

func (res *%[2]s) Serialize() []map[string]interface{} {
	sm := make([]map[string]interface{}, 0, 20)
	for _, model := range res.Ms {
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