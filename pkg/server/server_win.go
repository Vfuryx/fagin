//go:build windows
// +build windows

package server

import (
	"context"
	"fagin/app"
	"fagin/config"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func ListenAndServe(handler http.Handler) {
	// 设置服务
	srv := &http.Server{
		Addr:    "127.0.0.1:" + config.App().Port,
		Handler: handler,
		//ReadTimeout:    setting.ReadTimeout,
		//WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 32 << 20,
	}

	// 不是正式环境打印监听端口
	if ok := app.IsProd(); !ok {
		fmt.Printf("[GIN-debug] Listening and serving HTTP on %s \n", config.App().Port)
	}

	// 开始监听服务
	go serve(srv)

	// Wait for interrupt signal to gracefully shutdown the server with
	// a timeout of 5 seconds.
	quit := make(chan os.Signal)
	// kill (no param) default send syscanll.SIGTERM
	// kill -2 is syscall.SIGINT
	// kill -9 is syscall. SIGKILL but can"t be catch, so don't need add it
	signal.Notify(quit, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	log.Println("Shutdown Server ...")

	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}

	// catching ctx.Done(). timeout of 1 seconds.
	select {
	case <-ctx.Done():
		log.Println("timeout of 1 seconds.")
	}
	log.Println("Server exiting")
}

// 开始监听服务 (阻塞进程)
func serve(server *http.Server) {
	// service connections
	if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}