package rabbitmq

import (
	"context"
	"encoding/json"
	"fmt"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/streadway/amqp"
)

// Consumer 消费者接口
type Consumer interface {
	// GetName 获取消费者名称
	GetName() string

	// GetConfig 获取消费者配置
	GetConfig() ConsumerConfig

	// HandleMessage 处理消息
	HandleMessage(ctx context.Context, msg amqp.Delivery) error
}

// ConsumerConfig 消费者配置
type ConsumerConfig struct {
	Exchange      string // 交换机名称
	ExchangeType  string // 交换机类型，默认"topic"
	Queue         string // 队列名称
	RoutingKey    string // 路由键
	ConsumerTag   string // 消费者标签
	AutoAck       bool   // 是否自动确认，默认false
	PrefetchCount int    // 预取数量，默认1
	Durable       bool   // 是否持久化，默认true
}

// ConsumerManager 通用消费者管理器
type ConsumerManager struct {
	rb        *RabbitMQ
	ctx       context.Context
	consumers []Consumer
	wg        sync.WaitGroup
	done      chan struct{}
	once      sync.Once
}

// NewConsumerManager 创建消费者管理器
func NewConsumerManager(ctx context.Context) (*ConsumerManager, error) {
	rb, err := NewRabbitMQ(ctx)
	if err != nil {
		return nil, fmt.Errorf("创建RabbitMQ连接失败: %v", err)
	}

	return &ConsumerManager{
		rb:        rb,
		ctx:       ctx,
		consumers: make([]Consumer, 0),
		done:      make(chan struct{}),
	}, nil
}

// AddConsumer 添加消费者
func (cm *ConsumerManager) AddConsumer(consumer Consumer) {
	cm.consumers = append(cm.consumers, consumer)
	g.Log().Infof(cm.ctx, "添加消费者: %s", consumer.GetName())
}

// Start 启动所有消费者
func (cm *ConsumerManager) Start() error {
	g.Log().Info(cm.ctx, "启动消费者管理器")

	if len(cm.consumers) == 0 {
		g.Log().Warning(cm.ctx, "没有注册任何消费者")
		return nil
	}

	for _, consumer := range cm.consumers {
		cm.wg.Add(1)
		go cm.startConsumer(consumer)
	}

	g.Log().Infof(cm.ctx, "已启动 %d 个消费者", len(cm.consumers))
	return nil
}

// Stop 停止所有消费者
func (cm *ConsumerManager) Stop() {
	cm.once.Do(func() {
		g.Log().Info(cm.ctx, "停止消费者管理器")

		close(cm.done)
		cm.wg.Wait()

		if cm.rb != nil {
			cm.rb.Close()
		}

		g.Log().Info(cm.ctx, "消费者管理器已停止")
	})
}

// startConsumer 启动单个消费者
func (cm *ConsumerManager) startConsumer(consumer Consumer) {
	defer cm.wg.Done()

	config := consumer.GetConfig()
	name := consumer.GetName()

	// 设置默认值
	if config.ExchangeType == "" {
		config.ExchangeType = "topic"
	}
	if config.ConsumerTag == "" {
		config.ConsumerTag = fmt.Sprintf("%s_%d", name, time.Now().Unix())
	}
	if config.PrefetchCount == 0 {
		config.PrefetchCount = 1
	}
	config.Durable = true // 强制持久化

	// 设置队列
	err := cm.setupQueue(config)
	if err != nil {
		g.Log().Errorf(cm.ctx, "设置消费者 %s 队列失败: %v", name, err)
		return
	}

	// 设置QoS
	err = cm.rb.SetQoS(config.PrefetchCount, 0, false)
	if err != nil {
		g.Log().Errorf(cm.ctx, "设置消费者 %s QoS失败: %v", name, err)
		return
	}

	// 开始消费消息
	msgs, err := cm.rb.Consume(config.Queue, config.ConsumerTag, config.AutoAck)
	if err != nil {
		g.Log().Errorf(cm.ctx, "消费者 %s 启动失败: %v", name, err)
		return
	}

	g.Log().Infof(cm.ctx, "消费者 %s 已启动", name)

	for {
		select {
		case <-cm.done:
			g.Log().Infof(cm.ctx, "消费者 %s 停止", name)
			return
		case msg, ok := <-msgs:
			if !ok {
				g.Log().Infof(cm.ctx, "消费者 %s 消息通道关闭", name)
				return
			}
			go cm.handleMessage(consumer, msg)
		}
	}
}

// setupQueue 设置队列
func (cm *ConsumerManager) setupQueue(config ConsumerConfig) error {
	// 声明交换机
	err := cm.rb.DeclareExchange(config.Exchange, config.ExchangeType)
	if err != nil {
		return fmt.Errorf("声明交换机失败: %v", err)
	}

	// 声明队列
	q, err := cm.rb.DeclareQueue(config.Queue)
	if err != nil {
		return fmt.Errorf("声明队列失败: %v", err)
	}

	// 绑定队列
	err = cm.rb.QueueBind(q.Name, config.RoutingKey, config.Exchange)
	if err != nil {
		return fmt.Errorf("绑定队列失败: %v", err)
	}

	return nil
}

// handleMessage 处理消息
func (cm *ConsumerManager) handleMessage(consumer Consumer, msg amqp.Delivery) {
	defer func() {
		if r := recover(); r != nil {
			g.Log().Errorf(cm.ctx, "消费者 %s 处理消息时发生panic: %v", consumer.GetName(), r)
			msg.Nack(false, false) // 拒绝消息，不重新入队
		}
	}()

	start := time.Now()
	err := consumer.HandleMessage(cm.ctx, msg)
	duration := time.Since(start)

	if err != nil {
		g.Log().Errorf(cm.ctx, "消费者 %s 处理消息失败 (耗时 %v): %v",
			consumer.GetName(), duration, err)

		// 根据错误类型决定是否重新入队
		if shouldRequeue(err) {
			msg.Nack(false, true) // 重新入队
		} else {
			msg.Nack(false, false) // 不重新入队
		}
		return
	}

	g.Log().Debugf(cm.ctx, "消费者 %s 成功处理消息 (耗时 %v)",
		consumer.GetName(), duration)

	if !consumer.GetConfig().AutoAck {
		msg.Ack(false) // 确认消息
	}
}

// shouldRequeue 判断是否应该重新入队
func shouldRequeue(err error) bool {
	// 可以根据错误类型来判断是否重新入队
	// 例如：网络错误、临时性错误等可以重新入队
	// 数据格式错误、业务逻辑错误等不重新入队
	return false // 默认不重新入队
}

// BaseConsumer 基础消费者实现，可以被具体消费者嵌入
type BaseConsumer struct {
	name   string
	config ConsumerConfig
}

// NewBaseConsumer 创建基础消费者
func NewBaseConsumer(name string, config ConsumerConfig) *BaseConsumer {
	return &BaseConsumer{
		name:   name,
		config: config,
	}
}

// GetName 获取消费者名称
func (bc *BaseConsumer) GetName() string {
	return bc.name
}

// GetConfig 获取消费者配置
func (bc *BaseConsumer) GetConfig() ConsumerConfig {
	return bc.config
}

// UnmarshalEvent 通用事件解析helper
func UnmarshalEvent(data []byte, event interface{}) error {
	err := json.Unmarshal(data, event)
	if err != nil {
		return fmt.Errorf("解析事件数据失败: %v", err)
	}
	return nil
}
