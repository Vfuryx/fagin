package mq

import (
	"fagin/app"
	"runtime/debug"
)

// Init 初始化
func Init() {
	rabbitmq, err := NewAdminLogMQ()
	if err != nil {
		go app.Log().Error(err, string(debug.Stack()))
		panic(err)
	}

	go func() {
		err = rabbitmq.Consume()
		if err != nil {
			go app.Log().Error(err, string(debug.Stack()))
		}

		_ = rabbitmq.Destroy()
	}()
}
