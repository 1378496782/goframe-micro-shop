package consumer

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/streadway/amqp"
	"shop-goframe-micro-service-refacotor/app/goods/internal/logic/user_coupon_info"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
)

// StartUserRegisteredConsumer 启动用户注册事件消费者
func StartUserRegisteredConsumer() {
	ctx := context.Background()

	// 初始化RabbitMQ连接
	rb, err := rabbitmq.NewRabbitMQ(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer rb.Close()

	// 声明交换机和队列
	exchange := g.Cfg().MustGet(ctx, "rabbitmq.exchange.user").String()
	queueName := g.Cfg().MustGet(ctx, "rabbitmq.queue.user.registered").String()
	routingKey := g.Cfg().MustGet(ctx, "rabbitmq.routingKey.user.registered").String()

	err = rb.DeclareExchange(exchange, "topic")
	if err != nil {
		g.Log().Errorf(ctx, "Failed to declare exchange: %v", err)
		return
	}

	q, err := rb.DeclareQueue(queueName)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to declare queue: %v", err)
		return
	}

	err = rb.QueueBind(q.Name, routingKey, exchange)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to bind queue: %v", err)
		return
	}

	// 开始消费消息
	msgs, err := rb.Consume(q.Name, "goods_service", false)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to consume messages: %v", err)
		return
	}

	g.Log().Info(ctx, "Started user registered event consumer")

	for msg := range msgs {
		go handleUserRegisteredEvent(ctx, msg)
	}
}

// 处理用户注册事件
func handleUserRegisteredEvent(ctx context.Context, msg amqp.Delivery) {
	var event rabbitmq.UserRegisteredEvent
	err := json.Unmarshal(msg.Body, &event)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to unmarshal event: %v", err)
		msg.Nack(false, false) // 拒绝消息，不重新入队
		return
	}

	g.Log().Infof(ctx, "Received user registered event: %+v", event)

	// 调用发放优惠券逻辑
	err = user_coupon_info.IssueCouponToUser(ctx, event.UserID)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to issue coupon for user %d: %v", event.UserID, err)
		msg.Nack(false, true) // 拒绝消息，重新入队
		return
	}

	g.Log().Infof(ctx, "Successfully issued coupon for user %d", event.UserID)
	msg.Ack(false) // 确认消息
}
