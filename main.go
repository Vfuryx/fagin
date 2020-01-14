package main

import (
	"fagin/pkg/server"
)

// @title server API
// @version 1.0
// @description fagin服务.
// @termsOfService https://github.com/Vfuryx/fagin
// @securityDefinitions.basic BasicAuth
// @in header
// @name Authorization
func main() {
	server.Run()
}
