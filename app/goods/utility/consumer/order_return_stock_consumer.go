package consumer

import (
	"context"
	"encoding/json"
	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
	"shop-goframe-micro-service-refacotor/app/goods/internal/logic/goods_info"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"
)

// ReturnStockConsumer 订单库存返还事件消费者
type ReturnStockConsumer struct {
	*rabbitmq.BaseConsumer
}

// NewReturnStockConsumer 订单库存返还事件消费者
func NewReturnStockConsumer(ctx context.Context) *ReturnStockConsumer {
	config := rabbitmq.ConsumerConfig{
		Exchange:     g.Cfg().MustGet(ctx, "rabbitmq.exchange.goodsStockExchange").String(),
		ExchangeType: "topic",
		Queue:        g.Cfg().MustGet(ctx, "rabbitmq.queue.goodsStockQueue").String(),
		RoutingKey:   g.Cfg().MustGet(ctx, "rabbitmq.routingKey.goodsStock").String(),
		ConsumerTag:  "goods_service_order_goods_return_stock",
	}
	return &ReturnStockConsumer{
		BaseConsumer: rabbitmq.NewBaseConsumer("ReturnStockConsumer", config),
	}
}

func (c *ReturnStockConsumer) HandleMessage(ctx context.Context, msg amqp.Delivery) error {
	var event rabbitmq.OrderStockReturnEvent
	if err := json.Unmarshal(msg.Body, &event); err != nil {
		g.Log().Errorf(ctx, "反序列化 OrderStockReturnEvent 失败,err: %v", err)
		return err // Nack the message
	}

	g.Log().Infof(ctx, "OrderStockReturnEvent: %+v", event)

	if len(event.GoodsInfo) == 0 {
		g.Log().Errorf(ctx, "订单{%d} 没有商品需要返还库存", event.OrderId)
		return nil
	}

	goodsInfoArr, err := goods_info.ReturnStock(ctx, &event)
	if err != nil {
		g.Log().Infof(ctx, "返回库存失败", event.OrderId)
		return err
	}
	// todo 自动重试（带最大次数限制+延迟）、细分错误处理
	if len(goodsInfoArr) != 0 {
		//go rabbitmq.PublishReturnStockEvent(event.OrderId, goodsInfoArr)
		//g.Log().Warningf(ctx, "部分库存返还失败，已重新发布重试事件，失败商品数: %d", len(goodsInfoArr))
	}
	g.Log().Infof(ctx, "订单{%d} 返还库存成功", event.OrderId)
	return nil
}
