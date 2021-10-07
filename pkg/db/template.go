package db

import (
	"fagin/config"
	"fagin/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

// CreateModelTemplate 创建模版
func CreateModelTemplate(path, name string) {
	modelPath := config.App().AppPath + "/models/" + path + "/" + name + ".go"
	daoPath := config.App().AppPath + "/models/" + path + "/" + name + "_dao.go"
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

// New 实例化
func New() *%[2]s {
	return &%[2]s{}
}

// Dao Dao
type Dao struct {
	db.Dao
}

var _ db.IDao = &Dao{}

// Dao 获取Dao
func (m *%[2]s) Dao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(m)
	return dao
}

// NewDao 实例化
func NewDao() *Dao {
	dao := &Dao{}
	dao.Dao.SetModel(New())
	return dao
}

// All 所有数据
func (d *Dao) All(columns []string) (*[]%[2]s, error) {
	var model []%[2]s
	err := db.ORM().Select(columns).Find(&model).Error
	return &model, err
}

// Query 通用查询
func (d *Dao) Query(params map[string]interface{}, columns []string, with map[string]interface{}) db.IDao {
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

// Deletes 批量删除
func (d *Dao) Deletes(ids []uint) error {
	return db.ORM().Where("id in (?)", ids).Delete(d.GetModel()).Error
}
`
	content := fmt.Sprintf(ModelTemp, packageName, utils.Camel(name))
	if _, err = modelFile.WriteString(content); err != nil {
		panic(err)
	}
	content = fmt.Sprintf(DaoTemp, packageName, utils.Camel(name), config.App().Name)
	if _, err = daoFile.WriteString(content); err != nil {
		panic(err)
	}

	log.Printf("model create run successfully")
}
