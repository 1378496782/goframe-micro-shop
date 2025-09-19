package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/cenkalti/backoff/v4"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/streadway/amqp"
	"time"
)

type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	ctx     context.Context
}

// NewRabbitMQ 创建RabbitMQ实例（带指数退避重试）
func NewRabbitMQ(ctx context.Context) (*RabbitMQ, error) {
	var rb *RabbitMQ
	var err error

	// 创建指数退避策略
	expBackoff := backoff.NewExponentialBackOff()
	// 初始重试间隔，第一次重试等待2秒
	expBackoff.InitialInterval = 2 * time.Second
	// 最大重试间隔，重试间隔不会超过30秒
	expBackoff.MaxInterval = 30 * time.Second
	// 最大总重试时间，5分钟后停止重试
	expBackoff.MaxElapsedTime = 5 * time.Minute
	// 随机化因子，添加随机性避免多个客户端同时重试（雪崩效应）
	expBackoff.RandomizationFactor = 0.5

	// 重试操作
	operation := func() error {
		rb, err = createConnection(ctx)
		if err != nil {
			g.Log().Warningf(ctx, "RabbitMQ连接失败，正在重试: %v", err)
			return err
		}
		return nil
	}

	// 执行重试
	g.Log().Info(ctx, "正在尝试连接RabbitMQ（带重试机制）...")
	err = backoff.Retry(operation, expBackoff)
	if err != nil {
		return nil, fmt.Errorf("重试多次后仍无法连接到RabbitMQ: %v", err)
	}

	g.Log().Info(ctx, "RabbitMQ连接成功")
	return rb, nil
}

// NewRabbitMQ 创建RabbitMQ实例
func createConnection(ctx context.Context) (*RabbitMQ, error) {
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

// PublishWithDelay 发布延迟消息
func (r *RabbitMQ) PublishWithDelay(exchange, routingKey string, message interface{}, delayMs int) error {
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
			Headers: amqp.Table{
				"x-delay": delayMs, // 延迟时间，单位毫秒
			},
			DeliveryMode: amqp.Persistent, // 持久化消息
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
	args := amqp.Table{}
	
	// 如果是延迟交换机，需要设置特殊参数
	if kind == "x-delayed-message" {
		args["x-delayed-type"] = "direct" // 指定延迟交换机的底层类型
	}
	
	return r.channel.ExchangeDeclare(
		name,
		kind,
		true,  // durable
		false, // autoDelete
		false, // internal
		false, // noWait
		args,  // arguments
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

// SetQoS 设置服务质量
func (r *RabbitMQ) SetQoS(prefetchCount, prefetchSize int, global bool) error {
	return r.channel.Qos(prefetchCount, prefetchSize, global)
}
