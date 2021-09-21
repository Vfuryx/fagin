package service

import (
	"fagin/config"
	"fagin/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateServiceTemplate(path, name string) {
	filePath := config.App.AppPath + "/service/" + path + ".go"
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

	const temp = `package %[1]s

import (
	"fagin/pkg/db"
	"github.com/gin-gonic/gin"
	"fagin/app/models/%[2]s"
)

type %[2]sService struct{}

var %[3]s %[2]sService

func (*%[2]sService) Index(params gin.H, columns []string, with gin.H, p *db.Paginator) (ms []%[2]s.%[3]s, err error) {
	err = %[2]s.NewDao().Query(params, columns, with).Paginate(&ms, p)
	return
}

func (*%[2]sService) Show(id uint, columns []string) (*%[2]s.%[3]s, error) {
	b := %[2]s.New()
	err := b.NewDao().FindById(id, columns)
	return b, err
}

func (*%[2]sService) Create(m *%[2]s.%[3]s) error {
	return %[2]s.Dao().Create(m)
}

func (*%[2]sService) Update(id uint, data gin.H) error {
	return %[2]s.Dao().Update(id, data)
}

func (*%[2]sService) Delete(id uint) error {
	return %[2]s.Dao().Destroy(id)
}

func (*%[2]sService) Deletes(ids []uint) error {
	return %[2]s.Dao().Deletes(ids)
}
`
	content := fmt.Sprintf(temp, packageName, structName, name, config.App.Name)

	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("service create run successfully")
}
