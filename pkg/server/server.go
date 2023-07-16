//go:build !windows
// +build !windows

package server

import (
	"fadmin/app"
	"fadmin/config"
	"fmt"
	"github.com/gofiber/fiber/v2"
	"log"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"

	"github.com/cloudflare/tableflip"
)

func ListenAndServe(srv *fiber.App) {
	upg, err := tableflip.New(tableflip.Options{
		PIDFile: config.App().RootPath() + "/app.pid",
	})
	if err != nil {
		panic(err)
	}

	defer upg.Stop()

	// Do an upgrade on SIGHUP
	go waitUpgrade(upg)

	// 开始监听服务
	go serve(srv)

	// 记录 pid
	err = os.WriteFile(config.App().RootPath()+"/app.pid", []byte(strconv.Itoa(syscall.Getpid())), os.ModePerm)
	if err != nil {
		panic(err)
	}

	if err = upg.Ready(); err != nil {
		panic(err)
	}

	// 不是正式环境打印监听端口
	if ok := app.IsProd(); !ok {
		fmt.Printf("[GIN-debug] Listening and serving HTTP on %s \n", config.App().Port())
	}

	log.Printf("Actual pid is %d", syscall.Getpid())
	log.Printf("ready")

	<-upg.Exit()

	const num = 30
	// Make sure to set a deadline on exiting the process
	// after upg.Exit() is closed. No new upgrades can be
	// performed if the parent doesn't exit.
	time.AfterFunc(num*time.Second, func() {
		log.Println("Graceful shutdown timed out")
		os.Exit(1)
	})

	// Wait for connections to drain.
	_ = srv.Shutdown()
}

// waitUpgrade Do an upgrade on SIGHUP
// 等待升级
func waitUpgrade(upg *tableflip.Upgrader) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGHUP)

	for range sig {
		if err := upg.Upgrade(); err != nil {
			log.Println("Upgrade failed:", err)
		}
	}
}

// 开始监听服务 (阻塞进程)
func serve(srv *fiber.App) {
	// service connections
	if err := srv.Listen("127.0.0.1:" + config.App().Port()); err != nil && err != http.ErrServerClosed {
		log.Fatalf("listen: %s\n", err)
	}
}
