package config_test

import (
	"bufio"
	"fagin/config"
	"os"
	"strings"
	"testing"
)

// 强行根据 go.mod 来获取 name ，同时可以作为引入的路径
// 找到能直接获取 mod 名称的方法要替换这个方法
func getAppName() string {
	var err error

	file, err := os.Open(config.App.RootPath + "/go.mod")
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Scan()
	s := strings.Split(scanner.Text(), " ")

	if err = scanner.Err(); err != nil {
		panic(err)
	}

	if err = file.Close(); err != nil {
		panic(err)
	}

	return s[1]
}

func TestName(t *testing.T) {
	if config.App.GetName() != getAppName() {
		t.Fatal("获取项目名出错")
	}
	t.Log("项目名称：", config.App.GetName())
}
