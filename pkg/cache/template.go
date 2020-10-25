package cache

import (
	"fagin/app"
	"fagin/config"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateCacheTemplate( name string) {
	Path := config.App.AppPath + "/cache/" + name + ".go"
	sl := strings.Split(Path, "/")
	packageName := sl[len(sl)-2]
	dirPath := strings.Join(sl[:len(sl)-1], "/")

	//os.Stat获取文件信息
	if _, err := os.Stat(Path); err == nil {
		panic("文件已存在")
	}

	// 创建文件夹 可以多层
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		panic(err)
	}

	file, err := os.Create(Path)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	const Temp = `package %[1]s

import (
	"encoding/json"
	"fagin/app/models/model"
	"fagin/pkg/cache"
	"time"
)

// 网站配置缓存管理
type %[2]s struct {
	cache.SCache
	Content func() (model, error)
}

// 实例
var %[3]s = New%[3]s()

func New%[3]s() *%[2]s {
	var c = new(%[2]s)
	c.Prefix = "prefix"
	c.LifeTime = 60 * time.Second
	c.SetFunc(c)
	return c
}

// 获取键名称
func (c *%[2]s) Key(value string) string {
	return c.Prefix + value
}

// 默认存在时间
func (c *%[2]s) Lift() time.Duration {
	return c.LifeTime
}

//获取数据
func (c *%[2]s) GetContent(id string) (string, error) {
	str, err := c.Content()
	if err != nil {
		return "", err
	}
	data, err := json.Marshal(str)
	if err != nil {
		return "", err
	}
	return string(data), nil
}
`

	n := app.Camel(name)
	content := fmt.Sprintf(Temp, packageName, app.LFirst(n), n)
	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	log.Printf("cache create run successfully")
}
