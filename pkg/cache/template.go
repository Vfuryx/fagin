package cache

import (
	"fagin/app/utils"
	"fagin/config"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateCacheTemplate(path, name string) {
	filePath := config.App.AppPath + "/cache/" + path + ".go"
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

	const Temp = `package %[1]s

import (
	"fagin/pkg/cache"
	"time"
)

// 网站配置缓存管理
type %[2]s struct {
	cache.SCache
}

func New%[3]s(f cache.GetterFunc) *%[2]s {
	var c = new(%[2]s)
	c.Prefix = "prefix"
	c.LifeTime = 60 * time.Second
	c.Content = f
	c.SetFunc(c)
	return c
}

// Key 获取键名称
func (c *%[2]s) Key(value string) string {
	return c.Prefix + value
}

// Lift 默认存在时间
func (c *%[2]s) Lift() time.Duration {
	return c.LifeTime
}
`

	content := fmt.Sprintf(Temp, packageName, structName, name)
	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("cache create run successfully")
}
