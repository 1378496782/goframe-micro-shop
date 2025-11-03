package consumer

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
	"shop-goframe-micro-service-refacotor/app/goods/internal/logic/goods_info"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
)

// OrderCreatedConsumer 订单创建事件消费者
type OrderCreatedConsumer struct {
	*rabbitmq.BaseConsumer
}

// NewOrderCreatedConsumer 创建订单创建事件消费者
func NewOrderCreatedConsumer(ctx context.Context) *OrderCreatedConsumer {
	config := rabbitmq.ConsumerConfig{
		Exchange:     g.Cfg().MustGet(ctx, "rabbitmq.exchange.orderExchange").String(),
		ExchangeType: "topic",
		Queue:        g.Cfg().MustGet(ctx, "rabbitmq.queue.orderCreatedQueue").String(),
		RoutingKey:   g.Cfg().MustGet(ctx, "rabbitmq.routingKey.orderCreated").String(),
		ConsumerTag:  "goods_service_order_created",
	}
	return &OrderCreatedConsumer{
		BaseConsumer: rabbitmq.NewBaseConsumer("OrderCreatedConsumer", config),
	}
}

func (c *OrderCreatedConsumer) HandleMessage(ctx context.Context, msg amqp.Delivery) error {
	var event rabbitmq.OrderCreatedEvent
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		g.Log().Errorf(ctx, "Failed to unmarshal OrderCreatedEvent: %v", err)
		return err // Nack the message
	}

	g.Log().Infof(ctx, "收到创建订单事件OrderCreatedEvent: %+v", event)

	if len(event.GoodsIds) == 0 {
		g.Log().Infof(ctx, "No goods to delete from cart for user %d", event.UserId)
		return nil // Ack the message
	}

	// Delete items from cart
	_, err := dao.CartInfo.Ctx(ctx).Where("user_id", event.UserId).WhereIn("goods_id", event.GoodsIds).Delete()
	if err != nil {
		g.Log().Errorf(ctx, "Failed to delete cart items for user %d: %v", event.UserId, err)
		return err // Nack the message
	}

	// 扣减库存
	// todo: 添加最大重试次数
	err = goods_info.ReduceStock(ctx, &event)
	if err != nil {
		//go rabbitmq.PublishOrderCreatedEvent(event)
		//g.Log().Warningf(ctx, "库存扣减失败，已重新发布重试事件，失败原因: %v", err)
		return err // Nack the message
	}

	g.Log().Infof(ctx, "成功从订单{%d}中扣减库存", event.OrderId)
	return nil // Ack the message
}
