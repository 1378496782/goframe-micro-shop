package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/streadway/amqp"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	ctx     context.Context
}

// NewRabbitMQ 创建RabbitMQ实例
func NewRabbitMQ(ctx context.Context) (*RabbitMQ, error) {
	host := g.Cfg().MustGet(ctx, "rabbitmq.default.host").String()
	port := g.Cfg().MustGet(ctx, "rabbitmq.default.port").String()
	user := g.Cfg().MustGet(ctx, "rabbitmq.default.user").String()
	password := g.Cfg().MustGet(ctx, "rabbitmq.default.password").String()
	vhost := g.Cfg().MustGet(ctx, "rabbitmq.default.vhost").String()

	url := fmt.Sprintf("amqp://%s:%s@%s:%s/%s", user, password, host, port, vhost)
	g.Log().Info(ctx, "NewRabbitMQ的配置：", url)
	conn, err := amqp.Dial(url)
	if err != nil {
		return nil, err
	}

	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, err
	}

	return &RabbitMQ{
		conn:    conn,
		channel: channel,
		ctx:     ctx,
	}, nil
}

// Publish 发布消息
func (r *RabbitMQ) Publish(exchange, routingKey string, message interface{}) error {
	body, err := json.Marshal(message)
	if err != nil {
		return err
	}

	return r.channel.Publish(
		exchange,
		routingKey,
		false,
		false,
		amqp.Publishing{
			ContentType: "application/json",
			Body:        body,
		},
	)
}

// Consume 消费消息
func (r *RabbitMQ) Consume(queue, consumer string, autoAck bool) (<-chan amqp.Delivery, error) {
	return r.channel.Consume(
		queue,
		consumer,
		autoAck,
		false,
		false,
		false,
		nil,
	)
}

// Close 关闭连接
func (r *RabbitMQ) Close() {
	if r.channel != nil {
		r.channel.Close()
	}
	if r.conn != nil {
		r.conn.Close()
	}
}

// DeclareExchange 声明交换机
func (r *RabbitMQ) DeclareExchange(name, kind string) error {
	return r.channel.ExchangeDeclare(
		name,
		kind,
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		nil,   // args
	)
}

// DeclareQueue 声明队列
func (r *RabbitMQ) DeclareQueue(name string) (amqp.Queue, error) {
	return r.channel.QueueDeclare(
		name,
		true,  // durable
		false, // autoDelete
		false, // exclusive
		false, // noWait
		nil,   // args
	)
}

// QueueBind 绑定队列到交换机
func (r *RabbitMQ) QueueBind(queue, key, exchange string) error {
	return r.channel.QueueBind(
		queue,
		key,
		exchange,
		false, // noWait
		nil,   // args
	)
}
