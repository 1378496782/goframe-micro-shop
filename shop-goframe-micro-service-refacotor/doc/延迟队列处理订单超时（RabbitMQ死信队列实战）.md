# 延迟队列处理订单超时（RabbitMQ死信队列实战）

每篇教程都讲清楚了概念，也讲清楚了在咱们项目中是如何实现和落地的。

## 1. 延迟队列和死信队列的基本概念

### 1.1 什么是延迟队列？

延迟队列是一种特殊的消息队列，它允许消息在发送后的一定时间延迟后才被消费。在电商系统中，延迟队列常用于处理订单超时自动取消、优惠券到期提醒、定时任务调度等场景。

### 1.2 什么是死信队列？

死信队列（Dead Letter Queue，DLQ）是用于存储无法被正常消费的消息的队列。当消息满足以下任一条件时，会被发送到死信队列：

1. 消息被拒绝（basic.reject 或 basic.nack）并且 requeue=false
2. 消息的 TTL（Time-To-Live）过期
3. 队列达到最大长度，无法再添加新消息

### 1.3 延迟队列的实现方式

在RabbitMQ中，实现延迟队列主要有两种方式：

1. **TTL + 死信队列**：设置消息的TTL，当消息过期后会被转发到死信队列
2. **插件方式**：使用 RabbitMQ Delayed Message Exchange 插件

本项目采用的是第二种方式，通过安装和配置 RabbitMQ Delayed Message Exchange 插件来实现延迟队列功能。

## 2. 为什么需要使用延迟队列处理订单超时？

在电商系统中，订单创建后通常需要用户在一定时间内完成支付，否则订单应该被自动取消。处理这种场景有几种常见方案：

### 2.1 常见方案对比

| 方案 | 优点 | 缺点 |
|------|------|------|
| 定时任务轮询 | 实现简单 | 1. 时间精度低<br>2. 对数据库压力大<br>3. 资源浪费 |
| Redis过期监听 | 性能好 | 1. 需要额外的Redis集群<br>2. 实现复杂度高<br>3. 存在消息丢失风险 |
| 延迟队列 | 1. 时间精度高<br>2. 解耦系统<br>3. 高可靠 | 1. 需要引入消息队列<br>2. 额外维护成本 |

### 2.2 延迟队列的优势

1. **解耦系统**：订单创建和超时处理逻辑解耦
2. **高可靠**：消息持久化，防止消息丢失
3. **时间精确**：可以精确控制消息的延迟时间
4. **削峰填谷**：有效处理流量峰值
5. **扩展性好**：可以轻松扩展其他延迟业务需求

## 3. RabbitMQ延迟队列插件安装

### 3.1 插件介绍

RabbitMQ Delayed Message Exchange 插件是一个官方维护的插件，它提供了一个延迟交换机类型 `x-delayed-message`，允许消息根据指定的延迟时间进行投递。

### 3.2 插件安装

从项目结构可以看到，插件已经放置在 `rabbitmq/plugins` 目录下：

```
rabbitmq/
└── plugins/
    └── rabbitmq_delayed_message_exchange-4.1.0.ez
```

在Docker环境中，通常需要在 `docker-compose.yml` 中配置启用该插件。

## 4. 项目中的延迟队列实现

### 4.1 核心组件设计

项目中实现延迟队列处理订单超时主要包含以下几个核心组件：

1. **RabbitMQ客户端**：封装了与RabbitMQ交互的核心功能
2. **订单超时事件发布**：在订单创建时发布延迟消息
3. **订单超时事件消费**：处理超时消息，执行订单取消操作
4. **订单状态更新**：更新订单状态为已取消
5. **库存返还**：取消订单后返还商品库存

### 4.2 RabbitMQ客户端封装

项目在 `utility/rabbitmq/rabbitmq.go` 中封装了RabbitMQ客户端，提供了连接管理、消息发布、消费等功能。

```go
// 关键方法：PublishWithDelay 发布延迟消息
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
```

特别注意：
- 使用 `Headers: amqp.Table{"x-delay": delayMs}` 设置延迟时间
- 设置 `DeliveryMode: amqp.Persistent` 确保消息持久化，防止服务重启导致消息丢失

