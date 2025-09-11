package rabbitmq

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
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
