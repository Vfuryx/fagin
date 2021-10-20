package cache

import (
	"fagin/config"
	"fagin/utils"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateCacheTemplate(path, name string) {
	filePath := config.App().AppPath + "/caches/" + path + ".go"
	sl := strings.Split(filePath, "/")
	dirPath := strings.Join(sl[:len(sl)-1], "/")
	packageName := sl[len(sl)-2]
	name = utils.Camel(name)
	//structName := strings.ToLower(string(name[0])) + name[1:]

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

// 缓存
func New%[2]s(f cache.GetterFunc) *cache.Cache {
	var c = new(cache.Cache)
	c.SetKeyFormat("prefix") 
	c.SetLifeTime(60 * time.Second)
	c.SetFunc(f)

	return c
}
`

	content := fmt.Sprintf(Temp, packageName, name)
	if _, err = file.WriteString(content); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	log.Printf("cache create run successfully")
}