### 4.3 延迟交换机声明

```go
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
```

延迟交换机需要指定 `kind` 为 `x-delayed-message`，并在 `args` 中设置 `x-delayed-type` 参数。

## 5. 订单超时处理流程实现

### 5.1 订单超时事件定义

```go
// 订单超时事件定义
type OrderTimeoutEvent struct {
    OrderId   int    `json:"order_id"`
    Type      string `json:"type"`
    TimeStamp string `json:"timestamp"`
}

// 事件类型常量
const (
    OrderTimeout = "order_timeout"
)
```

### 5.2 发布订单超时事件

当用户创建订单时，系统会发布一个延迟消息，设置一定的延迟时间（如30分钟）：

```go
// PublishOrderTimeoutEvent 发布订单超时事件
func PublishOrderTimeoutEvent(orderId int, delayMs int) {
    ctx := context.Background()

    // 初始化RabbitMQ连接
    rb, err := NewRabbitMQ(ctx)
    if err != nil {
        g.Log().Errorf(ctx, "Failed to connect to RabbitMQ: %v", err)
        return
    }
    defer rb.Close()

    // 声明延迟交换机
    exchange := g.Cfg().MustGet(ctx, "rabbitmq.exchange.orderDelayExchange").String()
    err = rb.DeclareExchange(exchange, "x-delayed-message")
    if err != nil {
        g.Log().Errorf(ctx, "Failed to declare delay exchange: %v", err)
        return
    }

    // 创建事件
    event := OrderTimeoutEvent{
        OrderId:   orderId,
        Type:      OrderTimeout,
        TimeStamp: time.Now().Format(time.RFC3339),
    }

    // 发布延迟事件
    routingKey := g.Cfg().MustGet(ctx, "rabbitmq.routingKey.orderTimeout").String()
    err = rb.PublishWithDelay(exchange, routingKey, event, delayMs)
    if err != nil {
        g.Log().Errorf(ctx, "Failed to publish orderTimeout event: %v", err)
    } else {
        g.Log().Infof(ctx, "Published orderTimeout event with %d ms delay: %+v", delayMs, event)
    }
}
```

### 5.3 订单超时消费者

订单超时消费者负责接收和处理超时消息：

