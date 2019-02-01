package main

import (
	"fagin/app/models"
	_ "fagin/config" // 加载配置
	"fagin/routes"
	"github.com/gin-gonic/gin"
	"net/http"
)

var Gin = gin.Default()

func main() {

	// 获取路由配置
	router := routes.InitRoute()

	// 关闭orm
	defer models.ORM.Close()

	// 设置服务
	s := &http.Server{
		Addr:    ":8080",
		Handler: router,
		//ReadTimeout:    setting.ReadTimeout,
		//WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 8 << 20,
	}

	// 开启监听
	s.ListenAndServe()
}
