package config

import (
	"fmt"
	"sync/atomic"
)

var amqpConfig atomic.Value

func AMQP() *AMQPConfig {
	if c, ok := amqpConfig.Load().(*AMQPConfig); ok {
		return c
	}

	return &AMQPConfig{}
}

func amqpConfigInit() {
	c := &AMQPConfig{
		open:     false,
		user:     "admin",
		password: "admin",
		addr:     "127.0.0.1",
		port:     "5672",
		host:     "log",
	}

	amqpConfig.Store(c)
}

type AMQPConfig struct {
	open     bool   // 是否开启
	user     string // 用户名
	password string // 密码
	addr     string // 地址
	port     string // 端口
	host     string // 虚拟 host
}

func (amqp *AMQPConfig) Open() bool {
	return amqp.open
}

func (amqp *AMQPConfig) User() string {
	return amqp.user
}

func (amqp *AMQPConfig) Password() string {
	return amqp.password
}

func (amqp *AMQPConfig) Addr() string {
	return amqp.addr
}

func (amqp *AMQPConfig) Port() string {
	return amqp.port
}

func (amqp *AMQPConfig) Host() string {
	return amqp.host
}

// GetConnectLink 获取 conn 链接
func (amqp *AMQPConfig) GetConnectLink() string {
	const link = "amqp://%s:%s@%s:%s/%s"
	return fmt.Sprintf(link, amqp.user, amqp.password, amqp.addr, amqp.password, amqp.host)
}