```go
// OrderTimeoutConsumer 订单超时未支付消费者
type OrderTimeoutConsumer struct {
    *rabbitmq.BaseConsumer
}

// NewOrderTimeoutConsumer 创建订单超时未支付消费者
func NewOrderTimeoutConsumer(ctx context.Context) *OrderTimeoutConsumer {
    config := rabbitmq.ConsumerConfig{
        Exchange:      g.Cfg().MustGet(ctx, "rabbitmq.exchange.orderDelayExchange").String(),
        ExchangeType:  "x-delayed-message",
        Queue:         g.Cfg().MustGet(ctx, "rabbitmq.queue.orderTimeoutQueue").String(),
        RoutingKey:    g.Cfg().MustGet(ctx, "rabbitmq.routingKey.orderTimeout").String(),
        ConsumerTag:   "DEMO_WECHAT_OPEN_ID",
        AutoAck:       false,
        PrefetchCount: 1,
        Durable:       true,
    }

    return &OrderTimeoutConsumer{
        BaseConsumer: rabbitmq.NewBaseConsumer("OrderTimeoutConsumer", config),
    }
}

// HandleMessage 处理订单超时未支付消息
func (c *OrderTimeoutConsumer) HandleMessage(ctx context.Context, msg amqp.Delivery) error {
    var event rabbitmq.OrderTimeoutEvent
    err := rabbitmq.UnmarshalEvent(msg.Body, &event)
    if err != nil {
        g.Log().Errorf(ctx, "解析订单超时未支付结果事件失败: %v", err)
        return err
    }
    g.Log().Infof(ctx, "收到订单超时未支付事件: %+v", event)
    if event.Type != rabbitmq.OrderTimeout {
        g.Log().Errorf(ctx, "不是订单超时未支付的事件,event.Type:%s", event.Type)
        return gerror.WrapCode(gcode.CodeInvalidParameter, fmt.Errorf("不是订单超时未支付的事件,event.Type:%s", event.Type))
    }
    eventTime, err := time.Parse(time.RFC3339, event.TimeStamp)
    if err != nil {
        return fmt.Errorf("解析事件时间戳失败: %v", err)
    }

    // 判断是否过期：事件时间 + 30s < 当前时间
    expireTime := g.Cfg().MustGet(ctx, "business.orderTimeout").String()
    expireMs, err := strconv.Atoi(expireTime)
    if err != nil {
        return fmt.Errorf("订单超时时间配置无效: %v", err)
    }
    expireDuration := time.Duration(expireMs) * time.Millisecond
    if time.Now().Before(eventTime.Add(expireDuration)) {
        g.Log().Infof(ctx, "订单未到取消时间，跳过处理: order_id=%d, event_time=%s", event.OrderId, event.TimeStamp)
        return nil
    }

    // 调用订单超时未支付处理逻辑
    err = order_info.HandleOrderTimeoutResult(ctx, event.OrderId)
    if err != nil {
        g.Log().Errorf(ctx, "处理订单 %d 的超时未支付失败: %v", event.OrderId, err)
        return err
    }
    g.Log().Infof(ctx, "成功处理订单 %d 的超时未支付事件", event.OrderId)

    // 取消库存
    eventReq, err := order_info.GetOrderDetail(ctx, event.OrderId)
    if err != nil {
        g.Log().Errorf(ctx, "获取订单 %v 对应的商品信息失败,err: %v", event.OrderId, err)
        return err
    }
    go rabbitmq.PublishReturnStockEvent(event.OrderId, eventReq)

    return nil
}
```

消费者的主要职责：
1. 解析订单超时事件消息
2. 验证事件类型和时间
3. 调用订单超时处理逻辑
4. 触发库存返还操作

### 5.4 订单超时处理逻辑

```go
// HandleOrderTimeoutResult 处理订单超时结果
func HandleOrderTimeoutResult(ctx context.Context, orderId int) error {
    // 更新字段
    updateData := g.Map{
        "status":     consts.OrderStatusCancelled,
        "updated_at": gtime.Now(), // 可选：更新时间戳
    }
    // 更新订单状态
    result, err := dao.OrderInfo.Ctx(ctx).Where("id=? AND status=?", orderId, consts.OrderStatusPendingPayment).Update(updateData)
    if err != nil {
        return gerror.WrapCode(gcode.CodeDbOperationError, err)
    }

    row, _ := result.RowsAffected()
    if row == 0 {
        g.Log().Infof(ctx, "订单已取消，无需再取消, orderId=%d", orderId)
        return nil
    }

    g.Log().Infof(ctx, "订单状态更新成功, 订单编号:{%s}, 新状态: %d", orderId, consts.OrderStatusPendingPayment)
    return nil
}
```

这个函数的主要逻辑：
1. 准备更新数据，设置订单状态为已取消
2. 使用 `WHERE id=? AND status=?` 条件进行乐观锁更新，确保只更新待支付状态的订单
3. 检查更新结果，记录日志

## 6. 完整业务流程

### 6.1 流程图

```
┌───────────────┐      ┌────────────────────┐      ┌──────────────────────┐
│  创建订单     │ ──>  │  发布延迟消息      │ ──>  │  延迟交换机存储      │
└───────────────┘      └────────────────────┘      └──────────┬─────────┘
                                                             │ 延迟时间到
                                                             ▼
┌───────────────────────┐      ┌───────────────────────┐      ┌─────────────────┐
│  返还商品库存         │ <─── │  更新订单状态为已取消 │ <─── │  消费超时消息   │
└───────────────────────┘      └───────────────────────┘      └─────────────────┘
```

### 6.2 流程步骤详解

