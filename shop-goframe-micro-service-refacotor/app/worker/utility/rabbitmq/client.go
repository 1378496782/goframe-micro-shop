package rabbitmq

import (
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
)

// RabbitMQClient RabbitMQ客户端封装
type RabbitMQClient struct {
	conn    *amqp.Connection
	channel *amqp.Channel
}

// NewRabbitMQClient 创建RabbitMQ客户端实例
func NewRabbitMQClient(ctx context.Context) (*RabbitMQClient, error) {
	// 从配置获取RabbitMQ连接地址
	address := g.Cfg().MustGet(ctx, "rabbitmq.default.address").String()
	if address == "" {
		return nil, fmt.Errorf("RabbitMQ连接地址不能为空，请检查配置文件")
	}

	g.Log().Info(ctx, "正在连接RabbitMQ服务器: "+address)

	// 建立连接
	conn, err := amqp.Dial(address)
	if err != nil {
		return nil, fmt.Errorf("连接RabbitMQ失败: %v", err)
	}

	// 创建通道
	channel, err := conn.Channel()
	if err != nil {
		conn.Close()
		return nil, fmt.Errorf("创建RabbitMQ通道失败: %v", err)
	}

	client := &RabbitMQClient{
		conn:    conn,
		channel: channel,
	}

	// 初始化交换机和队列
	err = client.initExchangeAndQueue(ctx)
	if err != nil {
		client.Close()
		return nil, err
	}

	g.Log().Info(ctx, "RabbitMQ客户端初始化成功")
	return client, nil
}

// initExchangeAndQueue 初始化延迟交换机和队列
func (r *RabbitMQClient) initExchangeAndQueue(ctx context.Context) error {
	// 获取交换机名称
	exchangeName := g.Cfg().MustGet(ctx, "rabbitmq.default.exchange.orderDelayExchange").String()
	if exchangeName == "" {
		return fmt.Errorf("交换机名称不能为空，请检查配置文件")
	}

	g.Log().Info(ctx, "正在创建延迟交换机: "+exchangeName)

	// 声明延迟交换机
	err := r.channel.ExchangeDeclare(
		exchangeName,        // 交换机名称
		"x-delayed-message", // 交换机类型
		true,                // 持久化
		false,               // 自动删除
		false,               // 内部使用
		false,               // 不等待服务器响应
		amqp.Table{
			"x-delayed-type": "direct", // 指定延迟交换机的底层类型
		},
	)
	if err != nil {
		return fmt.Errorf("创建延迟交换机失败: %v", err)
	}

	// 获取队列名称
	queueName := g.Cfg().MustGet(ctx, "rabbitmq.default.queue.orderTimeoutQueue").String()
	if queueName == "" {
		return fmt.Errorf("队列名称不能为空，请检查配置文件")
	}

	g.Log().Info(ctx, "正在创建队列: "+queueName)

	// 声明队列
	_, err = r.channel.QueueDeclare(
		queueName, // 队列名称
		true,      // 持久化
		false,     // 自动删除
		false,     // 独占
		false,     // 不等待服务器响应
		nil,       // 参数
	)
	if err != nil {
		return fmt.Errorf("创建队列失败: %v", err)
	}

	g.Log().Info(ctx, "正在绑定队列到交换机")

	// 绑定队列到交换机
	err = r.channel.QueueBind(
		queueName,       // 队列名称
		"order.timeout", // 路由键
		exchangeName,    // 交换机名称
		false,           // 不等待服务器响应
		nil,             // 参数
	)
	if err != nil {
		return fmt.Errorf("绑定队列到交换机失败: %v", err)
	}

	g.Log().Info(ctx, "RabbitMQ交换机和队列初始化完成")
	return nil
}

// Close 关闭RabbitMQ连接
func (r *RabbitMQClient) Close() error {
	var err error
	if r.channel != nil {
		err = r.channel.Close()
		if err != nil {
			g.Log().Error(context.Background(), "关闭RabbitMQ通道失败: "+err.Error())
		}
	}
	if r.conn != nil {
		if closeErr := r.conn.Close(); closeErr != nil && err == nil {
			err = closeErr
			g.Log().Error(context.Background(), "关闭RabbitMQ连接失败: "+closeErr.Error())
		}
	}
	return err
}

// TestConnection 测试连接是否正常
func (r *RabbitMQClient) TestConnection(ctx context.Context) error {
	// 创建一个临时队列来测试连接
	testQueue := "test.connection.queue"

	_, err := r.channel.QueueDeclare(
		testQueue, // 队列名称
		false,     // 不持久化
		true,      // 自动删除
		false,     // 不独占
		false,     // 不等待服务器响应
		nil,       // 参数
	)
	if err != nil {
		return fmt.Errorf("测试连接失败: %v", err)
	}

	// 发布测试消息
	err = r.channel.Publish(
		"",        // 使用默认交换机
		testQueue, // 路由键
		false,     // 不强制
		false,     // 不立即
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte("测试消息 " + time.Now().Format(time.RFC3339)),
		})
	if err != nil {
		return fmt.Errorf("发布测试消息失败: %v", err)
	}

	g.Log().Info(ctx, "RabbitMQ连接测试成功")
	return nil
}
