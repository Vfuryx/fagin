package mq

import (
	"encoding/json"
	"errors"
	"fagin/app"
	"fagin/app/caches"
	"fagin/app/models/admin_operation_log"
	"fagin/app/service"
	"fagin/config"
	"fagin/pkg/logger"
	"fagin/pkg/rabbitmq"
	"fagin/utils"
	"fmt"
	"github.com/streadway/amqp"
	"runtime/debug"
	"sync"
	"time"
)

type adminLogMQ struct {
	*rabbitmq.RabbitMQ
	sync.Mutex
}

// AdminLog 日志
type AdminLog struct {
	LatencyTime time.Duration
	Body        []byte
	AdminID     uint
	LogType     uint8
	Method,
	ClientIP,
	UserName,
	UserAgent,
	URI,
	Path,
	Content string
}

// AdminLogMQ 消息队列日志
var AdminLogMQ *adminLogMQ

var _ rabbitmq.AMQP = &adminLogMQ{}

func init() {
	var err error
	if config.AMQP().Open {
		AdminLogMQ, err = NewAdminLogMQ()
		if err != nil {
			panic(err)
		}
	}
}

// NewAdminLogMQ 消息队列日志
func NewAdminLogMQ() (mq *adminLogMQ, err error) {
	mq = new(adminLogMQ)
	mq.RabbitMQ, err = rabbitmq.NewRabbitMQ("admin_log_mq", "", "")
	return
}

func (a *adminLogMQ) Publish(msg amqp.Publishing) error {
	if a.Channel == nil {
		return errors.New("channel 不存在")
	}
	err := a.Channel.Publish(
		a.Exchange,  // exchange
		a.QueueName, // routing key
		false,       // mandatory
		false,       // immediate
		msg,
	)
	if err != nil {
		fmt.Println("channel 不存在", err)
		return err
	}
	return nil
}

func (a *adminLogMQ) Consume() error {
	if a.Channel == nil {
		return errors.New("空")
	}
	_, err := a.Channel.QueueDeclare(
		a.QueueName, // name
		true,        // durable
		false,       // delete when usused
		false,       // exclusive
		false,       // no-wait
		nil,         // arguments
	)
	if err != nil {
		return err
	}

	err = a.Channel.Qos(1, 0, false)
	if err != nil {
		return err
	}

	msg, err := a.Channel.Consume(
		a.QueueName, // queue
		"",          // consumer
		false,       // auto-ack
		false,       // exclusive
		false,       // no-local
		false,       // no-wait
		nil,         // args
	)
	if err != nil {
		return err
	}

	forever := make(chan bool)
	go func() {
		for d := range msg {
			var l AdminLog
			err := json.Unmarshal(d.Body, &l)
			if err != nil {
				go app.Log(logger.AdminMode).Error(err, string(debug.Stack()))
			} else {
				LogStore(
					l.Body, l.LatencyTime, l.AdminID, l.Method, l.ClientIP,
					l.UserName, l.UserAgent, l.URI, l.Path,
				)
				err = d.Ack(false)
				if err != nil {
					go app.Log(logger.AdminMode).Error(err, string(debug.Stack()))
				}
			}
			time.Sleep(1 * time.Second)
		}
	}()
	<-forever
	return nil
}

// LogStore 日志
func LogStore(
	body []byte, latencyTime time.Duration, adminID uint, reqMethod, clientIP,
	userName, userAgent, uri, path string,
) {
	log := admin_operation_log.New()
	log.User = userName
	log.Method = reqMethod
	log.IP = clientIP
	log.Location = utils.GetLocation(clientIP)
	log.LatencyTime = latencyTime.String()
	log.UserAgent = userAgent
	log.Status = 1
	log.AdminID = adminID

	// 请求路由
	log.Path = uri
	log.Operation = map[string]string{
		"POST":   admin_operation_log.OperationADD,
		"PUT":    admin_operation_log.OperationUpdate,
		"DELETE": admin_operation_log.OperationDelete,
	}[reqMethod]
	menuCache := caches.NewAdminOperationLog(func() ([]byte, error) {
		// 缓存菜单
		m, err := service.AdminMenuService.FindByPath(reqMethod, path, []string{"title"})
		if err != nil {
			return nil, err
		}
		return []byte(m.Title), nil
	})
	str, err := menuCache.Get(reqMethod, path)
	if err != nil {
		log.Module = ""
	} else {
		log.Module = string(str)
	}
	log.Input = string(body)

	if err = log.Dao().Create(log); err != nil {
		go app.Log(logger.AdminMode).Error(err, string(debug.Stack()))
	}
}
