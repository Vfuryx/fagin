package main

import (
	"fagin/pkg/consoles"
	"flag"
)

// @title server
// @version 2.0
// @description fagin服务.
// @termsOfService https://github.com/Vfuryx/fagin
func main() {
	flag.Parse()
	consoles.Execute()
}