1. **订单创建**：用户提交订单，系统创建订单记录，状态为"待支付"
2. **发布延迟消息**：调用 `PublishOrderTimeoutEvent` 方法，发布一个延迟消息，延迟时间通常设置为订单超时时间（如30分钟）
3. **消息存储**：延迟消息被发送到延迟交换机并存储
4. **消息延迟**：消息在延迟交换机中等待，直到延迟时间到期
5. **消息路由**：延迟时间到期后，消息被路由到订单超时队列
6. **消息消费**：订单超时消费者 `OrderTimeoutConsumer` 从队列中获取消息
7. **订单状态检查**：验证订单是否仍然是"待支付"状态
8. **更新订单状态**：调用 `HandleOrderTimeoutResult` 更新订单状态为"已取消"
9. **返还库存**：调用 `PublishReturnStockEvent` 发布库存返还事件

## 7. 代码优化建议

### 7.1 错误处理优化

在 `HandleMessage` 方法中，建议添加 `msg.Ack(false)` 和 `msg.Nack(false, false)` 确保消息正确确认或拒绝：

```go
func (c *OrderTimeoutConsumer) HandleMessage(ctx context.Context, msg amqp.Delivery) error {
    defer func() {
        if err := recover(); err != nil {
            g.Log().Errorf(ctx, "处理订单超时消息发生 panic: %v", err)
            msg.Nack(false, false) // 拒绝消息，不重新入队
        }
    }()
    
    // 原有代码...
    
    // 成功处理后确认消息
    msg.Ack(false)
    return nil
}
```

### 7.2 超时时间配置优化

建议将超时时间配置从字符串改为数字类型，避免每次转换：

```go
// 配置文件中
// business:
//   orderTimeoutMs: 1800000  # 30分钟

// 代码中直接获取整数
var orderTimeoutMs int
orderTimeoutMs = g.Cfg().MustGet(ctx, "business.orderTimeoutMs").Int()
```

### 7.3 消息幂等性处理

为了防止重复处理订单超时消息，可以添加幂等性处理：

```go
func HandleOrderTimeoutResult(ctx context.Context, orderId int) error {
    // 添加分布式锁防止并发处理
    lockKey := fmt.Sprintf("order:timeout:lock:%d", orderId)
    unlock, err := redis.SetNX(ctx, lockKey, time.Now().Unix(), 30*time.Second)
    if err != nil || !unlock {
        g.Log().Infof(ctx, "订单 %d 正在处理中，跳过", orderId)
        return nil
    }
    defer func() {
        if unlock {  // 确保只在获取锁成功的情况下释放锁
            redis.Del(ctx, lockKey)
        }
    }()
    
    // 原有更新订单状态的代码...
}
```

### 7.4 监控和告警

建议添加关键节点的监控和告警：

1. 订单超时消息发布失败监控
2. 订单超时处理失败监控
3. 库存返还失败监控
4. 订单超时率统计

## 8. 总结

### 8.1 核心优势

1. **高可靠性**：消息持久化、指数退避重试等机制确保消息不丢失
2. **精确控制**：可以精确控制订单超时时间
3. **系统解耦**：订单创建和超时处理逻辑完全解耦
4. **可扩展性**：相同的模式可以应用于其他需要延迟处理的场景

### 8.2 学习要点

1. **延迟队列概念**：理解延迟队列的基本原理和应用场景
2. **RabbitMQ插件使用**：掌握 RabbitMQ Delayed Message Exchange 插件的配置和使用
3. **消息持久化**：理解消息持久化的重要性和配置方式
4. **消费者实现**：学习如何实现高可靠的消息消费者
5. **幂等性处理**：理解并实现幂等性处理，避免重复操作

### 8.3 应用场景扩展

除了订单超时处理，延迟队列还可以用于以下场景：

1. **预约提醒**：用户预约某服务前的提醒通知
2. **会员到期提醒**：会员到期前的自动提醒
3. **定时任务**：不需要高精度的定时任务调度
4. **异步任务补偿**：失败任务的延迟重试
5. **优惠券过期通知**：优惠券即将过期的提醒

通过本实战案例，相信大家已经掌握了如何使用RabbitMQ延迟队列来处理订单超时问题，以及相关的最佳实践和优化方向。