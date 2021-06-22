package main

import (
	"fagin/pkg/consoles"
	"fagin/pkg/server"
	"flag"
)

// IsConsole 是否命令模式
var IsConsole bool

func init() {
	flag.BoolVar(&IsConsole, "c", false, "cmd 代表 true")
}

// @title server
// @version 2.0
// @description fagin服务.
// @termsOfService https://github.com/Vfuryx/fagin
func main() {
	flag.Parse()
	if IsConsole { // 命令模式
		consoles.Execute()
	} else {
		server.Run()
	}
}
