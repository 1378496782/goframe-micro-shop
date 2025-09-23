package consumer

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/streadway/amqp"
	order_info "shop-goframe-micro-service-refacotor/app/order/internal/logic/order_info"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
)

// CouponResultConsumer 优惠券确认结果消费者
type CouponResultConsumer struct {
	*rabbitmq.BaseConsumer
}

// NewCouponResultConsumer 创建优惠券确认结果消费者
func NewCouponResultConsumer(ctx context.Context) *CouponResultConsumer {
	config := rabbitmq.ConsumerConfig{
		Exchange:      g.Cfg().MustGet(ctx, "rabbitmq.exchange.couponConfirmResultExchange").String(),
		ExchangeType:  "topic",
		Queue:         g.Cfg().MustGet(ctx, "rabbitmq.queue.couponConfirmResultQueue").String(),
		RoutingKey:    g.Cfg().MustGet(ctx, "rabbitmq.routingKey.couponConfirmResult").String(),
		ConsumerTag:   "DEMO_WECHAT_OPEN_ID",
		AutoAck:       false,
		PrefetchCount: 1,
		Durable:       true,
	}

	return &CouponResultConsumer{
		BaseConsumer: rabbitmq.NewBaseConsumer("CouponResultConsumer", config),
	}
}

// HandleMessage 处理优惠券确认结果消息
func (c *CouponResultConsumer) HandleMessage(ctx context.Context, msg amqp.Delivery) error {
	var event rabbitmq.CouponConfirmResultEvent
	err := rabbitmq.UnmarshalEvent(msg.Body, &event)
	if err != nil {
		g.Log().Errorf(ctx, "解析优惠券确认结果事件失败: %v", err)
		return err
	}

	g.Log().Infof(ctx, "收到优惠券确认结果事件: %+v", event)

	// 调用优惠券确认结果处理逻辑
	err = order_info.HandleCouponResult(ctx, event.OrderId, event.Success, event.Message)
	if err != nil {
		g.Log().Errorf(ctx, "处理订单 %d 的优惠券确认结果失败: %v", event.OrderId, err)
		return err
	}

	g.Log().Infof(ctx, "成功处理订单 %d 的优惠券确认结果", event.OrderId)
	return nil
}