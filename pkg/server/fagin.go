package server

import (
	"context"
	"fagin/config"
	"fagin/pkg/db"
	"fagin/routes"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	// 关闭orm
	defer db.Close()

	// 获取路由配置
	router := routes.InitRoute()

	// 设置服务
	srv := &http.Server{
		Addr:    ":" + config.App.Port,
		Handler: router,
		//ReadTimeout:    setting.ReadTimeout,
		//WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 8 << 20,
	}

	// 优雅重启
	go func() {
		// service connections
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	signal.Notify(quit, os.Interrupt)
	<-quit
	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 30 * time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}