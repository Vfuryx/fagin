package main

import (
	_ "github.com/Vfuryx/fagin/database/migrations"
	"github.com/Vfuryx/fagin/database"
	"github.com/Vfuryx/fagin/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Gin = gin.Default()

func main() {
	router := routes.InitRoute()

	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		//ReadTimeout:    setting.ReadTimeout,
		//WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 8 << 20,
	}

	// 关闭orm
	defer database.ORM.Close()

	s.ListenAndServe()
}
