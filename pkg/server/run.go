//go:build !windows
// +build !windows

package server

import (
	"context"
	"fagin/app"
	"fagin/config"
	"fagin/pkg/router"
	"fmt"
	"github.com/cloudflare/tableflip"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

func listenAndServe() {
	upg, err := tableflip.New(tableflip.Options{
		PIDFile: config.App().RootPath + "/app.pid",
	})
	if err != nil {
		panic(err)
	}
	defer upg.Stop()

	// Do an upgrade on SIGHUP
	go func() {
		sig := make(chan os.Signal, 1)
		signal.Notify(sig, syscall.SIGHUP)
		for range sig {
			if err := upg.Upgrade(); err != nil {
				log.Println("Upgrade failed:", err)
			}
		}
	}()

	// Listen must be called before Ready
	ln, err := upg.Listen("tcp", ":"+config.App().Port)
	if err != nil {
		log.Fatalln("Can't listen:", err)
	}

	// 设置服务
	server := &http.Server{
		Addr:    ":" + config.App().Port,
		Handler: router.New(),
		//ReadTimeout:    setting.ReadTimeout,
		//WriteTimeout:   setting.WriteTimeout,
		MaxHeaderBytes: 32 << 20,
	}

	go func() {
		if err := server.Serve(ln); err != http.ErrServerClosed {
			log.Println("HTTP server:", err)
		}
	}()

	err = ioutil.WriteFile(config.App().RootPath+"/app.pid", []byte(strconv.Itoa(os.Getpid())), 0755)
	if err != nil {
		panic(err)
	}

	// 不是正式环境打印监听端口
	if ok := app.IsProd(); !ok {
		fmt.Printf("[GIN-debug] Listening and serving HTTP on %s \n", config.App().Port)
	}
	log.Printf("Actual pid is %d", syscall.Getpid())
	log.Printf("ready")

	if err := upg.Ready(); err != nil {
		panic(err)
	}
	<-upg.Exit()

	// Make sure to set a deadline on exiting the process
	// after upg.Exit() is closed. No new upgrades can be
	// performed if the parent doesn't exit.
	time.AfterFunc(30*time.Second, func() {
		log.Println("Graceful shutdown timed out")
		os.Exit(1)
	})

	// Wait for connections to drain.
	_ = server.Shutdown(context.Background())
}
