package utility

import (
	"context"
	"fmt"

	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
)

var (
	flashSaleRabbitMQ *RabbitMQ
)

type RabbitMQ struct {
	Conn    *amqp.Connection
	Channel *amqp.Channel
	Ctx     context.Context
}

// InitFlashSaleRabbitMQ 初始化秒杀服务RabbitMQ
func InitFlashSaleRabbitMQ(ctx context.Context) error {
	// 创建RabbitMQ连接配置
	config := g.Cfg().MustGet(ctx, "rabbitmq").MapStrStr()

	// 建立连接
	conn, err := amqp.Dial(fmt.Sprintf("amqp://%s:%s@%s:%s/",
		config["user"], config["password"], config["host"], config["port"]))
	if err != nil {
		return err
	}

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return err
	}

	flashSaleRabbitMQ = &RabbitMQ{
		Conn:    conn,
		Channel: channel,
		Ctx:     ctx,
	}

	// 声明交换机和队列
	if err := declareFlashSaleExchangesAndQueues(ctx); err != nil {
		channel.Close()
		conn.Close()
		return err
	}

	g.Log().Info(ctx, "秒杀服务RabbitMQ初始化成功")
	return nil
}

// declareFlashSaleExchangesAndQueues 声明秒杀相关交换机和队列
func declareFlashSaleExchangesAndQueues(ctx context.Context) error {
	// 声明交换机
	if err := flashSaleRabbitMQ.Channel.ExchangeDeclare(
		"flash_sale.exchange", // 交换机名称
		"direct",              // 交换机类型
		true,                  // 持久化
		false,                 // 自动删除
		false,                 // 内部使用
		false,                 // 等待确认
		nil,                   // 参数
	); err != nil {
		return err
	}

	// 声明队列
	if _, err := flashSaleRabbitMQ.Channel.QueueDeclare(
		"flash_sale.order.queue", // 队列名称
		true,                     // 持久化
		false,                    // 自动删除
		false,                    // 排他性
		false,                    // 等待确认
		nil,                      // 参数
	); err != nil {
		return err
	}

	// 绑定队列到交换机
	if err := flashSaleRabbitMQ.Channel.QueueBind(
		"flash_sale.order.queue", // 队列名称
		"flash_sale.order",       // 路由键
		"flash_sale.exchange",    // 交换机名称
		false,                    // 等待确认
		nil,                      // 参数
	); err != nil {
		return err
	}

	return nil
}

// GetFlashSaleRabbitMQ 获取秒杀RabbitMQ实例
func GetFlashSaleRabbitMQ() *RabbitMQ {
	return flashSaleRabbitMQ
}

// PublishFlashSaleOrder 发布秒杀订单消息
func (r *RabbitMQ) PublishFlashSaleOrder(message []byte) error {
	if r.Channel == nil {
		return fmt.Errorf("RabbitMQ channel is not initialized")
	}

	return r.Channel.Publish(
		"flash_sale.exchange", // 交换机名称
		"flash_sale.order",    // 路由键
		false,                 // 强制
		false,                 // 立即
		amqp.Publishing{
			ContentType:  "application/json",
			Body:         message,
			DeliveryMode: amqp.Persistent, // 消息持久化
		},
	)
}

// Close 关闭RabbitMQ连接
func (r *RabbitMQ) Close() error {
	if r.Channel != nil {
		r.Channel.Close()
	}
	if r.Conn != nil {
		return r.Conn.Close()
	}
	return nil
}
