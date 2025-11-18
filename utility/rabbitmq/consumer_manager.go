package rabbitmq

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"sync"
	"time"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	amqp "github.com/rabbitmq/amqp091-go"

	"shop-goframe-micro-service-refacotor/utility/idempotent"
)

// Consumer 消费者接口
type Consumer interface {
	// GetName 获取消费者名称
	GetName() string

	// GetConfig 获取消费者配置
	GetConfig() ConsumerConfig

	// HandleMessage 处理消息
	HandleMessage(ctx context.Context, msg amqp.Delivery) error
	
	// GetBusinessID 从消息中提取业务ID，用于幂等性检查（可选实现）
	// 如果不实现，将使用消息头中的business_id或空字符串
	GetBusinessID(data []byte, event map[string]interface{}) string
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
	MaxRetries    int    // 最大重试次数，默认3次；设置为0表示不限制
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
	if config.MaxRetries == 0 {
		config.MaxRetries = 3 // 默认最大重试次数为3次
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

	// 记录消息接收日志
	g.Log().Debugf(cm.ctx, "接收到消息: consumer=%s, exchange=%s, routingKey=%s, messageID=%s",
		consumer.GetName(), msg.Exchange, msg.RoutingKey, msg.MessageId)

	// 幂等性检查
	if err := cm.checkIdempotency(consumer, msg); err != nil {
		g.Log().Warningf(cm.ctx, "消息幂等性检查失败或已处理: consumer=%s, messageID=%s, error=%v",
			consumer.GetName(), msg.MessageId, err)
		// 已处理过的消息直接确认
		if !consumer.GetConfig().AutoAck {
			msg.Ack(false)
		}
		return
	}

	start := time.Now()
	err := consumer.HandleMessage(cm.ctx, msg)
	duration := time.Since(start)

	if err != nil {
		g.Log().Errorf(cm.ctx, "消费者 %s 处理消息失败 (耗时 %v): %v",
			consumer.GetName(), duration, err)

		// 检查重试次数
		retryCount := cm.getMessageRetryCount(msg)
		config := consumer.GetConfig()
		
		// 根据错误类型和重试次数决定是否重新入队
		if shouldRequeue(err) && (config.MaxRetries <= 0 || retryCount < config.MaxRetries) {
			// 增加重试计数
			newRetryCount := retryCount + 1
			g.Log().Infof(cm.ctx, "消息将重新入队，当前重试次数: %d, 最大重试次数: %d", 
				newRetryCount, config.MaxRetries)
			
			// 更新消息头中的重试次数
			if msg.Headers == nil {
				msg.Headers = make(amqp.Table)
			}
			msg.Headers["x-retry-count"] = newRetryCount
			
			// 在重新入队前发布一条带有更新后重试次数的消息
			// 注意：这里我们仍然使用Nack让消息重新入队，但会记录重试次数
			// 在实际场景中，如果需要更精确的控制，可以先Nack(false,false)然后重新发布带有新header的消息
			msg.Nack(false, true) // 重新入队
		} else {
			g.Log().Warningf(cm.ctx, "消息达到最大重试次数或为永久性错误，不再重试，重试次数: %d, 最大重试次数: %d", 
			retryCount, config.MaxRetries)
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

// checkIdempotency 检查消息的幂等性
func (cm *ConsumerManager) checkIdempotency(consumer Consumer, msg amqp.Delivery) error {
	// 获取消息相关信息
	messageID := msg.MessageId
	if messageID == "" {
		// 如果没有messageID，生成一个基于消息体的简单标识
		messageID = fmt.Sprintf("%x", msg.Body[:32]) // 使用消息体前32字节作为标识
	}

	// 尝试从消息中提取业务ID
	// 1. 首先尝试使用消费者的GetBusinessID方法
	businessID := ""
	// 解析消息体以获取event数据
	event := make(map[string]interface{})
	if err := UnmarshalEvent(msg.Body, &event); err == nil {
		businessID = consumer.GetBusinessID(msg.Body, event)
	}
	
	// 2. 如果消费者没有提供业务ID，尝试从消息头获取
	if businessID == "" {
		if businessIDHeader, exists := msg.Headers["business_id"]; exists {
			businessID = gconv.String(businessIDHeader)
		}
	}

	// 生成幂等键
	// 格式: rabbitmq:consumer_name:message_id:business_id
	key := idempotent.GenerateMessageKey(
		fmt.Sprintf("rabbitmq:%s", consumer.GetName()),
		messageID,
		businessID,
	)

	// 设置幂等键的过期时间，默认24小时
	expiration := 24 * time.Hour
	if ttlHeader, exists := msg.Headers["idempotent_ttl"]; exists {
		if ttl, ok := ttlHeader.(int64); ok {
			expiration = time.Duration(ttl) * time.Second
		}
	}

	// 尝试获取幂等锁
	locked, err := idempotent.TryLock(cm.ctx, key, expiration)
	if err != nil {
		// 幂等性服务错误，为了不阻塞业务流程，暂时允许继续处理
		g.Log().Errorf(cm.ctx, "幂等性检查服务错误: key=%s, error=%v", key, err)
		return nil // 允许处理，不视为错误
	}

	if !locked {
		// 已存在幂等锁，说明消息已处理过
		return fmt.Errorf("消息已处理过")
	}

	return nil
}

// 定义错误类型，用于控制重试行为
type (
	// TemporaryError 临时性错误，表示可以重试
	TemporaryError struct {
		Err error
	}
	
	// PermanentError 永久性错误，表示不应该重试
	PermanentError struct {
		Err error
	}
)

// Error 实现error接口
func (e TemporaryError) Error() string {
	return fmt.Sprintf("临时性错误: %v", e.Err)
}

// Error 实现error接口
func (e PermanentError) Error() string {
	return fmt.Sprintf("永久性错误: %v", e.Err)
}

// Unwrap 实现errors.Unwrap接口
func (e TemporaryError) Unwrap() error {
	return e.Err
}

// Unwrap 实现errors.Unwrap接口
func (e PermanentError) Unwrap() error {
	return e.Err
}

// IsTemporary 判断是否为临时性错误
func IsTemporary(err error) bool {
	var tempErr TemporaryError
	return errors.Is(err, &tempErr)
}

// IsPermanent 判断是否为永久性错误
func IsPermanent(err error) bool {
	var permErr PermanentError
	return errors.Is(err, &permErr)
}

// getMessageRetryCount 获取消息的当前重试次数
func (cm *ConsumerManager) getMessageRetryCount(msg amqp.Delivery) int {
	// 从消息头获取重试次数
	if retryCount, exists := msg.Headers["x-retry-count"]; exists {
		// 尝试将各种类型转换为整数
		return gconv.Int(retryCount)
	}
	return 0
}

// shouldRequeue 判断是否应该重新入队
func shouldRequeue(err error) bool {
	// 检查是否为临时性错误
	var tempErr TemporaryError
	if errors.As(err, &tempErr) {
		return true
	}
	
	// 对于其他类型的错误，可以根据错误信息判断
	// 例如网络错误、连接错误等临时性问题可以重试
	errMsg := err.Error()
	temporaryErrorPatterns := []string{
		"connection",
		"timeout",
		"network",
		"refused",
		"unavailable",
		"reset by peer",
		"deadline exceeded",
	}
	
	for _, pattern := range temporaryErrorPatterns {
		if strings.Contains(strings.ToLower(errMsg), strings.ToLower(pattern)) {
			return true
		}
	}
	
	// 默认不重新入队
	return false
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

// GetBusinessID 默认实现，返回空字符串
// 具体消费者可以覆盖此方法以提供自定义的业务ID提取逻辑
func (bc *BaseConsumer) GetBusinessID(data []byte, event map[string]interface{}) string {
	return ""
}

// UnmarshalEvent 通用事件解析helper
func UnmarshalEvent(data []byte, event interface{}) error {
	err := json.Unmarshal(data, event)
	if err != nil {
		return fmt.Errorf("解析事件数据失败: %v", err)
	}
	return nil
}
