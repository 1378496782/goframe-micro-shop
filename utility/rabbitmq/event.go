package rabbitmq

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"time"
)

// UserRegisteredEvent 用户注册事件
type UserRegisteredEvent struct {
	UserID int `json:"user_id"`
}

// CouponIssuedEvent 优惠券发放事件
type CouponIssuedEvent struct {
	UserID int `json:"user_id"`
}

// PublishUserRegisteredEvent 发布用户注册事件
func PublishUserRegisteredEvent(userID int) {
	ctx := context.Background()

	// 初始化RabbitMQ连接
	rb, err := NewRabbitMQ(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer rb.Close()

	// 声明交换机
	exchange := g.Cfg().MustGet(ctx, "rabbitmq.exchange.user").String()
	err = rb.DeclareExchange(exchange, "topic")
	if err != nil {
		g.Log().Errorf(ctx, "Failed to declare exchange: %v", err)
		return
	}

	// 创建事件
	event := UserRegisteredEvent{
		UserID: userID,
	}

	// 发布事件
	routingKey := g.Cfg().MustGet(ctx, "rabbitmq.routingKey.user.registered").String()
	err = rb.Publish(exchange, routingKey, event)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to publish user registered event: %v", err)
	} else {
		g.Log().Infof(ctx, "Published user registered event: %+v", event)
	}
}

// CouponConfirmEvent 优惠券确认事件
type CouponConfirmEvent struct {
	OrderId  int `json:"order_id"`
	UserID   int `json:"user_id"`
	CouponId int `json:"coupon_id"`
}

// CouponConfirmResultEvent 优惠券确认结果事件
type CouponConfirmResultEvent struct {
	OrderId int    `json:"order_id"`
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// PublishCouponConfirmEvent 发布优惠券确认事件
func PublishCouponConfirmEvent(orderId int32, userID int32, couponId int32) {
	ctx := context.Background()

	// 初始化RabbitMQ连接
	rb, err := NewRabbitMQ(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer rb.Close()

	// 声明交换机
	exchange := g.Cfg().MustGet(ctx, "rabbitmq.exchange.couponConfirmExchange").String()
	err = rb.DeclareExchange(exchange, "topic")
	if err != nil {
		g.Log().Errorf(ctx, "Failed to declare exchange: %v", err)
		return
	}

	// 创建事件
	event := CouponConfirmEvent{
		OrderId:  int(orderId),
		UserID:   int(userID),
		CouponId: int(couponId),
	}

	// 发布事件
	routingKey := g.Cfg().MustGet(ctx, "rabbitmq.routingKey.couponConfirm").String()
	err = rb.Publish(exchange, routingKey, event)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to publish PublishCouponConfirmEvent event: %v", err)
	} else {
		g.Log().Infof(ctx, "Published PublishCouponConfirmEvent event: %+v", event)
	}
}

// PublishCouponConfirmResultEvent 优惠券确认结果事件
func PublishCouponConfirmResultEvent(orderId int, success bool, message string) {
	ctx := context.Background()

	// 初始化RabbitMQ连接
	rb, err := NewRabbitMQ(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer rb.Close()

	// 声明交换机
	exchange := g.Cfg().MustGet(ctx, "rabbitmq.exchange.couponConfirmResultExchange").String()
	err = rb.DeclareExchange(exchange, "topic")
	if err != nil {
		g.Log().Errorf(ctx, "Failed to declare exchange: %v", err)
		return
	}

	// 创建事件
	event := CouponConfirmResultEvent{
		OrderId: orderId,
		Success: success,
		Message: message,
	}

	// 发布事件
	routingKey := g.Cfg().MustGet(ctx, "rabbitmq.routingKey.couponConfirmResult").String()
	err = rb.Publish(exchange, routingKey, event)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to publish couponConfirmResult event: %v", err)
	} else {
		g.Log().Infof(ctx, "Published couponConfirmResult event: %+v", event)
	}
}

type OrderTimeoutEvent struct {
	OrderId   int    `json:"order_id"`
	Type      string `json:"type"`
	TimeStamp string `json:"timestamp"`
}

func PublishOrderTimeoutEvent(orderId int, delayMs int) {
	ctx := context.Background()

	// 初始化RabbitMQ连接
	rb, err := NewRabbitMQ(ctx)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to connect to RabbitMQ: %v", err)
		return
	}
	defer rb.Close()

	// 声明延迟交换机
	exchange := g.Cfg().MustGet(ctx, "rabbitmq.exchange.orderDelayExchange").String()
	err = rb.DeclareExchange(exchange, "x-delayed-message")
	if err != nil {
		g.Log().Errorf(ctx, "Failed to declare delay exchange: %v", err)
		return
	}

	// 创建事件
	event := OrderTimeoutEvent{
		OrderId:   orderId,
		Type:      "order_timeout",
		TimeStamp: time.Now().Format(time.RFC3339),
	}

	// 发布延迟事件
	routingKey := g.Cfg().MustGet(ctx, "rabbitmq.routingKey.orderTimeout").String()
	err = rb.PublishWithDelay(exchange, routingKey, event, delayMs)
	if err != nil {
		g.Log().Errorf(ctx, "Failed to publish orderTimeout event: %v", err)
	} else {
		g.Log().Infof(ctx, "Published orderTimeout event with %d ms delay: %+v", delayMs, event)
	}
}
