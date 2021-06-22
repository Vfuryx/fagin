package mq

import (
	"fagin/app"
	"runtime/debug"
)

func InitMQ()  {
	rabbitmq, err := NewAdminLogMQ()
	if err != nil {
		go app.Log().Println(err, string(debug.Stack()))
		panic(err)
	}
	go func() {
		defer rabbitmq.Destroy()
		err = rabbitmq.Consume()
		if err != nil {
			go app.Log().Println(err, string(debug.Stack()))
		}
	}()
}