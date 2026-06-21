package consumer

import (
	"context"
	"database/sql"
	"errors"
	goods_info "shop-goframe-micro-service-refacotor/app/goods/api/goods_info/v1"
	"shop-goframe-micro-service-refacotor/app/order/internal/consts"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
	"shop-goframe-micro-service-refacotor/app/order/internal/model/entity"
	goods "shop-goframe-micro-service-refacotor/app/order/utility/goods_info"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderCancelledConsumer struct {
	*rabbitmq.BaseConsumer
}

func NewOrderCancelledConsumer(ctx context.Context) *OrderCancelledConsumer {
	config := rabbitmq.ConsumerConfig{
		Exchange:      g.Cfg().MustGet(ctx, "rabbitmq.exchange.orderCancelledExchange").String(),
		ExchangeType:  "topic",
		Queue:         g.Cfg().MustGet(ctx, "rabbitmq.queue.orderCancelledQueue").String(),
		RoutingKey:    g.Cfg().MustGet(ctx, "rabbitmq.routingKey.orderCancelled").String(),
		ConsumerTag:   "order_cancelled",
		PrefetchCount: 1,
		MaxRetries:    3,
	}
	return &OrderCancelledConsumer{
		BaseConsumer: rabbitmq.NewBaseConsumer("OrderCancelledConsumer", config),
	}
}

// 同一个 order_number 的消息会有业务幂等键。
// Redis/消息幂等只能辅助，真正兜底必须靠业务状态。
func (c *OrderCancelledConsumer) GetBusinessID(data []byte, event map[string]interface{}) string {
	if v, ok := event["order_number"]; ok {
		return gconv.String(v)
	}
	return ""
}

func (c *OrderCancelledConsumer) HandleMessage(ctx context.Context, msg amqp.Delivery) error {
	// 解析 OrderCancelledEvent
	var event rabbitmq.OrderCancelledEvent
	err := rabbitmq.UnmarshalEvent(msg.Body, &event)
	if err != nil {
		g.Log().Errorf(ctx, "解析订单取消事件失败: %v", err)
		return err
	}

	if event.OrderNumber == "" {
		g.Log().Warningf(ctx, "订单取消事件缺少 order_number: %+v", event)
		return nil
	}
	g.Log().Infof(ctx, "收到订单取消事件: %+v", event)

	// 查 order_info
	var order entity.OrderInfo
	err = dao.OrderInfo.Ctx(ctx).Where(dao.OrderInfo.Columns().Number, event.OrderNumber).Scan(&order)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			g.Log().Warningf(ctx, "订单不存在，跳过取消事件, order=%s", event.OrderNumber)
			return nil
		}
		g.Log().Errorf(ctx, "查询订单失败: %v", err)
		return err
	}
	if order.Id == 0 {
		g.Log().Warningf(ctx, "订单不存在，跳过取消事件, order=%s", event.OrderNumber)
		return nil
	}

	if order.Status != int(consts.OrderStatusCancelled) {
		g.Log().Warningf(ctx, "订单不是已取消状态，跳过取消事件, order=%s, status=%d", event.OrderNumber, order.Status)
		return nil
	}

	// 查 order_goods_info 商品快照
	var goodsList []*entity.OrderGoodsInfo
	err = dao.OrderGoodsInfo.Ctx(ctx).Where("order_id", order.Id).Scan(&goodsList)
	if err != nil {
		g.Log().Errorf(ctx, "查询订单商品失败, order=%s, err=%v", event.OrderNumber, err)
		return err
	}
	if len(goodsList) == 0 {
		g.Log().Warningf(ctx, "订单商品为空，跳过恢复库存, order=%s", event.OrderNumber)
		return nil
	}
	// 恢复库存
	goodsIds := make([]uint32, 0, len(goodsList))
	counts := make([]uint32, 0, len(goodsList))
	for _, item := range goodsList {
		goodsIds = append(goodsIds, uint32(item.GoodsId))
		counts = append(counts, uint32(item.Count))
	}
	_, err = goods.Client.RestoreStock(ctx, &goods_info.RestoreStockReq{
		GoodsIds: goodsIds,
		Counts:   counts,
	})
	if err != nil {
		g.Log().Errorf(ctx, "恢复商品库存失败: order=%s, err=%v", event.OrderNumber, err)
		// 保持订单为已取消状态，返回错误交给 MQ 重试恢复库存。
		return err
	}

	g.Log().Infof(ctx, "订单取消库存恢复成功, order=%s", event.OrderNumber)
	return nil
}
