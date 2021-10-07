package service

import (
	"fagin/config"
	"fagin/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

// CreateServiceTemplate 生成模版
func CreateServiceTemplate(path, name string) {
	filePath := config.App().AppPath + "/service/" + path + ".go"
	sl := strings.Split(filePath, "/")
	dirPath := strings.Join(sl[:len(sl)-1], "/")
	packageName := sl[len(sl)-2]
	varName := utils.Camel(name)
	structName := strings.ToLower(string(varName[0])) + varName[1:]

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

	const temp = `package %[1]s

import (
	"fagin/app/models/%[4]s"
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
)

type %[2]sService struct{}

var %[3]s %[2]sService

func (*%[2]sService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) (ms []%[4]s.%[3]s, err error) {
	err = %[4]s.NewDao().Query(params, columns, with).Paginate(&ms, p)
	return
}

func (*%[2]sService) Show(id uint, columns []string) (*%[4]s.%[3]s, error) {
	b := %[4]s.New()
	err := b.Dao().FindById(id, columns)
	return b, err
}

func (*%[2]sService) Create(m *%[4]s.%[3]s) error {
	return %[4]s.NewDao().Create(m)
}

func (*%[2]sService) Update(id uint, data gin.H) error {
	return %[4]s.NewDao().Update(id, data)
}

func (*%[2]sService) Delete(id uint) error {
	return %[4]s.NewDao().Destroy(id)
}

func (*%[2]sService) Deletes(ids []uint) error {
	return %[4]s.NewDao().Deletes(ids)
}
`
	fmt.Println(packageName, structName, varName, name)
	content := fmt.Sprintf(temp, packageName, structName, varName, name, config.App().Name)

	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("service create run successfully")
}
