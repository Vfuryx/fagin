package config

import (
	"fmt"
)

type amqp struct {
	Open     bool   // 是否开启
	User     string // 用户名
	Password string // 密码
	Addr     string // 地址
	Port     string // 端口
	Host     string // 虚拟 host
}

var AMQP amqp

func init() {
	AMQP.Open = false
	AMQP.User = "admin"
	AMQP.Password = "admin"
	AMQP.Addr = "127.0.0.1"
	AMQP.Port = "5672"
	AMQP.Host = "log"
}

// GetConnectLink 获取 conn 链接
func (amqp *amqp) GetConnectLink() string {
	const link = "amqp://%s:%s@%s:%s/%s"
	return fmt.Sprintf(link, amqp.User, amqp.Password, amqp.Addr, amqp.Port, amqp.Host)
}
