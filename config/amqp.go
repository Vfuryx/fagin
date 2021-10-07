package config

import (
	"fmt"
)

type AMQPConfig struct {
	Open     bool   // 是否开启
	User     string // 用户名
	Password string // 密码
	Addr     string // 地址
	Port     string // 端口
	Host     string // 虚拟 host
}

var amqpConfig = new(AMQPConfig)

func AMQP() AMQPConfig {
	return *amqpConfig
}

func (amqp *AMQPConfig) init() {
	amqp.Open = false
	amqp.User = "admin"
	amqp.Password = "admin"
	amqp.Addr = "127.0.0.1"
	amqp.Port = "5672"
	amqp.Host = "log"
}

// GetConnectLink 获取 conn 链接
func (amqp AMQPConfig) GetConnectLink() string {
	const link = "amqp://%s:%s@%s:%s/%s"
	return fmt.Sprintf(link, amqp.User, amqp.Password, amqp.Addr, amqp.Port, amqp.Host)
}
