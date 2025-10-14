package cmd

import (
	"context"
	"os"
	"os/signal"
<<<<<<< HEAD
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/bargain_history"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/bargain_info"
=======
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/recommend_goods_info"
>>>>>>> master
	"syscall"

	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/cart_info"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/category_info"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/coupon_info"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/goods_images"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/goods_info"
	"shop-goframe-micro-service-refacotor/app/goods/internal/controller/user_coupon_info"
	"shop-goframe-micro-service-refacotor/app/goods/utility/consumer"
	"shop-goframe-micro-service-refacotor/utility/rabbitmq"

	"github.com/gogf/gf/contrib/rpc/grpcx/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
	"google.golang.org/grpc"
)

var (
	Main = gcmd.Command{
		Name:  "main",
		Usage: "main",
		Brief: "goods grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 创建消费者管理器
			consumerManager, err := rabbitmq.NewConsumerManager(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "创建消费者管理器失败: %v", err)
				return err
			}

			// 注册goods服务的消费者
			setupConsumers(ctx, consumerManager)

			// 启动消费者管理器
			err = consumerManager.Start()
			if err != nil {
				g.Log().Errorf(ctx, "启动消费者管理器失败: %v", err)
				return err
			}

			// 设置优雅关闭
			go func() {
				quit := make(chan os.Signal, 1)
				signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
				<-quit
				g.Log().Info(ctx, "正在关闭消费者管理器...")
				consumerManager.Stop()
			}()

			// 启动gRPC服务
			c := grpcx.Server.NewConfig()
			c.Options = append(c.Options, []grpc.ServerOption{
				grpcx.Server.ChainUnary(
					grpcx.Server.UnaryValidate,
				)}...,
			)
			s := grpcx.Server.New(c)
			goods_info.Register(s)
			goods_images.Register(s)
			category_info.Register(s)
			cart_info.Register(s)
			coupon_info.Register(s)
			user_coupon_info.Register(s)
<<<<<<< HEAD
			bargain_history.Register(s)
			bargain_info.Register(s)
=======
			recommend_goods_info.Register(s)
>>>>>>> master
			s.Run()
			return nil
		},
	}
)

// setupConsumers 设置goods服务的消费者
func setupConsumers(ctx context.Context, manager *rabbitmq.ConsumerManager) {
	// 添加用户注册事件消费者
	userConsumer := consumer.NewUserRegisteredConsumer(ctx)
	manager.AddConsumer(userConsumer)

	// 添加优惠券确认消费者
	couponConsumer := consumer.NewCouponConfirmConsumer(ctx)
	manager.AddConsumer(couponConsumer)

<<<<<<< HEAD
=======
	// 添加订单创建事件消费者
	orderCreatedConsumer := consumer.NewOrderCreatedConsumer(ctx)
	manager.AddConsumer(orderCreatedConsumer)

>>>>>>> master
	// 可以继续添加更多消费者...
	// anotherConsumer := consumer.NewAnotherConsumer(ctx)
	// manager.AddConsumer(anotherConsumer)
}
