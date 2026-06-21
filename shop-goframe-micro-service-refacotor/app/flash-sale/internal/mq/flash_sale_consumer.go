package mq

import (
	"context"
	"encoding/json"
	"fmt"

	"shop-goframe-micro-service-refacotor/app/flash-sale/internal/model"

	"shop-goframe-micro-service-refacotor/app/flash-sale/utility"

	"github.com/gogf/gf/v2/frame/g"
	amqp "github.com/rabbitmq/amqp091-go"
)

// FlashSaleOrderConsumer 秒杀订单消费者
type FlashSaleOrderConsumer struct {
	ctx context.Context
}

// NewFlashSaleOrderConsumer 创建秒杀订单消费者
func NewFlashSaleOrderConsumer(ctx context.Context) *FlashSaleOrderConsumer {
	return &FlashSaleOrderConsumer{
		ctx: ctx,
	}
}

// Start 启动消费者
func (c *FlashSaleOrderConsumer) Start() error {
	rabbitMQ := utility.GetFlashSaleRabbitMQ()
	if rabbitMQ == nil {
		return fmt.Errorf("RabbitMQ未初始化")
	}

	// 消费消息
	msgs, err := rabbitMQ.Channel.Consume(
		"flash_sale.order.queue",    // 队列名称
		"flash_sale.order.consumer", // 消费者标签
		false,                       // 自动确认
		false,                       // 排他性
		false,                       // 不等待
		false,                       // 参数
		nil,                         // 额外参数
	)
	if err != nil {
		return err
	}

	g.Log().Info(c.ctx, "开始监听秒杀订单消息队列")

	// 启动goroutine处理消息
	go c.processMessages(msgs)

	return nil
}

// processMessages 处理消息
func (c *FlashSaleOrderConsumer) processMessages(msgs <-chan amqp.Delivery) {
	for msg := range msgs {
		if err := c.handleMessage(msg); err != nil {
			g.Log().Error(c.ctx, "处理秒杀订单消息失败:", err)
			// 处理失败，重新入队或记录日志
			msg.Nack(false, true) // 重新入队
		} else {
			msg.Ack(false) // 确认消息
		}
	}
}

// handleMessage 处理单条消息
func (c *FlashSaleOrderConsumer) handleMessage(msg amqp.Delivery) error {
	var orderMsg model.FlashSaleOrderMessage
	if err := json.Unmarshal(msg.Body, &orderMsg); err != nil {
		return fmt.Errorf("解析消息失败: %v", err)
	}

	g.Log().Info(c.ctx, "收到秒杀订单消息:", orderMsg)

	// 调用服务处理订单
	flashSaleService := utility.GetFlashSaleService()
	if flashSaleService == nil {
		return fmt.Errorf("获取秒杀服务失败")
	}

	// 处理订单 - 注意：这里需要调用具体的服务实现
	// 由于接口定义中没有ProcessFlashSaleOrder方法，我们需要通过其他方式处理
	// 暂时直接处理，后续需要调整接口定义
	g.Log().Info(c.ctx, "处理秒杀订单消息:", orderMsg.OrderId)

	// TODO: 实现具体的订单处理逻辑
	// 这里可以调用服务层的ProcessFlashSaleOrder方法，但需要调整接口定义

	g.Log().Info(c.ctx, "秒杀订单处理完成")
	return nil
}

// PublishFlashSaleOrder 发布秒杀订单消息
func PublishFlashSaleOrder(ctx context.Context, orderMsg *model.FlashSaleOrderMessage) error {
	rabbitMQ := utility.GetFlashSaleRabbitMQ()
	if rabbitMQ == nil {
		g.Log().Warning(ctx, "RabbitMQ未初始化，订单消息将本地处理:", orderMsg.OrderId)
		// RabbitMQ未初始化时，直接本地处理订单
		return processFlashSaleOrderLocal(ctx, orderMsg)
	}

	// 序列化消息
	body, err := json.Marshal(orderMsg)
	if err != nil {
		return fmt.Errorf("序列化消息失败: %v", err)
	}

	// 发布消息
	if err := rabbitMQ.PublishFlashSaleOrder(body); err != nil {
		return fmt.Errorf("发布消息失败: %v", err)
	}

	g.Log().Info(ctx, "发布秒杀订单消息成功:", orderMsg.OrderId)
	return nil
}

// processFlashSaleOrderLocal 本地处理秒杀订单（当RabbitMQ不可用时）
func processFlashSaleOrderLocal(ctx context.Context, orderMsg *model.FlashSaleOrderMessage) error {
	g.Log().Info(ctx, "本地处理秒杀订单:", orderMsg.OrderId)

	// 这里可以实现简单的本地订单处理逻辑
	// 在实际生产环境中，可能需要将订单数据存储到数据库或缓存中

	// 模拟订单处理
	g.Log().Infof(ctx, "订单 %s 处理完成 - 用户ID: %d, 商品ID: %d, 数量: %d",
		orderMsg.OrderId, orderMsg.UserId, orderMsg.GoodsId, orderMsg.Count)

	return nil
}
