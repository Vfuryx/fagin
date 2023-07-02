package bootstrap

import (
	"fadmin/config"
	"fadmin/pkg/logger"
	"fadmin/pkg/router"
	"fadmin/pkg/server"
)

// Run 运行
func Run() {
	// 初始化配置
	config.Init()

	// 初始化日志
	logger.Init()

	// 初始化 web 服务
	server.ListenAndServe(router.Init())
}
