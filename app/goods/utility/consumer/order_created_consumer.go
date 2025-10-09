package consumer

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
	"shop-goframe-micro-service-refacotor/app/goods/internal/dao"
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

	g.Log().Infof(ctx, "Received OrderCreatedEvent: %+v", event)

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

	g.Log().Infof(ctx, "Successfully deleted %d goods from cart for user %d", len(event.GoodsIds), event.UserId)
	return nil // Ack the message
}
