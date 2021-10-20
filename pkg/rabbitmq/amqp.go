package rabbitmq

import (
	"fagin/config"
	"fmt"
	"github.com/streadway/amqp"
)

type AMQP interface {
	Publish(msg amqp.Publishing) error // 生产者
	Consume() error                    // 消费者
	Destroy() error                    // 需要关闭channel，才能在下次新建
}

type RabbitMQ struct {
	conn      *amqp.Connection
	Channel   *amqp.Channel
	QueueName string
	Exchange  string
	Key       string
}

func NewRabbitMQ(queueName, exchange, key string) (*RabbitMQ, error) {
	rabbitMQ := &RabbitMQ{
		QueueName: queueName,
		Exchange:  exchange,
		Key:       key,
	}
	link := config.AMQP().GetConnectLink()
	fmt.Println(link)
	var err error
	rabbitMQ.conn, err = amqp.Dial(link)
	if err != nil {
		return nil, err
	}
	rabbitMQ.Channel, err = rabbitMQ.conn.Channel()
	if err != nil {
		return nil, err
	}
	return rabbitMQ, nil
}

func NewRabbitMQSimple(queueName string) (*RabbitMQ, error) {
	return NewRabbitMQ(queueName, "", "")
}

// Destroy 需要关闭channel，才能在下次新建
func (r *RabbitMQ) Destroy() error {
	err := r.Channel.Close()
	if err != nil {
		return err
	}
	err = r.conn.Close()
	if err != nil {
		return err
	}
	return nil
}

//// 消费者
//func (r *RabbitMQ) Consume() error {
//	if r.channel == nil {
//		return errors.New("空")
//	}
//	_, err := r.channel.QueueDeclare(
//		r.QueueName, // name
//		true,        // durable
//		true,        // delete when usused
//		false,       // exclusive
//		false,       // no-wait
//		nil,         // arguments
//	)
//	if err != nil {
//		return err
//	}
//	msg, err := r.channel.Consume(
//		r.QueueName, // queue
//		"",          // consumer
//		true,        // auto-ack
//		false,       // exclusive
//		false,       // no-local
//		false,       // no-wait
//		nil,         // args
//	)
//	if err != nil {
//		return err
//	}
//	forever := make(chan bool)
//
//	go func() {
//		for d := range msg {
//			app.Log().Printf("Received a message: %s", d.Body, string(debug.Stack()))
//		}
//	}()
//
//	//log.Printf("Waiting for messages. To exit press CTRL+C")
//	<-forever
//	return nil
//}
//
//// 生产者
//func (r *RabbitMQ) Publish(msg amqp.Publishing) error {
//	if r.channel == nil {
//		return errors.New("channel 不存在")
//	}
//	err := r.channel.Publish(
//		r.Exchange,  // exchange
//		r.QueueName, // routing key
//		false,       // mandatory
//		false,       // immediate
//		msg,
//	)
//	if err != nil {
//		return err
//	}
//	return nil
//}
