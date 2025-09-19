package consumer

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/streadway/amqp"
	"shop-goframe-micro-service-refacotor/app/goods/internal/logic/user_coupon_info"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
)

// CouponConfirmConsumer 优惠券确认消费者
type CouponConfirmConsumer struct {
	*rabbitmq.BaseConsumer
}

// NewCouponConfirmConsumer 创建优惠券确认消费者
func NewCouponConfirmConsumer(ctx context.Context) *CouponConfirmConsumer {
	config := rabbitmq.ConsumerConfig{
		Exchange:      g.Cfg().MustGet(ctx, "rabbitmq.exchange.couponConfirmExchange").String(),
		ExchangeType:  "topic",
		Queue:         g.Cfg().MustGet(ctx, "rabbitmq.queue.couponConfirmRequestQueue").String(),
		RoutingKey:    g.Cfg().MustGet(ctx, "rabbitmq.routingKey.couponConfirmRequest").String(),
		ConsumerTag:   "goods_service_coupon",
		AutoAck:       false,
		PrefetchCount: 1,
		Durable:       true,
	}

	return &CouponConfirmConsumer{
		BaseConsumer: rabbitmq.NewBaseConsumer("CouponConfirmConsumer", config),
	}
}

// HandleMessage 处理优惠券确认消息
func (c *CouponConfirmConsumer) HandleMessage(ctx context.Context, msg amqp.Delivery) error {
	var event rabbitmq.CouponConfirmEvent
	err := rabbitmq.UnmarshalEvent(msg.Body, &event)
	if err != nil {
		g.Log().Errorf(ctx, "解析优惠券确认事件失败: %v", err)
		return err
	}

	g.Log().Infof(ctx, "收到优惠券确认事件: %+v", event)

	// 调用优惠券确认逻辑
	err = user_coupon_info.HandleOrderConfirmMessage(ctx, event.OrderId, event.UserID, event.CouponId)
	if err != nil {
		g.Log().Errorf(ctx, "处理用户 %d 的优惠券确认失败: %v", event.UserID, err)
		return err
	}

	g.Log().Infof(ctx, "成功处理用户 %d 的优惠券确认", event.UserID)
	return nil
}
