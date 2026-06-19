package cmd

import (
	"context"
	"os"
	"os/signal"
	cart_info "shop-goframe-micro-service-refacotor/app/order/utility/cart_info"
	goods "shop-goframe-micro-service-refacotor/app/order/utility/goods_info"
	"shop-goframe-micro-service-refacotor/app/order/utility/payment"
	"syscall"

	"shop-goframe-micro-service-refacotor/app/order/internal/controller/order_info"
	"shop-goframe-micro-service-refacotor/app/order/internal/controller/refund_info"
	"shop-goframe-micro-service-refacotor/app/order/internal/job"
	"shop-goframe-micro-service-refacotor/app/order/utility/consumer"
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
		Brief: "order grpc service",
		Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
			// 创建消费者管理器
			consumerManager, err := rabbitmq.NewConsumerManager(ctx)
			if err != nil {
				g.Log().Errorf(ctx, "创建消费者管理器失败: %v", err)
				return err
			}
			// 创建支付客户端
			if err := payment.InitWechatClient(); err != nil {
				g.Log().Errorf(ctx, "支付客户端初始化失败:%v", err)
				return err
			}

			// 注册order服务的消费者
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
			order_info.Register(s)
			refund_info.Register(s)
			goods.Register()
			cart_info.Register()

			job.StartOrderCompensateJob(ctx)

			// 启动 Outbox 中继任务：轮询发件箱并投递到 RabbitMQ
			job.StartOutboxRelayJob(ctx)

			s.Run()
			return nil
		},
	}
)

// setupConsumers 设置order服务的消费者
func setupConsumers(ctx context.Context, manager *rabbitmq.ConsumerManager) {
	// 添加优惠券确认结果消费者
	couponResultConsumer := consumer.NewCouponResultConsumer(ctx)
	manager.AddConsumer(couponResultConsumer)

	// 添加订单超时未支付消费者
	orderTimeoutConsumer := consumer.NewOrderTimeoutConsumer(ctx)
	manager.AddConsumer(orderTimeoutConsumer)

	// 添加订单支付成功消费者
	orderPaidSalesConsumer := consumer.NewOrderPaidSalesConsumer(ctx)
	manager.AddConsumer(orderPaidSalesConsumer)

	// 添加订单取消消费者
	orderCancelledConsumer := consumer.NewOrderCancelledConsumer(ctx)
	manager.AddConsumer(orderCancelledConsumer)

	// 可以继续添加更多消费者...
	// anotherConsumer := consumer.NewAnotherConsumer(ctx)
	// manager.AddConsumer(anotherConsumer)
}
