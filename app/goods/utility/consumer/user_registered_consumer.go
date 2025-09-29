package consumer

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
	"shop-goframe-micro-service-refacotor/app/goods/internal/logic/user_coupon_info"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
)

// UserRegisteredConsumer 用户注册事件消费者
type UserRegisteredConsumer struct {
	*rabbitmq.BaseConsumer
}

// NewUserRegisteredConsumer 创建用户注册事件消费者
func NewUserRegisteredConsumer(ctx context.Context) *UserRegisteredConsumer {
	config := rabbitmq.ConsumerConfig{
		Exchange:      g.Cfg().MustGet(ctx, "rabbitmq.exchange.user").String(),
		ExchangeType:  "topic",
		Queue:         g.Cfg().MustGet(ctx, "rabbitmq.queue.user.registered").String(),
		RoutingKey:    g.Cfg().MustGet(ctx, "rabbitmq.routingKey.user.registered").String(),
		ConsumerTag:   "goods_service_user",
		AutoAck:       false,
		PrefetchCount: 1,
		Durable:       true,
	}

	return &UserRegisteredConsumer{
		BaseConsumer: rabbitmq.NewBaseConsumer("UserRegisteredConsumer", config),
	}
}

// HandleMessage 处理用户注册事件消息
func (c *UserRegisteredConsumer) HandleMessage(ctx context.Context, msg amqp.Delivery) error {
	var event rabbitmq.UserRegisteredEvent
	err := rabbitmq.UnmarshalEvent(msg.Body, &event)
	if err != nil {
		g.Log().Errorf(ctx, "解析用户注册事件失败: %v", err)
		return err
	}

	g.Log().Infof(ctx, "收到用户注册事件: %+v", event)

	// 调用发放优惠券逻辑
	err = user_coupon_info.IssueCouponToUser(ctx, event.UserID)
	if err != nil {
		g.Log().Errorf(ctx, "为用户 %d 发放优惠券失败: %v", event.UserID, err)
		return err
	}

	g.Log().Infof(ctx, "成功为用户 %d 发放优惠券", event.UserID)
	return nil
}
