package consumer

import (
	"context"
	"shop-goframe-micro-service-refacotor/app/order/internal/consts"
	"shop-goframe-micro-service-refacotor/app/order/internal/dao"
	order_logic "shop-goframe-micro-service-refacotor/app/order/internal/logic/order_info"
	"shop-goframe-micro-service-refacotor/app/order/internal/model/entity"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	amqp "github.com/rabbitmq/amqp091-go"
)

type OrderPaidSalesConsumer struct {
	*rabbitmq.BaseConsumer
}

func NewOrderPaidSalesConsumer(ctx context.Context) *OrderPaidSalesConsumer {
	config := rabbitmq.ConsumerConfig{
		Exchange:      g.Cfg().MustGet(ctx, "rabbitmq.exchange.orderExchange").String(),
		ExchangeType:  "topic",
		Queue:         g.Cfg().MustGet(ctx, "rabbitmq.queue.orderPaidSalesQueue").String(),
		RoutingKey:    g.Cfg().MustGet(ctx, "rabbitmq.routingKey.orderPaid").String(),
		ConsumerTag:   "order_paid_sales",
		PrefetchCount: 1,
		MaxRetries:    3,
	}
	return &OrderPaidSalesConsumer{
		BaseConsumer: rabbitmq.NewBaseConsumer("OrderPaidSalesConsumer", config),
	}
}

// 同一个 order_number 的消息会有业务幂等键。
// Redis/消息幂等只能辅助，真正兜底必须靠业务状态。
func (c *OrderPaidSalesConsumer) GetBusinessID(data []byte, event map[string]interface{}) string {
	if v, ok := event["order_number"]; ok {
		return gconv.String(v)
	}
	return ""
}

/*
HandleMessage 的业务逻辑：处理订单支付成功的消息

	解析 OrderPaidEvent
	查 order_info
	如果订单不存在：记录日志，Ack
	如果订单不是已支付：记录日志，Ack
	如果 sales_status = 已同步：说明重复消息，直接 Ack
	如果 sales_status = 未同步：抢占为同步中，然后执行 IncreaseOrderGoodsSales
	成功：同步中 -> 已同步（销量已经增加成功时，以已同步为最终状态）
	失败：同步中 -> 同步失败
*/
func (c *OrderPaidSalesConsumer) HandleMessage(ctx context.Context, msg amqp.Delivery) error {
	// 解析 OrderPaidEvent
	var event rabbitmq.OrderPaidEvent
	err := rabbitmq.UnmarshalEvent(msg.Body, &event)
	if err != nil {
		g.Log().Errorf(ctx, "解析订单支付成功事件失败: %v", err)
		return err
	}

	if event.OrderNumber == "" {
		g.Log().Warningf(ctx, "订单支付成功事件缺少 order_number: %+v", event)
		return nil
	}
	g.Log().Infof(ctx, "收到订单支付成功事件: %+v", event)

	// 查 order_info
	var order entity.OrderInfo
	err = dao.OrderInfo.Ctx(ctx).Where(dao.OrderInfo.Columns().Number, event.OrderNumber).Scan(&order)
	if err != nil {
		g.Log().Errorf(ctx, "查询订单失败: %v", err)
		return err
	}
	if order.Id == 0 {
		g.Log().Warningf(ctx, "订单不存在，跳过支付成功事件, order=%s", event.OrderNumber)
		return nil
	}

	if order.Status != int(consts.OrderStatusPaid) {
		g.Log().Warningf(ctx, "订单不是已支付状态，跳过销量同步, order=%s, status=%d", event.OrderNumber, order.Status)
		return nil
	}

	if order.SalesStatus == int(consts.OrderSalesStatusSynced) {
		g.Log().Infof(ctx, "订单销量已同步，重复消息直接跳过, order=%s", event.OrderNumber)
		return nil
	}

	// 尝试更新订单销量的状态
	success, err := order_logic.TryUpdateOrderSalesStatus(ctx, order.Number, consts.OrderSalesStatusPending, consts.OrderSalesStatusSyncing)
	if err != nil {
		g.Log().Errorf(ctx, "抢占订单销量同步任务失败, order=%s, err=%v", event.OrderNumber, err)
		return err
	}
	if !success {
		g.Log().Infof(ctx, "订单销量同步任务未抢占成功，可能已被处理, order=%s, sales_status=%d", event.OrderNumber, order.SalesStatus)
		return nil
	}

	// 增加订单商品的销量
	err = order_logic.IncreaseOrderGoodsSales(ctx, order.Number)
	if err != nil {
		g.Log().Errorf(ctx, "增加订单销量失败, 订单编号: %s, 错误: %v", order.Number, err)
		_, updateErr := order_logic.TryUpdateOrderSalesStatus(ctx, order.Number, consts.OrderSalesStatusSyncing, consts.OrderSalesStatusFailed)
		if updateErr != nil {
			g.Log().Errorf(ctx, "标记订单销量同步失败状态失败: %v", updateErr)
		}
		return nil
	}

	if updateErr := order_logic.UpdateOrderSalesStatusByNumber(ctx, order.Number, consts.OrderSalesStatusSynced); updateErr != nil {
		g.Log().Errorf(ctx, "标记订单销量已同步状态失败, 订单编号: %s, 错误: %v", order.Number, updateErr)
		return nil
	}
	g.Log().Infof(ctx, "订单销量已同步, 订单编号: %s", order.Number)
	return nil
}
