package server

import (
	"context"
	"fagin/app"
	"fagin/app/mq"
	"fagin/config"
	"fagin/pkg/cache"
	"fagin/pkg/casbins"
	"fagin/pkg/db"
	"fagin/pkg/request"
	"fagin/pkg/router"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func Run() {
	db.Init()
	// 关闭orm
	defer func() {
		err := db.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()

	casbins.Init()

	// 初始化redis并且延迟关闭redis
	defer func() {
		if cache.Redis != nil {
			cache.Redis.Close()
		}
	}()

	// 初始化翻译器
	if err := request.InitTrans(config.App.Locale); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	// MQ
	if config.AMQP.Open {
		mq.InitMQ()
	}

	// 设置服务
	srv := &http.Server{
		Addr:    "127.0.0.1:" + config.App.Port,
		Handler: router.New(),
		//ReadTimeout:    setting.ReadTimeout,
		//WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 32 << 20,
	}

	// 不是正式环境打印监听端口
	if ok := app.IsProd(); !ok {
		fmt.Printf("[GIN-debug] Listening and serving HTTP on %s \n", config.App.Port)
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

	ctx, cancel := context.WithTimeout(context.Background(), 30*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server Shutdown:", err)
	}
	log.Println("Server exiting")
}
