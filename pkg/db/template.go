package db

import (
	"fagin/app/utils"
	"fagin/config"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateModelTemplate(path, name string) {
	modelPath := config.App.AppPath + "/models/" + path + "/" + name + ".go"
	daoPath := config.App.AppPath + "/models/" + path + "/" + name + "_dao.go"
	sl := strings.Split(modelPath, "/")
	packageName := sl[len(sl)-2]
	dirPath := strings.Join(sl[:len(sl)-1], "/")

	//os.Stat获取文件信息
	if _, err := os.Stat(modelPath); err == nil {
		panic("文件已存在")
	}
	//os.Stat获取文件信息
	if _, err := os.Stat(daoPath); err == nil {
		panic("文件已存在")
	}

	// 创建文件夹 可以多层
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	modelFile, err := os.Create(modelPath)
	if err != nil {
		panic(err)
	}
	defer modelFile.Close()

	daoFile, err := os.Create(daoPath)
	if err != nil {
		panic(err)
	}
	defer daoFile.Close()

	const ModelTemp = `package %s

import (
	"time"
)

type %s struct {
	ID        	uint 		` + "`gorm:\"primary_key\"`" + `
	CreatedAt 	time.Time
	UpdatedAt	time.Time
}
`
	const DaoTemp = `package %[1]s

import (
	"%[3]s/pkg/db"
)

func New() *%[2]s {
	return &%[2]s{}
}

type dao struct {
	db.Dao
}

var _ db.IDao = &dao{}

func (m *%[2]s) Dao() *dao {
	dao := &dao{}
	dao.Dao.M = m
	return dao
}

func Dao() *dao {
	dao := &dao{}
	dao.Dao.M = New()
	return dao
}

func (dao) All(columns []string) (*[]%[2]s, error) {
	var model []%[2]s
	err := db.ORM().Select(columns).Find(&model).Error
	return &model, err
}

func (d *dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.IDao {
	model := db.ORM().Select(columns)

	var (
		v  interface{}
		ok bool
	)
	if v, ok = params["id"]; ok {
		model = model.Where("id = ?", v)
	}

	if v, ok = params["orderBy"]; ok {
		model = model.Order(v)
	}

	d.DB = d.With(model, with)
	return d
}

func (d *dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.M).Error
}
`
	content := fmt.Sprintf(ModelTemp, packageName, utils.Camel(name))
	if _, err = modelFile.WriteString(content); err != nil {
		panic(err)
	}
	content = fmt.Sprintf(DaoTemp, packageName, utils.Camel(name), config.App.Name)
	if _, err = daoFile.WriteString(content); err != nil {
		panic(err)
	}

	log.Printf("model create run successfully")
}

