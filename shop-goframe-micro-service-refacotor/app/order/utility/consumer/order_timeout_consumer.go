package consumer

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
	order_info "shop-goframe-micro-service-refacotor/app/order/internal/logic/order_info"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
	"strconv"
	"time"
)

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
		ConsumerTag:   "order_service_order_timeout",
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
