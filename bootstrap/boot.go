package bootstrap

import (
	"fagin/app/mq"
	"fagin/config"
	"fagin/pkg/cache"
	"fagin/pkg/casbins"
	"fagin/pkg/db"
	"fagin/pkg/logger"
	"fagin/pkg/request"
	"fagin/pkg/router"
	"fagin/pkg/server"
	"fmt"
)

// Run 运行
func Run() {
	// 初始化配置
	config.Init()

	// 初始化日志
	logger.Init()

	// 初始化 db
	db.Init()
	// 关闭orm
	defer db.Close()

	// 初始化 cache 并且延迟关闭redis
	cache.Init()
	defer cache.Close()

	// 初始化 casbins
	casbins.Init()

	// 初始化翻译器
	if err := request.InitTrans(config.App().Locale()); err != nil {
		fmt.Printf("init trans failed, err:%v\n", err)
		return
	}

	// 初始化MQ
	if config.AMQP().Open() {
		mq.Init()
	}

	// 初始化 web 服务
	server.ListenAndServe(router.Init())
}
